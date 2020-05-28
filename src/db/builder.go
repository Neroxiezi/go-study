
package db

import (
	"database/sql"
)

// 表操作接口
type IBuilder interface {
    InitConn(connector IConnector) IBuilder
    GetConn() *sql.DB
    //Table(tableName string) IBuilder
}
// 数据库操作接口
type ISchemaBuilder interface {
    InitConn(connector IConnector) IBuilder
    GetConn() *sql.DB
}

type SchemaBuilder struct {
	connector IConnector
}

type Builder struct {
	//connector IConnector
}

// 放入连接器
func (b *Builder) InitConn(connector IConnector) IBuilder {
	return nil
}

func (b *Builder) GetConn() *sql.DB {
	return nil
}
