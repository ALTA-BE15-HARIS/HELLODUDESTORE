package user

import (
	"database/sql"
	"fmt"
	"strconv"
)

func Transfer(db *sql.DB, noHPpemilik string) (err error) {
	var NoHp string
	var Nominal string
	fmt.Println("Masukkan Nomor Tujuan Dude :")
	fmt.Scanln(&NoHp)
	fmt.Println("Nominal yang ingin anda transfer Dude :")
	fmt.Scanln(&Nominal)

	if noHPpemilik == NoHp {
		fmt.Println("Tidak boleh mengirim ke nomor hp sendiri Dude!")
		return err
	}

	nominal, err := strconv.Atoi(Nominal)
	if err != nil {
		fmt.Println("Nominal harus angka Dude")
		return err
	}

	var saldo int64
	err = db.
		QueryRow("SELECT saldo FROM user WHERE no_hp = ? LIMIT 1", noHPpemilik).Scan(&saldo)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	if saldo < int64(nominal) {
		fmt.Println("Saldo anda tidak cukup Dude!")
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	defer tx.Rollback()

	_, err = tx.Exec("update user set saldo=saldo-? where no_hp = ?", Nominal, noHPpemilik)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	_, err = tx.Exec("update user set saldo=saldo+? where no_hp = ?", Nominal, NoHp)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	db.Exec("insert into transfer(no_hp_pengirim,saldo, no_hp_penerima) values (?, ?, ?)",
		noHPpemilik, Nominal, NoHp)

	fmt.Println("Yeay. Transfer berhasil Dude!")
	fmt.Printf("No Handphone : %s\nSaldo Terkirim : %s\n", noHPpemilik, Nominal)

	return err
}
