package querys

import (
	"database/sql"
)

type QuerysDB struct {
	conn           *sql.DB
	tabelClientsIP string
}

func New(conn *sql.DB) *QuerysDB {
	return &QuerysDB{
		conn:           conn,
		tabelClientsIP: "clients_ip",
	}
}
