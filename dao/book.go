package dao

import (
	"bookServer/models"
	"database/sql"
	"fmt"
)

var (
	username = "root"
	password = "colts1810"
	address = "127.0.0.1:3306"
)

func createDBIfDoesNotExist() {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/", username, password, address)

	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	_,err = db.Exec("CREATE DATABASE IF NOT EXISTS BookStore;")
	if err != nil {
		panic(err)
	}

	_,err = db.Exec("USE BookStore")
	if err != nil {
		panic(err)
	}

	_,err = db.Exec("CREATE TABLE IF NOT EXISTS books(" +
		"id int NOT NULL," +
		"name varchar(100)," +
		"author_first_name varchar(100)," +
		"author_last_name varchar(100)," +
		"price FLOAT," +
		"PRIMARY KEY (id));")
	if err != nil {
		panic(err)
	}
}

func QueryAllBooks() []models.Book {
	createDBIfDoesNotExist()

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/BookStore", username, password, address)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	result, err := db.Query("SELECT * FROM books")
	if err != nil {
		fmt.Println("Error With query statement")
		panic(err)
	}
	defer result.Close()

	var books []models.Book
	for result.Next() {
		var(
			id int
			name string
			authorFName string
			authorLName string
			price float64
		)

		if err := result.Scan(&id, &name, &authorFName, &authorLName, &price); err != nil {
			panic(err)
		}

		book := models.NewBook(id, price, name, authorFName, authorLName)
		books = append(books, *book)
	}

	return books
}

func Save(b models.Book) {
	createDBIfDoesNotExist()

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/BookStore", username, password, address)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	insert := fmt.Sprintf("INSERT INTO books VALUES (%d, \"%s\", \"%s\", \"%s\", %.2f);",
		b.Id, b.Name, b.AuthorFirstName, b.AuthorLastName, b.Price)

	_, err = db.Exec(insert)
	if err != nil {
		fmt.Println("Error With INSERT statement")
		panic(err)
	}
}