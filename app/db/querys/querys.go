package querys

import (
	"database/sql"
)

type QuerysDB struct {
	conn *sql.DB
}

func New(conn *sql.DB) *QuerysDB {
	return &QuerysDB{
		conn: conn,
	}
}
