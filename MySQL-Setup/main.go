package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8080)/pzn-go-test")
	defer db.Close()
	if err != nil {
		panic(err)
	}
}
