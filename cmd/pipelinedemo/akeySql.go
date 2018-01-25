package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/golang/net/websocket"
	_ "github.com/lib/pq"
	"net/http"
	"strings"
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

//sql語句解析處理函數
func QueryOrExec(sqlstr string) string {
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
		jsonData, err := json.Marshal(tableData)
		checkErr(err)
		return string(jsonData)

	} else {
		//表明是Insert,update或delete語句
		_, err = db.Exec(sqlstr)
		checkErr(err)
		return ""

	}
}

func h_index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

//通過websocket接收前台傳入的sql語句,並回傳解析處理後的結果(有可能是結果集的json串,也有可能是空串).
func h_akeySql(ws *websocket.Conn) {
	var sqlstr string

	for {
		err := websocket.Message.Receive(ws, &sqlstr)
		checkErr(err)
		fmt.Println("接收的sql語句: ", sqlstr)
		resultstr := QueryOrExec(sqlstr)
		err = websocket.Message.Send(ws, resultstr)
		checkErr(err)
	}

}

func main() {
	fmt.Println("启动时间")
	fmt.Println(time.Now())

	http.HandleFunc("/", h_index)
	//绑定socket方法
	http.Handle("/akeySql", websocket.Handler(h_akeySql))
	//开始监听
	http.ListenAndServe(":666", nil)

	//debug用
	//var sqlstr string = "select cust_no,cust_sname from customer"
	//result := QueryOrExce(sqlstr)
	//fmt.Println(result)

}
