package user

import (
	"database/sql"
	"fmt"
)

func Register(db *sql.DB) (err error) {

	var NamaUser string
	var PasswordUser string
	var NoHpUser string
	fmt.Println("Siapa nama anda Dude! :")
	fmt.Scanln(&NamaUser)
	fmt.Println("Buat sandi anda Dude! :")
	fmt.Scanln(&PasswordUser)
	fmt.Println("Masukkan No Hp anda Dude! :")
	fmt.Scanln(&NoHpUser)

	_, err = db.
		Exec("insert into user(name,password,no_hp) values (?, ?, ?)",
			NamaUser, PasswordUser, NoHpUser)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Yeay! Anda sudah berhasil mendaftar Dude! :")
	return err
}
