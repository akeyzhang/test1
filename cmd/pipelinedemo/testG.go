package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-adodb"
)

func checkErr(str string, err error) {
	if err != nil {
		fmt.Println(str)
		panic(err)
	}
}

func main() {
	db, err := sql.Open("adodb", "Provider=SQLOLEDB;Data Source=192.168.1.103;Initial Catalog=chen;User ID=sa;Password=dell1226")
	checkErr("1111", err)
	defer db.Close()

	rows, err := db.Query("select username, password from users")
	checkErr("2222", err)

	for rows.Next() {
		var username, password string
		rows.Scan(&username, &password)
		fmt.Println(username, password)
	}

	//db.Exec("insert into users(username,password) values('aaa','bbb')")
}
