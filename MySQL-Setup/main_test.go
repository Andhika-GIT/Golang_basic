package main

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8080)/pzn-go-test")

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
}
