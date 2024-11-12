package authorsbooks

import (
	"time"
)

type AuthorBook struct {
	BookId    int
	AuthorId  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AuthorsBooksUseCaseInterface interface {
	AddAuthorBook(authorBook AuthorBook) (AuthorBook, error)
	EditAuthorBook(authorBook AuthorBook, id int) (AuthorBook, error)
	DeleteAuthorBook(id int) (AuthorBook, error)
	GetAuthorBook(id int) (AuthorBook, error)
}

type AuthorsBooksRepoInterface interface {
	AddAuthorBook(authorBook AuthorBook) (AuthorBook, error)
	EditAuthorBook(authorBook AuthorBook, id int) (AuthorBook, error)
	DeleteAuthorBook(id int) (AuthorBook, error)
	GetAuthorBook(id int) (AuthorBook, error)
}
