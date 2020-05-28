package db

import (
	"database/sql"
)

type SqlSchemaBuilder struct {
    connector IConnector
}
func NewSqlSchemaBuilder() *SqlSchemaBuilder {
	return &SqlSchemaBuilder{}
}

func (ms *SqlSchemaBuilder) InitConn(connector IConnector) ISchemaBuilder{
    ms.connector = connector
    return ms
}

func (ms *SqlSchemaBuilder) GetConn() *sql.DB{
    return ms.connector.GetConn()
}
