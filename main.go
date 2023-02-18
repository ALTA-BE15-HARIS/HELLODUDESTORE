package main

import (
	"PROJECT1/database"
	"PROJECT1/user"
	"fmt"
)

func init() {
	database.Connect()
}

func main() {
	var (
		err     error
		noHp    string
		pilihan int
	)

	db := database.GetConnection()

	for {
		if pilihan != 2 {
			fmt.Println("1. Register")
			fmt.Println("2. Login")

			fmt.Println("Masukkan pilihan Dude!")
			fmt.Scanln(&pilihan)
		}

		switch pilihan {
		case 1:
			fmt.Println("Silahkan Register Dude!")
			if err = user.Register(db); err != nil {
				return
			}
			fallthrough
		case 2:
			if noHp == "" {
				fmt.Println("Silahkan Login Dude!")
				if noHp, err = user.Login(db); err != nil {
					pilihan = 0
					return
				}
			}

		default:
			fmt.Println("pilihan tidak ditemukan Dude!")
			return
		}

		var menu int
		fmt.Println("1. Read Account")
		fmt.Println("2. Updated Account")
		fmt.Println("3. Delete Account")
		fmt.Println("4. Transfer")
		fmt.Println("5. Top Up")
		fmt.Println("6. Histori Top Up")
		fmt.Println("7. Histori Transfer")
		fmt.Println("8. Cari user lain")
		fmt.Println("0. keluar")

		fmt.Println("Masukkan pilihan Dude!")
		fmt.Scanln(&menu)

		switch menu {
		case 1:
			if err = user.ReadAccount(db); err != nil {
				return
			}
		case 2:
			if err = user.Update(db, noHp); err != nil {
				return
			}
		case 3:
			if err = user.Delete(db, noHp); err != nil {
				return
			}
			pilihan = 0
		case 4:
			if err = user.Transfer(db, noHp); err != nil {
				return
			}
		case 5:
			// fmt.Println("Silahkan Login Dude!")
			if err = user.TopUp(db, noHp); err != nil {
				return
			}
		case 6:
			if err = user.HistoryTopUp(db, noHp); err != nil {
				return
			}
		case 7:
			if err = user.HistoryTransfer(db, noHp); err != nil {
				return
			}
		default:
			fmt.Println("Terima kasih telah bertransaksi!")
			fmt.Println("Senang bisa membantu anda Dude!")
			return
		}
	}
}
