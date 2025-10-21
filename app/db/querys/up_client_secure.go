package querys

import "fmt"

func (qrs *QuerysDB) UpdateClientSecure(ip string, isSecure bool) error {
	query := fmt.Sprintf(`
	    UPDATE %s
		SET
		    is_secure = $2
		WHERE client_ip = $1;
	`, qrs.tabelClientsIP)

	_, err := qrs.conn.Exec(query, ip, isSecure)
	return err
}
