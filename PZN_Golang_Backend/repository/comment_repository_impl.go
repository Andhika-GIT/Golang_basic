package repository

import (
	"PZN_Golang_Backend/PZN_Golang_Backend/entity"
	"context"
	"database/sql"
	"errors"
)

// create struct for comment
type CommentRepositoryImpl struct {
	DB *sql.DB
}

// create method for comment struct that implements comment repository interface
func (repository *CommentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	sqlQuery := "INSERT INTO comment(email,comment) VALUES (?, ?)"
	result, err := repository.DB.ExecContext(ctx, sqlQuery, comment.Email, comment.Comment)

	if err != nil {
		return comment, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return comment, err
	}

	comment.Id = int32(id)
	return comment, nil
}

func (repository *CommentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	sqlQuery := "SELECT id,email,comment FROM comment WHERE id = ?"
	row, err := repository.DB.QueryContext(ctx, sqlQuery, id)

	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}
	defer row.Close()
	if row.Next() {

		err := row.Scan(&comment.Id, &comment.Email, &comment.Comment)

		if err != nil {
			return comment, err
		}

		return comment, nil
	} else {
		return comment, errors.New("account not found")
	}
}

func (repository *CommentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	sqlQuery := "SELECT id,email,comment FROM comment"
	rows, err := repository.DB.QueryContext(ctx, sqlQuery)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []entity.Comment

	for rows.Next() {

		comment := entity.Comment{}

		err := rows.Scan(&comment.Id, &comment.Email, &comment.Comment)

		if err != nil {
			return comments, err
		}
		comments = append(comments, comment)

	}

	return comments, nil

}

// create new function to export the comment struct and it's method
func NewCommentRepository(db *sql.DB) CommentRepository {
	return &CommentRepositoryImpl{DB: db}
}
