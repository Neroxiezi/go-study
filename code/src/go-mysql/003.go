package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var p = fmt.Println
var db *sql.DB

func initDB() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/golang"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func main  () {
    initDB()
    sqlStr :="create table user(id int not null,name varchar(255),sex tinyint)"
    res,err := db.Exec(sqlStr)
    if err !=nil {
        p(err)
    }
    p(res)
}
