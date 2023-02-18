package user

import (
	"database/sql"
	"fmt"
)

func TopUp(db *sql.DB, noHPpemilik string) (err error) {
	var Nominal string
	fmt.Println("Mau Top-up berapa dude! :")
	fmt.Scanln(&Nominal)

	_, err = db.Exec("update user set saldo=saldo+? where no_hp = ?", Nominal, noHPpemilik)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	db.Exec("insert into topup(no_hp_pengirim,saldo) values (?, ?)",
		noHPpemilik, Nominal)

	fmt.Println("Yeay top-up succes DUde!")
	return err
}
