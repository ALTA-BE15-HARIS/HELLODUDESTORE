package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connect() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/project1")
	if err != nil {
		log.Fatal(err)
	}
}

func GetConnection() *sql.DB {
	return db
}
