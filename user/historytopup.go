package user

import (
	"database/sql"
	"fmt"
)

func HistoryTopUp(db *sql.DB, noHPpemilik string) (err error) {
	rows, err := db.Query("SELECT saldo, created_at FROM topup WHERE no_hp_pengirim = ?", noHPpemilik)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for rows.Next() {
		saldo := ""
		created_at := ""
		rows.Scan(
			&saldo,
			&created_at,
		)
		fmt.Println(saldo, created_at)
	}

	return err
}
