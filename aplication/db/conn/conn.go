package conn

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func Connect(config SocketConfigDB) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"user=%s password=%s host=%s dbname=%s sslmode=disable",
		config.USER,
		config.PASSWORD,
		config.HOST, // <-- isso será o caminho do socket
		config.DBNAME,
	)

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		defer db.Close()
		return nil, err
	}
	return db, nil
}
