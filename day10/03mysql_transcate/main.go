package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
var db *sql.DB
// 初始化数据库
func initDB()(err error){
	// 创建数据库连接
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/sql_test")
	if err != nil {
		fmt.Printf("open mysql failed, err: %v", err)
		return 
	}
	return
}

// 创建事务操作
func sqlTranscate()(err error){
		// 创建一个事务
	tx, err1 := db.Begin()
	if err1 != nil{
		fmt.Printf("begin transcation failed, err: %v\n", err1)
		return
	}

	// 执行sql语句
	_, err2 := tx.Exec("update User set age = age - 2 where name = ?", "李四")
	if err2 != nil{
		tx.Rollback()
		fmt.Printf("update 李四 failed, err: %v\n", err2)
		return
	}
	_, err3 := tx.Exec("update user set age = age + 2 where name = ?", "王五")
	if err3 != nil{
		tx.Rollback()
		fmt.Printf("update 王五 failed, err: %v\n", err3)
		return
	}
	// 提交事务
	err4 := tx.Commit()
	if err4 != nil{
		tx.Rollback()
		fmt.Printf("commit failed, err: %v\n", err4)
		return
	}
	fmt.Printf("update success\n")
	return
}

func main(){
	// 初始化数据库
	err := initDB()
	if err != nil {
		fmt.Printf("initDB failed, err: %v\n", err)
		return
	}
	// 创建事务
	err = sqlTranscate()
	if err != nil {
		fmt.Printf("sqlTranscate failed, err: %v\n", err)
		return
	}

}