
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
    InitConn(connector IConnector) ISchemaBuilder
    GetConn() *sql.DB
    // 创建数据库
    CreateDatabase(args ...string) (sql.Result, error)
    //创建表
    CreateTable(tableName string, call func(table IBlueprint)) error
}

type SchemaBuilder struct {
	connector IConnector
}

func (s *SchemaBuilder) InitConn(connector IConnector) ISchemaBuilder {
	s.connector = connector
	return s
}

func (s *SchemaBuilder) GetConn() *sql.DB {
	return s.connector.GetConn()
}

func (s *SchemaBuilder) CreateDatabase(args ...string) (sql.Result, error) {
	return nil, nil
}

func (s *SchemaBuilder) CreateTable(tableName string, call func(table IBlueprint)) error {
	return nil
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
