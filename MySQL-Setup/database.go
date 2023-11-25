package main

import (
	"database/sql"
	"time"
)

func getConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8080)/pzn-go-test")

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
