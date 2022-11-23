package sqll

import "database/sql"

func QueryWithResult(db *sql.DB, sql string) (*sql.Rows, error) {
	return db.Query(sql)
}

func QueryWithNoResult(db *sql.DB, sql string) error {
	_, err := db.Exec(sql)
	return err
}
