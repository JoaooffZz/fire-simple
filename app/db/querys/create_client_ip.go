package querys

import (
	"fmt"
	"time"
)

func (qrs *QuerysDB) CreateClientIP(ip string) error {
	query := fmt.Sprintf(`
	    INSERT INTO %s
		(client_ip, unix_first_access_date)
		VALUES ($1, $2);
	`, qrs.tabelClientsIP)

	tx, err := qrs.conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	unix := time.Now().Unix()
	_, err = tx.Exec(query, ip, unix)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
