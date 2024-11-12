package controller

import (
	booksstocks "bookservice/business/booksstocks"
	booksStocksProto "shared/proto/booksstocks"
)

func ToUsecase(book *booksStocksProto.BooksStocksRequest) *booksstocks.BookStock {
	return &booksstocks.BookStock{
		BookId:         int(book.BookId),
		TotalStock:     int(book.TotalStock),
		AvailableStock: int(book.AvailableStock),
	}
}
