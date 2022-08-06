package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	dbname := "root:123456@tcp(127.0.0.1:3306)/mysql"
	db,err := sql.Open("mysql", dbname)
	if err != nil {
		fmt.Println("open mysql failed, err: ", err)
		return
	}
	defer db.Close()
	err1 := db.Ping()
	if err1 != nil{
		fmt.Println("login database failed, err: ", err1)
		return
	}
	fmt.Println("链接数据库成功")
}