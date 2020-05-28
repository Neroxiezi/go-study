package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// 构建一个连接器接口，所有实现这个接口的结构体都可以作为参数传入，每个单独的连接器
type IConnector interface {
	InitConn(config map[string]string) // 设置连接器
	GetConn() *sql.DB                 // 获取数据库连接器接口
	Schema() ISchemaBuilder           // 获取一个空Builder，用于数据库的增删改查表
	//Table(tableName string) IBuilder  // 表的操作
}

type Connector struct {
	db *sql.DB
}

func (c *Connector) InitConn(config map[string]string) {

}

func (c *Connector) GetConn() *sql.DB {
	return nil
}

//func (c *Connector) Table(tableName string) IBuilder {
//	return nil
//}

func (c *Connector) Schema() IBuilder {
	return nil
}
