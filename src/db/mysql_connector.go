package db

import (
	"database/sql"
)

type MysqlConnect struct {
	connection   *sql.DB
	name         string // 链接用户
	password     string // 密码
	host         string // 主机名
	port         string // 端口
	databaseName string // 数据库名
	charset      string // 字符集
}

func (m *MysqlConnect) InitConn(config map[string]string) {
	// 开始初始化拿到config
	if n, ok := config["db_username"]; ok {
		m.name = n
	} else {
		panic("没有配置数据库用户名")
	}
	if n, ok := config["db_password"]; ok {
		m.password = n
	} else {
		panic("没有配置数据库密码")
	}
	if n, ok := config["db_host"]; ok {
		m.host = n
	} else {
		panic("没有配置数据库地址")
	}
	if n, ok := config["db_name"]; ok {
		m.databaseName = n
	} else {
		panic("没有配置数据库")
	}
	if n, ok := config["db_port"]; ok {
		m.port = n
	} else {
		m.port = "3306"
	}

	if n, ok := config["db_charset"]; ok {
		m.charset = n
	} else {
		m.charset = "utf8"
	}
	db, err := sql.Open("mysql", m.name+":"+m.password+"@tcp("+m.host+":"+m.port+")/"+m.databaseName+"?charset="+m.charset)
	if err != nil {
		panic(err)
	}
	m.connection = db // 赋值给结构体的属性
}

// 新建连接器
func NewConnect(config map[string]string) IConnector{
	c := &MysqlConnect{}
	c.InitConn(config)
	return c
}

// 返回链接对象
func (m *MysqlConnect) GetConn() *sql.DB {
	return m.connection
}

//func (m *MysqlConnect) Table(tableName string) IBuilder {
//	return NewSqlBuilder().InitConn(m).Table(tableName)
//}

func (m *MysqlConnect) Schema() ISchemaBuilder {
	return NewSqlSchemaBuilder().InitConn(m)
}
