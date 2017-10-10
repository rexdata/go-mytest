package main

import (

	"database/sql"
	"fmt"   
	_ "github.com/go-sql-driver/mysql"
)

func getdata() int{
	
	db, err := sql.Open("mysql", "root:rqdb123@tcp(202.168.1.21:3306)/cnrqhts?charset=utf8")
	CheckErr(err)
	rows, err := db.Query("select id,user_name from ca_user")
	CheckErr(err)
	var rec int =0;
	for rows.Next() {
		var id int
		var uname string
		err = rows.Scan(&id, &uname)
		
		CheckErr(err)
		fmt.Println(id)
		fmt.Println(uname)
		rec++
	}
	db.Close()
	return rec
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
	