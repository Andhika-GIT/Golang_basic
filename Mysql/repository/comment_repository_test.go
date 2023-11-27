package repository

import (
	MySQL "PZN_Golang_Backend/Mysql"
	"PZN_Golang_Backend/Mysql/entity"
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	var db = MySQL.GetConnection()
	CommentRepository := NewCommentRepository(db)

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repo1@test.com",
		Comment: "repo1 comment",
	}

	result, err := CommentRepository.Insert(ctx, comment)

	if err != nil {
		panic(err)
	}

	fmt.Println("new comment sucessfully : ", result)

}

func TestFindCommenyById(t *testing.T) {
	var db = MySQL.GetConnection()
	CommentRepository := NewCommentRepository(db)

	ctx := context.Background()
	comment, err := CommentRepository.FindById(ctx, 1)

	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindAllComments(t *testing.T) {
	var db = MySQL.GetConnection()
	CommentRepository := NewCommentRepository(db)

	ctx := context.Background()

	comments, err := CommentRepository.FindAll(ctx)

	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}
