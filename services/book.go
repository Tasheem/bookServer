package services

import (
	"bookServer/dao"
	"bookServer/models"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
)

func GetAllBooks() []models.Book {
	books := dao.QueryAllBooks()

	return books
}

func CreateBook(b models.Book) {
	b.Id = rand.Intn(1000)
	dao.Save(b)
}
