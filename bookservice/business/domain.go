package books

import (
	"time"
)

type Book struct {
	Id              int
	Title           string
	Isbn            string
	PublicationYear string
	Description     string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type BookUseCaseInterface interface {
	AddBook(book Book) (Book, error)
	EditBook(book Book, id int) (Book, error)
	DeleteBook(id int) (Book, error)
	GetBook(id int) (Book, error)
}

type BookRepoInterface interface {
	AddBook(book Book) (Book, error)
	EditBook(book Book, id int) (Book, error)
	DeleteBook(id int) (Book, error)
	GetBook(id int) (Book, error)
}
