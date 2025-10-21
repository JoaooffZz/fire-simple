package querys

import "os"

func (qrs *QuerysDB) InitTabelsDB() error {
	file, err := os.ReadFile("/usr/local/fire_simple/tables/init.sql")
	if err != nil {
		return err
	}

	sqlContent := string(file)

	_, err = qrs.conn.Exec(sqlContent)
	if err != nil {
		return err
	}
	return nil
}
