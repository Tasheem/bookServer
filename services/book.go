package services

import (
	"fmt"

	"github.com/Tasheem/bookServer/dao"
	"github.com/Tasheem/bookServer/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func GetAllBooks() ([]models.Book, error) {
	return dao.QueryAllBooks()
}

func CreateBook(b models.Book) error {
	b.Id = uuid.New()
	fmt.Printf("Id: %d\n", b.Id)
	return dao.Save(b)
}
