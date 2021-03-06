package dao

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Tasheem/bookServer/models"
	"github.com/google/uuid"
)

var (
	username = "root"
	password = "colts1810"
	address  = "127.0.0.1:3306"
)

func createDBIfDoesNotExist() (*sql.DB, error) {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/", username, password, address)

	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		fmt.Println(err)
		db.Close()
		return nil, err
	}

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS LibraryManagementSystem;")
	if err != nil {
		fmt.Println(err)
		db.Close()
		return nil, err
	}

	_, err = db.Exec("USE LibraryManagementSystem")
	if err != nil {
		fmt.Println(err)
		db.Close()
		return nil, err
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
		db.Close()
		return nil, err
	}

	return db, nil
}

func QueryAllBooks() ([]models.Book, error) {
	db, err := createDBIfDoesNotExist()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		// returning empty slice of books and the error.
		return make([]models.Book, 0), err
	}

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
	db, err := createDBIfDoesNotExist()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}

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

func UpdatePrice(id uuid.UUID, price float64) error {
	db, err := createDBIfDoesNotExist()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}

	update := fmt.Sprintf("UPDATE books SET price = %.2f WHERE id = \"%s\";", price, id.String())

	result, err := db.Exec(update)
	if err != nil {
		fmt.Println("dao->UpdatePrice: Error With Update statement")
		fmt.Println(err)
		fmt.Printf("Update Statement: %s", update)
		return err
	}

	if rowsAffected, err := result.RowsAffected(); rowsAffected == 0 {
		if err != nil {
			fmt.Println("Error thrown by RowsAffected() function.")
			return err
		}
		return errors.New("row not found")
	}

	return nil
}

func DeleteBook(id string) error {
	db, err := createDBIfDoesNotExist()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}

	deleteStatement := fmt.Sprintf("DELETE FROM books WHERE id = \"%s\";", id)

	_, err = db.Exec(deleteStatement)
	if err != nil {
		fmt.Println("dao->DeleteBook: Error With Delete statement")
		fmt.Println(err)
		fmt.Printf("Delete Statement: %s", deleteStatement)
		return err
	}

	return nil
}