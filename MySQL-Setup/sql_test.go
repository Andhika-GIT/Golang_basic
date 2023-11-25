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

	insertQuery := "INSERT INTO customer(id,name) VALUES('1asdfsf','Andhika')"
	_, err := db.ExecContext(ctx, insertQuery)

	if err != nil {
		panic(err)
	}

	fmt.Println("successfully insert to database")

}
