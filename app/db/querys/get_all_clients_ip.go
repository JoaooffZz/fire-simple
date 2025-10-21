package querys

import "fmt"

func (qrs *QuerysDB) GetAllClientsIP() ([]string, error) {
	query := fmt.Sprintf(`
	    SELECT client_ip
		FROM %s;
	`, qrs.tabelClientsIP)

	rows, err := qrs.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []string
	for rows.Next() {
		var client string
		err := rows.Scan(&client)
		if err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}
	return clients, nil
}
