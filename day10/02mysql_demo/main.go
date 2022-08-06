package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 声明一个连接池
var db *sql.DB
// 声明一个结构体来保存数据
type user struct {
	id int
	name string
	age int
}
var u []user
// init DB 创建连接池
func initDB()(err error) {
	// 创建连接池 db
	dbname := "root:123456@tcp(127.0.0.1:3306)/sql_test"
	db, err = sql.Open("mysql", dbname)
	if err != nil {
		return fmt.Errorf("open mysql failed, err: %v", err)
	}

	db.SetMaxOpenConns(10)  // 设置连接池的最大连接数
	db.SetMaxIdleConns(5)	// 设置连接池的最大闲置连接数

	return
}
// sql 单行查询并将查询结果插入到指定结构体中
func queryOneRow(querystr string, arg... interface{})(err error){
	// 通过 sql 语句获取数据表对象
	rowObj := db.QueryRow(querystr, arg...)
	// 将获取的数据表对象数据插入到结构体中
	var u1 user
	err1 := rowObj.Scan(&u1.id, &u1.name, &u1.age)
	if err1 != nil{
		return errors.New("query failed")
	}
	u = append(u, u1)
	return nil
}

// sql 多行查询，并将查询结果插入到指定结构体中
func queryMultRow(querystr string, args... interface{})(err error){
	rowsObj, err := db.Query(querystr, args...)
	if err != nil{
		fmt.Printf("query failed, err: %v\n",err)
		return
	}
	defer rowsObj.Close()
	var u1 user
	for rowsObj.Next() {
		rowsObj.Scan(&u1.id, &u1.name, &u1.age)
		u = append(u, u1)
	}
	return 
}
// sql 插入数据
func insertData(insertstr string, args... interface{})(err error){
	ret, err := db.Exec(insertstr, args...)
	if err != nil{
		fmt.Printf("insert failed, err: %v\n",err)
		return
	}
	id, err1 := ret.LastInsertId()
	if err1 != nil{
		fmt.Printf("get id failed, err: %v\n",err1)
		return
	}
	fmt.Println(id)
	return nil
}

// sql 更新数据
func updateData(updatestr string, args... interface{})(err error){
	ret, err := db.Exec(updatestr, args...)
	if err != nil{
		fmt.Printf("update failed, err: %v\n",err)
		return
	}
	affect, err1 := ret.RowsAffected()
	if err1 != nil{
		fmt.Printf("get affect failed, err: %v\n",err1)
		return
	}
	fmt.Println(affect)
	return nil
}
// sql 删除数据
func deleteData(deletestr string, args... interface{})(err error){
	ret, err := db.Exec(deletestr, args...)
	if err != nil{
		fmt.Printf("delete failed, err: %v\n",err)
		return
	}
	affect, err1 := ret.RowsAffected()
	if err1 != nil{
		fmt.Printf("get affect failed, err: %v\n",err1)
		return
	}
	fmt.Println(affect)
	return nil
}

// sql 预处理
func prepareSql(){
	// 预处理语句
	sqlstr := "insert into user(name, age) values(?, ?)"
	// 预处理对象
	stmt, err := db.Prepare(sqlstr)
	if err != nil{
		fmt.Printf("prepare failed, err: %v\n",err)
		return
	}
	// 关闭预处理对象
	defer stmt.Close()
	// 执行预处理语句
	var m = make(map[string]interface{}, 10)
	m = map[string]interface{}{
		"张三": 18,
		"李四": 20,
		"王五": 22,
	}
	for k, v := range m{
		_, err := stmt.Exec(k, v)
		if err != nil{
			fmt.Printf("exec failed, err: %v\n",err)
			return
		}
	}
}
func main(){
	// 初始化连接池
	err := initDB()
	if err != nil {
		fmt.Printf("连接 MySQL 失败：%v\n", err)
		return
	}
	// 设置 sql 查询语句
	// querystr := "select id,name,age from user"
	// 设置 sql 插入语句
	// insertstr := "insert into user(name,age) values(?,?)"
	// 设置 sql 更新语句
	// updatestr := "update user set age = ? where name = ?"
	// 设置 sql 删除语句
	// deletestr := "delete from user where name = ?"
	prepareSql()
}