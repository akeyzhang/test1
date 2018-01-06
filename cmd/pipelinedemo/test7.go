package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

/*func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func insertdata() {
	db, err := sql.Open("mysql", "root:akey123@tcp(127.0.0.1:3306)/test?charset=utf8")
	checkErr(err)
	defer db.Close()
	//insert data
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	checkErr(err)
	res, err := stmt.Exec("zxzxr", "电脑部", "2002-11-15")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
}

func updatedata() {
	db, err := sql.Open("mysql", "root:akey123@tcp(127.0.0.1:3306)/test?charset=utf8")
	checkErr(err)
	defer db.Close()
	stmt, err := db.Prepare("UPDATE userinfo SET departname=? where username=?")
	checkErr(err)
	_, err = stmt.Exec("manager room", "zxzxr")
	checkErr(err)

}

func deletedata() {
	db, err := sql.Open("mysql", "root:akey123@tcp(127.0.0.1:3306)/test?charset=utf8")
	checkErr(err)
	defer db.Close()
	stmt, err := db.Prepare("DELETE FROM userinfo  where username=?")
	checkErr(err)
	_, err = stmt.Exec("zxr")
	checkErr(err)
}

func querydata() {
	var username, departname, created string
	db, err := sql.Open("mysql", "root:akey123@tcp(127.0.0.1:3306)/test?charset=utf8")
	checkErr(err)
	defer db.Close()
	rows, err := db.Query("SELECT username,departname,created from userinfo ORDER BY username")
	checkErr(err)
	for rows.Next() {
		err = rows.Scan(&username, &departname, &created)
		fmt.Println(username, departname, created)
	}

}
*/

func main() {
	//insertdata()
	//updatedata()
	//deletedata()
	//querydata()

}
