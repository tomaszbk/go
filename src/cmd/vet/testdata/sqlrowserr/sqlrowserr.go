// This file contains the test for failure to call sql.Rows.Err.

package sqlrowserr

import "database/sql"

func missingErr(db *sql.DB) {
	rows, err := db.Query("") // ERROR "sql.Rows .rows. is used in Next loop at line 17 without final check of rows.Err"
	if err != nil {
		return
	}
	defer rows.Close() // ignore error
	for rows.Next() {  // L17
		println(rows.Scan())
	}
}
