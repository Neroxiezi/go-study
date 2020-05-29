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
	config["db_name"] = "pfinal"
	config["db_charset"] = "utf8"
	c := db.NewConnect(config)
	//res, err := c.Schema().CreateDatabaseIfNotExists("pfinal")
	err := c.Schema().CreateTable("test", func(table db.IBlueprint) {
		table.Increments("id").Comment("自增id") // 设置备注与主键
		table.String("name").Comment("姓名")
		table.Integer("age").Nullable().Comment("年龄") // 允许为空
		// 添加普通索引  _index
		table.Index("user_id")
		// 添加唯一索引  _unique_index
		table.UniqueIndex("user_id")
		// 添加组合索引
		table.Index("user_id", "name")
		// 添加全文索引,只对MyISAM表有效
		table.FullTextIndex("user_id")
	})
	if err != nil {
		p(err)
	}

}
