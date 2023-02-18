package user

import (
	"database/sql"
	"fmt"
)

func ReadAccount(db *sql.DB) (err error) {
	var NoHpUser string
	fmt.Println("Masukkan nomor Telepon tujuan Dude !:")
	fmt.Scanln(&NoHpUser)

	var user User
	err = db.
		QueryRow("SELECT name, no_hp FROM user WHERE no_hp = ? LIMIT 1",
			NoHpUser).
		Scan(&user.Nama, &user.NoHp)
	if err != nil {
		fmt.Println("hmm .. Akun tidak ada Dude!")
		return err
	}
	fmt.Printf("No Handphone : %s\nNama : %s\n", user.NoHp, user.Nama)

	return err
}
