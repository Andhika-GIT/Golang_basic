package PZN_Golang_Backend

import (
	"context"
	"fmt"
	"strconv"
	"testing"
)

func TestExecSqlQuery(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// insertQuery := "INSERT INTO customer(id,name) VALUES('1asdfsf','Andhika')"
	selectQuery := "SELECT id,name FROM customer"
	rows, err := db.QueryContext(ctx, selectQuery)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	// looping through all data rows
	for rows.Next() {
		var id, name string

		err := rows.Scan(&id, &name)

		// check if error

		if err != nil {
			panic(err)
		}

		fmt.Printf("id : %v\tname: %v\n", id, name)
	}

	// fmt.Println("successfully insert to database")

}
func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "admin"

	selectQuery := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, selectQuery, username, password)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	// looping through all data rows
	if rows.Next() {
		var username string

		err := rows.Scan(&username)

		// check if error

		if err != nil {
			panic(err)
		}

		fmt.Printf("username : %v", username)
	} else {
		fmt.Println("gagal login")
	}

	// fmt.Println("successfully insert to database")

}

func TestGetLastId(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	email := "dhika@gmail.com"
	comment := "this is a second comment"

	ctx := context.Background()
	sqlQuery := "INSERT INTO comments(email,comment) VALUES(? , ?)"
	result, err := db.ExecContext(ctx, sqlQuery, email, comment)

	if err != nil {
		panic(err)
	}

	// get the last inserted id
	insertId, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}

	fmt.Println("success inserted new comment with id : ", insertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	stmt, err := db.PrepareContext(ctx, "INSERT INTO comments(email,comment) VALUES(? , ?)")

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	for i := 0; i < 10; i++ {
		// create email and comment until ten times
		email := "dhikaemailke" + strconv.Itoa(i) + "@gmail.com"
		comment := "ini komen ke " + strconv.Itoa(i)
		result, err := stmt.ExecContext(ctx, email, comment)

		if err != nil {
			panic(err)
		}

		lastInsertId, _ := result.LastInsertId()
		fmt.Println("ini adalah comment id ke : ", lastInsertId)
	}
}

func TestTransaction(t *testing.T) {

	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	// begin transaction (save to database, only when transaction is commit)
	tr, err := db.Begin()

	if err != nil {
		panic(err)
	}

	sqlQuery := "INSERT INTO comments(email,comment) VALUES(? , ?)"
	// do transaction

	for i := 0; i < 10; i++ {
		// create email and comment until ten times
		email := "dhikaemailke" + strconv.Itoa(i) + "@gmail.com"
		comment := "ini komen ke " + strconv.Itoa(i)
		result, err := tr.ExecContext(ctx, sqlQuery, email, comment)

		if err != nil {
			panic(err)
		}

		lastInsertId, _ := result.LastInsertId()
		fmt.Println("ini adalah comment id ke : ", lastInsertId)
	}

	// commit transaction (save data to database)
	// err = tr.Commit()

	// rollback transaction (canceling saving process to database)
	err = tr.Rollback()

	// check for commit error (can be caused by server error, or wrong id, or duplicate data)
	if err != nil {
		panic(err)
	}

}
