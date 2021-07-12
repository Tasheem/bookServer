package dao

import (
	"database/sql"
	"fmt"

	"github.com/Tasheem/bookServer/models"
	"github.com/google/uuid"
)

var (
	username = "root"
	password = "colts1810"
	address  = "127.0.0.1:3306"
)

func createDBIfDoesNotExist() error {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/", username, password, address)

	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS BookStore;")
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = db.Exec("USE BookStore")
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS books(" +
		"id varchar(36) NOT NULL," +
		"name varchar(100)," +
		"author_first_name varchar(100)," +
		"author_last_name varchar(100)," +
		"price FLOAT," +
		"PRIMARY KEY (id));")
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func QueryAllBooks() ([]models.Book, error) {
	err := createDBIfDoesNotExist()
	if err != nil {
		fmt.Println(err)
		// returning empty slice of books and the error.
		return make([]models.Book, 0), err
	}

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/BookStore", username, password, address)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println(err)
		return make([]models.Book, 0), err
	}
	defer db.Close()

	result, err := db.Query("SELECT * FROM books")
	if err != nil {
		fmt.Println("Error With query statement")
		fmt.Println(err)
		return make([]models.Book, 0), err
	}
	defer result.Close()

	var books []models.Book
	for result.Next() {
		var id string

		book := models.Book{}
		err := result.Scan(&id, &book.Name, &book.AuthorFirstName, &book.AuthorLastName, &book.Price)
		if err != nil {
			fmt.Println(err)
			return make([]models.Book, 0), err
		}

		book.Id = uuid.MustParse(id)

		books = append(books, book)
	}

	return books, nil
}

func Save(b models.Book) error {
	err := createDBIfDoesNotExist()
	if err != nil {
		fmt.Println(err)
		return err
	}

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/BookStore", username, password, address)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println("dao->Save: Error Opening SQL Connection statement")
		fmt.Println(err)
		return err
	}
	defer db.Close()

	insert := fmt.Sprintf("INSERT INTO books VALUES (\"%s\", \"%s\", \"%s\", \"%s\", %.2f);",
		b.Id.String(), b.Name, b.AuthorFirstName, b.AuthorLastName, b.Price)

	_, err = db.Exec(insert)
	if err != nil {
		fmt.Println("dao->Save: Error With INSERT statement")
		fmt.Println(err)
		fmt.Printf("Insert Statement: %s", insert)
		return err
	}

	return nil
}
