package user

import (
	"database/sql"
	"fmt"
)

func Update(db *sql.DB, noHp string) (err error) {
	var NamaUser string
	var PasswordUser string
	fmt.Println("Nama yang ingin di ubah apa Dude :")
	fmt.Scanln(&NamaUser)
	fmt.Println("Password yang ingin di ubah apa Dude :")
	fmt.Scanln(&PasswordUser)

	_, err = db.
		Exec("update user set name=?, password=? where no_hp = ?",
			NamaUser, PasswordUser, noHp)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Yeay. Update berhasil Dude!")
	fmt.Printf("No Handphone : %s\nNama : %s\n", noHp, NamaUser)

	return err
}
