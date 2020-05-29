package db

import (
	"database/sql"
	"errors"
	"strings"
)

type SqlSchemaBuilder struct {
	connector IConnector
}

func NewSqlSchemaBuilder() *SqlSchemaBuilder {
	return &SqlSchemaBuilder{}
}

func (ms *SqlSchemaBuilder) InitConn(connector IConnector) ISchemaBuilder {
	ms.connector = connector
	return ms
}

func (ms *SqlSchemaBuilder) GetConn() *sql.DB {
	return ms.connector.GetConn()
}

func (ms *SqlSchemaBuilder) CreateDatabase(args ...string) (sql.Result, error) {
	if len(args) == 0 {
		return nil, errors.New(`CreateDatabase 函数必须传递参数`)
	}
	dbName := args[0]              // 指定数据库名称
	dbCharset := "utf8"            // 指定字符集
	dbCollate := "utf8_general_ci" // 指定默认排序规则
	switch len(args) {
	case 1:
	case 2:
		if strings.ToUpper(args[1]) == "UTF8" || strings.ToUpper(args[1]) == "UTF-8" {
			dbCharset = "utf8"
		} else if strings.ToUpper(args[1]) == "GBK" {
			dbCharset = "gbk"
			dbCollate = "gbk_chinese_ci"
		} else {
			// 除了这utf8和gbk这两种，其他的，必须指明排序规则
			return nil, errors.New("you must set your collate when create a new database")
		}
	case 3:
	default:
		return nil, errors.New(`too many arguments in the CreateDatabase function`)
	}
	stmt, err := ms.GetConn().Prepare(`CREATE DATABASE ` + dbName + ` DEFAULT CHARSET ` + dbCharset + ` COLLATE ` + dbCollate)
	if err != nil {
		return nil, err
	}
	return stmt.Exec()
}

func (ms *SqlSchemaBuilder) CreateDatabaseIfNotExists(args ...string) (sql.Result, error) {
	if len(args) == 0 {
		return nil, errors.New(`CreateDatabaseIfNotExists 函数必须传递参数`)
	}
	dbName := args[0]              // 指定数据库名称
	dbCharset := "utf8"            // 指定字符集
	dbCollate := "utf8_general_ci" // 指定默认排序规则
	switch len(args) {
	case 1:
	case 2:
		if strings.ToUpper(args[1]) == "UTF8" || strings.ToUpper(args[1]) == "UTF-8" {
			dbCharset = "utf8"
		} else if strings.ToUpper(args[1]) == "GBK" {
			dbCharset = "gbk"
			dbCollate = "gbk_chinese_ci"
		} else {
			// 除了这utf8和gbk这两种，其他的，必须指明排序规则
			return nil, errors.New("you must set your collate when create a new database")
		}
	case 3:
	default:
		return nil, errors.New(`too many arguments in the CreateDatabase function`)
	}
	stmt, err := ms.GetConn().Prepare(`CREATE DATABASE ` + dbName + ` DEFAULT CHARSET ` + dbCharset + ` COLLATE ` + dbCollate)
	if err != nil {
		return nil, err
	}
	return stmt.Exec()
}

