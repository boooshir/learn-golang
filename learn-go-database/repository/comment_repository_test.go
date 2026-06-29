package repository

import (
	"context"
	"fmt"
	"learn-go-database/db"
	"learn-go-database/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommnetInsert(t *testing.T) {
	commentRepository := NewCommentRepository(db.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repository@mail.com",
		Comment: "Test comment from repository",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestCommnetFindById(t *testing.T) {
	commentRepository := NewCommentRepository(db.GetConnection())

	comment, err := commentRepository.FindById(context.Background(), 24)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}
func TestCommentFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(db.GetConnection())

	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	for _, comment := range comments {

		fmt.Println(comment)
	}
}
