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

func main() {
	initDB()
//	p(db)
	sqlStr := "show databases"
    rows, err := db.Query(sqlStr)
    if err!=nil {
        p("查询失败\n")
        return
    }
    defer rows.Close()
    for rows.Next() {
        var dbs string
        err = rows.Scan(&dbs)
        p(dbs)
    }
}
