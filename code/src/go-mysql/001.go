package main
import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)
var p = fmt.Println
func checkError(e error) {
    if e != nil {
        panic(e)
    }
}

func main (){
    db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang?charset=utf8")
    checkError(err)
    p(db)
}
