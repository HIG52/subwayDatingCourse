package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DbConnection() {

	var db *sql.DB

	var connectionString = "datero:datero1234@tcp(localhost:3306)/datero"
	var err error

	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	fmt.Println("connect success", db)
	defer db.Close()

}
