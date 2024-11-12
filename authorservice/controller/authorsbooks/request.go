package controller

import (
	authorsbooks "authorservice/business/authorsbooks"
	authorProto "shared/proto/authors"
)

func ToUsecase(book *authorProto.AuthorsBooksRequest) *authorsbooks.AuthorBook {
	return &authorsbooks.AuthorBook{
		BookId:   int(book.BookId),
		AuthorId: int(book.AuthorId),
	}
}
