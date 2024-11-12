package controller

import (
	booksstocks "bookservice/business/booksstocks"
	booksStocksProto "shared/proto/booksstocks"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func FromUsecase(book booksstocks.BookStock) *booksStocksProto.BooksStocksResponse {
	return &booksStocksProto.BooksStocksResponse{
		Id:             int32(book.Id),
		BookId:         int32(book.BookId),
		TotalStock:     int32(book.TotalStock),
		AvailableStock: int32(book.AvailableStock),
		CreatedAt:      timestamppb.New(book.CreatedAt),
		UpdatedAt:      timestamppb.New(book.UpdatedAt),
	}
}

func FromUsecaseList(book []booksstocks.BookStock) []*booksStocksProto.BooksStocksResponse {
	var bookResponse []*booksStocksProto.BooksStocksResponse

	for _, v := range book {
		bookResponse = append(bookResponse, FromUsecase(v))
	}

	return bookResponse
}
