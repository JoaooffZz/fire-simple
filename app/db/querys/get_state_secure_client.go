package querys

import (
	"database/sql"
	"errors"
	"fmt"
)

func (qrs *QuerysDB) GetStateSecureClient(clientIP string) (bool, error) {
	query := fmt.Sprintf(`
	    SELECT is_secure
		FROM %s
		WHERE client_ip =$1;
	`, qrs.tabelClientsIP)

	var isSecure bool
	err := qrs.conn.QueryRow(query, clientIP).Scan(&isSecure)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return true, nil
		}
		return false, err
	}
	return isSecure, nil
}
