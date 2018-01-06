package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func insertdata() {
	db, err := sql.Open("postgres", "postgresql://akey@localhost:26257/test?sslmode=disable")
	checkErr(err)
	defer db.Close()
	_, err = db.Exec("insert into userinfo(username,departname,created) values('zxr', 'IT dept', '2002-11-15')")
	checkErr(err)
}

func updatedata() {
	db, err := sql.Open("postgres", "postgresql://akey@localhost:26257/test?sslmode=disable")
	checkErr(err)
	defer db.Close()
	_, err = db.Exec("update userinfo set departname='管理部' where username='jhf'")
	checkErr(err)
}

func deletedata() {
	db, err := sql.Open("postgres", "postgresql://akey@localhost:26259/test?sslmode=disable")
	checkErr(err)
	defer db.Close()
	_, err = db.Exec("delete from userinfo where username='zxr'")
	checkErr(err)
}

func querydata() {
	var username, departname, created string
	db, err := sql.Open("postgres", "postgresql://akey@localhost:26257/test?sslmode=disable")
	checkErr(err)
	defer db.Close()
	rows, err := db.Query("SELECT * from userinfo ORDER BY username")
	checkErr(err)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&username, &departname, &created)
		fmt.Println(username, departname, created)
	}

}

func main() {
	//insertdata()
	//updatedata()
	//deletedata()
	querydata()

}
