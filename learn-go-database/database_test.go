package main

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "bangsyir:qwerty123!@#@tcp(localhost:3306)/go_database")

	if err != nil {
		panic(err)
	}
	defer db.Close()
}
