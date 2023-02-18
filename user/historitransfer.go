package user

import (
	"database/sql"
	"fmt"
)

func HistoryTransfer(db *sql.DB, noHPpemilik string) (err error) {
	rows, err := db.Query("SELECT no_hp_penerima, saldo, created_at FROM transfer WHERE no_hp_pengirim = ?", noHPpemilik)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for rows.Next() {
		no_hp_penerima := ""
		saldo := ""
		created_at := ""
		rows.Scan(
			&no_hp_penerima,
			&saldo,
			&created_at,
		)
		fmt.Println(saldo, created_at)
	}

	return err
}
