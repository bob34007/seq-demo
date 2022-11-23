package sqll

import (
	"database/sql"
)

func ConnectDB(dsn string) (*sql.DB, error) {
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return DB, nil
}

func CloseConn(db *sql.DB) error {
	return db.Close()
}
