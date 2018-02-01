package main

import (
	"database/sql"
	"fmt"
	"github.com/golang/net/websocket"
	//"github.com/json-iterator/go"
	"encoding/json"
	_ "github.com/lib/pq"
	"net/http"
	"strings"
	"test1/akeyfunction"
	"time"
)

/*
akeySql功能: sql執行器,將傳入的sql語句傳給數據庫執行並返回jason結果,
             暫且使用cockroachDB做為數據庫
*/

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func QueryOrExce(sqlstr string) string {
	fmt.Println(time.Now().Format(akeyfunction.STANDARDDATEFORMAT), ": 开始访问数据库...")
	var dbConnStr string = "postgresql://akey@113.108.248.46:26257/test?sslmode=disable"
	db, err := sql.Open("postgres", dbConnStr)
	checkErr(err)
	defer db.Close()
	if strings.Index(strings.ToUpper(strings.TrimLeft(sqlstr, " ")), "SELECT") == 0 {
		//表明是Query語句
		//fmt.Println("found!")
		rows, err := db.Query(sqlstr)
		checkErr(err)
		columns, err := rows.Columns() // 得到列
		checkErr(err)

		count := len(columns)
		tableData := make([]map[string]interface{}, 0)
		values := make([]interface{}, count)
		scanArgs := make([]interface{}, count)
		for rows.Next() {
			for i := 0; i < count; i++ {
				scanArgs[i] = &values[i]
			}
			rows.Scan(scanArgs...)
			entry := make(map[string]interface{})
			for i, col := range columns {
				var v interface{}
				val := values[i]
				b, ok := val.([]byte)
				if ok {
					v = string(b)
				} else {
					v = val
				}
				entry[col] = v
			}
			tableData = append(tableData, entry)
		}
		fmt.Println(time.Now().Format(akeyfunction.STANDARDDATEFORMAT), ": 数据生成完毕.")
		jsonData, err := json.Marshal(tableData)
		checkErr(err)
		fmt.Println(time.Now().Format(akeyfunction.STANDARDDATEFORMAT), ": 打包JSON串完毕.")
		return string(jsonData)

	} else {
		//表明是Insert,update或delete語句
		_, err = db.Exec(sqlstr)
		checkErr(err)
		return ""

	}
}

//tagstr拆包函数
func splitTagstr(tagstr string) (string, string) {
	var tag, sqlstr string
	ipos := strings.Index(tagstr, "#**#")
	tag = tagstr[0:ipos]
	sqlstr = tagstr[ipos+4 : len(tagstr)]
	return tag, sqlstr

}

func h_index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func h_akeySql(ws *websocket.Conn) {
	var tagstr string

	for {
		err := websocket.Message.Receive(ws, &tagstr)
		checkErr(err)
		fmt.Println(time.Now().Format(akeyfunction.STANDARDDATEFORMAT), ": 接收到的tagstr: ", tagstr)
		//将tagstr拆包成两部分,tag和sqlstr.
		tag, sqlstr := splitTagstr(tagstr)
		//预处理
		presqlstr := sqlstr + " limit 500"
		resultstr := QueryOrExce(presqlstr)
		if resultstr != "" { //resultstr返回空白字串表示處理的是非select語句,不向客戶端返回處理結果.
			resultstr = tag + "#**#" + resultstr
			fmt.Println(time.Now().Format(akeyfunction.STANDARDDATEFORMAT), ": 向客户端发送预处理JSON串...")
			err = websocket.Message.Send(ws, resultstr)
			checkErr(err)
			resultstr = QueryOrExce(sqlstr)
			//再将resultstr用bkno打包发送出去.
			resultstr = tag + "#**#" + resultstr
			fmt.Println(time.Now().Format(akeyfunction.STANDARDDATEFORMAT), ": 向客户端发送JSON串...")
			err = websocket.Message.Send(ws, resultstr)
			checkErr(err)
		}
	}

}

func main() {
	fmt.Println(time.Now().Format(akeyfunction.STANDARDDATEFORMAT), ": 开始启动...")
	fmt.Println()

	http.HandleFunc("/", h_index)
	//绑定socket方法
	http.Handle("/akeySql", websocket.Handler(h_akeySql))
	//开始监听
	http.ListenAndServe(":666", nil)

	//debug用
	/*	var sqlstr string = "select delv_no, delv_date, cust_no from delvmain order by delv_date"
		result := QueryOrExce(sqlstr)
		fmt.Println(result)*/

}
