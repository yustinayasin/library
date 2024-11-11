package controller

import (
	books "bookservice/business/books"
	booksProto "shared/proto/books"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func FromUsecase(book books.Book) *booksProto.BookResponse {
	return &booksProto.BookResponse{
		Id:              int32(book.Id),
		Title:           book.Title,
		Isbn:            book.Isbn,
		PublicationYear: book.PublicationYear,
		Description:     book.Description,
		CreatedAt:       timestamppb.New(book.CreatedAt),
		UpdatedAt:       timestamppb.New(book.UpdatedAt),
	}
}

func FromUsecaseList(book []books.Book) []*booksProto.BookResponse {
	var bookResponse []*booksProto.BookResponse

	for _, v := range book {
		bookResponse = append(bookResponse, FromUsecase(v))
	}

	return bookResponse
}
