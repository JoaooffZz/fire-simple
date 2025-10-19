package conn

import (
	"database/sql"
	"fmt"

	q "github.com/JoaooffZz/fire-simple/app/db/querys"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func Connect(config SocketConfigDB) (*q.QuerysDB, error) {
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
