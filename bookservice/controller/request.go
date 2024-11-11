package controller

import (
	books "bookservice/business"
	booksProto "shared/proto/books"
)

func ToUsecase(book *booksProto.BookRequest) *books.Book {
	return &books.Book{
		Title:           book.Title,
		Isbn:            book.Isbn,
		PublicationYear: book.PublicationYear,
		Description:     book.Description,
	}
}
