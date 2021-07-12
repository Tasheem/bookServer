package models

import "github.com/google/uuid"

type Book struct {
	Id uuid.UUID
	Name string
	AuthorFirstName string
	AuthorLastName string
	Price float64
}

func NewBook(id uuid.UUID, price float64, name, authorFName, authorLName string) *Book {
	var book *Book = new(Book)
	book.Id = id
	book.Price = price
	book.Name = name
	book.AuthorFirstName = authorFName
	book.AuthorLastName = authorLName
	return book
}

/*func (b Book) getID() string {
	return b.id
}

func (b *Book) setID(id string) {
	b.id = id
}

func (b Book) getName() string {
	return b.name
}

func (b *Book) setName(name string) {
	b.name = name
}

func (b Book) getAuthorFirstName() string {
	return b.authorFirstName
}

func (b *Book) setAuthorFirstName(fName string) {
	b.authorFirstName = fName
}

func (b Book) getAuthorLastName() string {
	return b.authorLastName
}

func (b *Book) setAuthorLastName(lName string) {
	b.authorLastName = lName
}*/
