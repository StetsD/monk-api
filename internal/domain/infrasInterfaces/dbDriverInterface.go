package infrasInterfaces

import "database/sql"

type DbDriver interface {
	Query(qString string, fields ...interface{}) (*sql.Rows, error)
}
