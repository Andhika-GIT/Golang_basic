package repository

import (
	"PZN_Golang_Backend/Mysql-learning/entity"
	"context"
	"database/sql"
	"errors"
)

type CommentRepositoryImpl struct {
	DB *sql.DB
}

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
	defer row.Close()
	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}

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

func (repository *CommentRepositoryImpl) FindAll(ctx context.Context, id int32) ([]entity.Comment, error) {
	sqlQuery := "SELECT id,email,comment FROM comment WHERE id = ?"
	row, err := repository.DB.QueryContext(ctx, sqlQuery, id)
	defer row.Close()

	if err != nil {
		return nil, err
	}

	var comments []entity.Comment

	for row.Next() {

		comment := entity.Comment{}

		err := row.Scan(&comment.Id, &comment.Email, &comment.Comment)

		if err != nil {
			return comments, err
		}
		comments = append(comments, comment)

	}

	return comments, nil

}
