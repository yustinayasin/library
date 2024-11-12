package authors

import (
	proto "shared/proto/authors"
	"time"
)

type Author struct {
	Id        int
	Name      string
	Bio       string
	Birthdate time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AuthorUseCaseInterface interface {
	AddAuthor(author Author) (Author, error)
	EditAuthor(author Author, id int) (Author, error)
	DeleteAuthor(id int) (Author, error)
	GetAuthor(id int) (Author, error)
	GetAuthorExist(id int) (proto.AuthorResponseExist, error)
}

type AuthorRepoInterface interface {
	AddAuthor(author Author) (Author, error)
	EditAuthor(author Author, id int) (Author, error)
	DeleteAuthor(id int) (Author, error)
	GetAuthor(id int) (Author, error)
	GetAuthorExist(id int) (proto.AuthorResponseExist, error)
}
