package main

import (
	"context"
	"fmt"
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
