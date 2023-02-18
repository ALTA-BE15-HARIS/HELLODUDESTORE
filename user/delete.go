package user

import (
	"database/sql"
	"fmt"
)

func Delete(db *sql.DB, noHp string) (err error) {

	var konfirmasi string
	fmt.Println("Anda serius mau hapus akun Dude ? (Y/n)")
	fmt.Scanln(&konfirmasi)

	if konfirmasi != "Y" {
		fmt.Println("Tidak jadi menghapus akun Dude")
		return err
	}

	_, err = db.
		Exec("delete from user where no_hp = ?", noHp)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("delete success!")

	return err
}
