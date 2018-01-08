package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var dbConnStr string = "postgresql://akey@113.108.248.46:26257/test?sslmode=disable"

//var dbConnStr string = "postgresql://postgres:akey123@localhost:5432/test?sslmode=disable"

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func insertdata() {
	db, err := sql.Open("postgres", dbConnStr)
	checkErr(err)
	defer db.Close()
	_, err = db.Exec("insert into person(perno,pername) values('1005', 'jhf')")
	checkErr(err)
}

func updatedata() {
	db, err := sql.Open("postgres", dbConnStr)
	checkErr(err)
	defer db.Close()
	_, err = db.Exec("update person set pername='zxzxr' where perno='1004'")
	checkErr(err)
}

func deletedata() {
	db, err := sql.Open("postgres", dbConnStr)
	checkErr(err)
	defer db.Close()
	_, err = db.Exec("delete from person where perno='1004'")
	checkErr(err)
}

func querydata() {
	var perno, pername, subject string
	var sccore int
	db, err := sql.Open("postgres", dbConnStr)
	checkErr(err)
	defer db.Close()
	rows, err := db.Query("select A.*,B.subject,B.sccore from person A left outer join chenji B on A.perno=B.perno order by A.perno")
	checkErr(err)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&perno, &pername, &subject, &sccore)
		fmt.Println(perno, pername, subject, sccore)
	}

}

func main() {
	//insertdata()
	//updatedata()
	//deletedata()
	querydata()

}
