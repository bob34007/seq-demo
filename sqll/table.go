package sqll

import (
	"database/sql"
)

func CreateDatabase(db *sql.DB, sql string) error {
	_, err := db.Exec(sql)
	return err
}

func CreateTable(db *sql.DB, sql string) error {
	_, err := db.Exec(sql)
	return err
}
