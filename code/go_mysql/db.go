package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(172.100.0.5:3306)/go_study?charset=utf8")
	checkErr(err)
	// 插入数据
	stmt, err := db.Prepare("INSERT userinfo SET username=?,department=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "研发部门", "2021-6-21")
	checkErr(err)
	fmt.Println(res)

	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)

	// 更新数据

	stmt1, err := db.Prepare("UPDATE userinfo SET username=? where uid=?")
	checkErr(err)
	res1, err := stmt1.Exec("PFinal", id)
	checkErr(err)
	affect, err := res1.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	// 查询数据

	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	for rows.Next() {
		var uid int
		var username, department, created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid, username, department, created)
	}

	// 删除数据
	stmt2, err := db.Prepare("DELETE FROM userinfo WHERE uid=?")
	checkErr(err)
	res2, err := stmt2.Exec(id)
	checkErr(err)
	affect, err = res2.RowsAffected()
	fmt.Println(affect)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
