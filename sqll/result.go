package sqll

import (
	"database/sql"
)

// pos need sorted
func ParseResult(r *sql.Rows) (int64, error) {
	defer r.Close()
	var res int64
	var err error
	for r.Next() {
		err = r.Scan(&res)
		if err != nil {
			return res, err
		}
	}
	return res, err
}
