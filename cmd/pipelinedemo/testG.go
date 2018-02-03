package main

/*從sql2000導數據到cockroachDB的工具.*/

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-adodb"
	"strconv"
	"strings"
)

func checkErr(tag string, err error) {
	if err != nil {
		fmt.Println(tag)
		panic(err)
	}
}

func processDelvmain() {
	mdb, err := sql.Open("adodb", "Provider=SQLOLEDB;Data Source=192.168.1.195;Initial Catalog=MRPII;User ID=sa;Password=123456")
	checkErr("1111", err)
	defer mdb.Close()
	cdb, err := sql.Open("postgres", "postgresql://akey@192.168.1.129:26257/test?sslmode=disable")
	checkErr("2222", err)
	defer cdb.Close()

	//從sql2000取數據 insert into cockroachDB.
	rows, err := mdb.Query("select delv_no,delv_date,cust_no,user_no,upd_date from delvmain where upd_date>='2001-01-01' and upd_date<'2005-01-01'")
	checkErr("3333", err)

	fmt.Println("開始處理......")
	icount := 0
	for rows.Next() {
		var delv_no, delv_date, cust_no, upd_date string
		var user_no string
		//fetch a row from sql2000.
		rows.Scan(&delv_no, &delv_date, &cust_no, &user_no, &upd_date)
		//fmt.Println(delv_no, delv_date, cust_no, user_no, upd_date)
		sqlstr := "insert into delvmain values('" + strings.Trim(delv_no, " ") + "','" + delv_date + "','" + strings.Trim(cust_no, " ") + "'," + user_no + ",'" + upd_date + "')"
		//fmt.Println(sqlstr)
		//insert into cockroachDB delvmain table.
		_, err = cdb.Exec(sqlstr)
		checkErr("444", err)
		icount += 1
		if icount%1000 == 0 {
			fmt.Println("已處理: ", icount)
		}
	}
	fmt.Println("處理結束.")
}

func processDelvdetail() {
	mdb, err := sql.Open("adodb", "Provider=SQLOLEDB;Data Source=192.168.1.195;Initial Catalog=MRPII;User ID=sa;Password=123456")
	checkErr("1111", err)
	defer mdb.Close()
	cdb, err := sql.Open("postgres", "postgresql://akey@192.168.1.129:26257/test?sslmode=disable")
	checkErr("2222", err)
	defer cdb.Close()

	//從sql2000取數據 insert into cockroachDB.
	rows, err := mdb.Query("select delv_no,ord_no,delv_qty,delv_spare,weight as unt_weight,remark from delvdetail where delv_no in (select delv_no from delvmain where delv_date>=getdate()-600)")
	checkErr("3333", err)

	fmt.Println("開始處理......")
	icount := 0
	for rows.Next() {
		var delv_no, ord_no, remark string
		var delv_qty, delv_spare, unt_weight int
		//fetch a row from sql2000.
		rows.Scan(&delv_no, &ord_no, &delv_qty, &delv_spare, &unt_weight, &remark)
		//fmt.Println(delv_no, ord_no, delv_qty, delv_spare, unt_weight, remark)
		sqlstr := "insert into delvdetail values('" + strings.Trim(delv_no, " ") + "','" + strings.Trim(ord_no, " ") + "'," + strconv.Itoa(delv_qty) + "," + strconv.Itoa(delv_spare) + "," + strconv.Itoa(unt_weight) + ",'" + remark + "')"
		//fmt.Println(sqlstr)
		//insert into cockroachDB delvdetail table.
		_, err = cdb.Exec(sqlstr)
		checkErr("444", err)
		icount += 1
		if icount%1000 == 0 {
			fmt.Println("已處理: ", icount)
		}
	}
	fmt.Println("處理結束.", icount)

}

func main() {
	//processDelvmain()
	processDelvdetail()
}
