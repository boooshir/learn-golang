package main

import (
	"context"
	"database/sql"
	"fmt"
	"learn-go-database/db"
	"strconv"
	"testing"
	"time"
)

func TestExecSQL(t *testing.T) {
	db := db.GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUES('joko', 'anwar')"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")

}

func TestQuerySQL(t *testing.T) {
	db := db.GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("id", id)
		fmt.Println("name", name)
	}
}

func TestQuerySQLComplex(t *testing.T) {
	db := db.GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool

		err = rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("===============================")
		fmt.Println("id", id)
		fmt.Println("name", name)
		if email.Valid {
			fmt.Println("email", email.String)
		}
		fmt.Println("balance", balance)
		fmt.Println("rating", rating)
		if birthDate.Valid {
			fmt.Println("birthDate", birthDate.Time.String())
		}
		fmt.Println("createdAt", createdAt)
		fmt.Println("married", married)
	}
}

func TestSQLInjection(t *testing.T) {
	db := db.GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "admin"

	script := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	fmt.Println(script)
	rows, err := db.QueryContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses login", username)
	} else {
		fmt.Println("gagal login")
	}
}

func TestExecSQLSafe(t *testing.T) {
	db := db.GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "arief'; DROP TABLE user; #"
	password := "arief"
	script := "INSERT INTO user(username, password) VALUES(?,?)"
	_, err := db.ExecContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")

}
func TestAutoIncrement(t *testing.T) {
	db := db.GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "test@gmail.com"
	comment := "hello this is first comment"
	script := "INSERT INTO comments(email, comment) VALUES(?,?)"
	result, err := db.ExecContext(ctx, script, email, comment)
	if err != nil {
		panic(err)
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert new comments with id", insertId)

}

func TestPreparateStatement(t *testing.T) {
	db := db.GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO comments(email, comment) VALUES(?,?)"
	statement, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		email := strconv.Itoa(i) + "test@gmail.com"
		comment := "this is test comment" + strconv.Itoa(i)
		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		commentId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Success create comment with id :", commentId)
	}
	defer statement.Close()
}

func TestTransaction(t *testing.T) {
	db := db.GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	script := "INSERT INTO comments(email, comment) VALUES(?,?)"
	// do transaction
	for i := 0; i < 10; i++ {
		email := strconv.Itoa(i) + "sudin@gmail.com"
		comment := "this is sudin comment " + strconv.Itoa(i)
		result, err := tx.ExecContext(ctx, script, email, comment)
		if err != nil {
			panic(err)
		}
		commentId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Success create comment with id :", commentId)
	}

	err = tx.Rollback()
	if err != nil {
		panic(err)
	}
}
