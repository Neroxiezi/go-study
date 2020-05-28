package main

import (
	"db"
	"fmt"
)

var p = fmt.Println

func main() {
	config := make(map[string]string)
	config["db_username"] = "root"
	config["db_password"] = "root"
	config["db_host"] = "127.0.0.1"
	config["db_port"] = "3306"
	config["db_name"] = "golong"
	config["db_charset"] = "utf8"
	c := db.NewConnect(config)
    res, err := c.Schema().CreateDatabaseIfNotExists("test")
    p(res)
}
