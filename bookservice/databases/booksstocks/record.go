package booksstocks

import (
	booksStocks "bookservice/business/booksstocks"
	"time"
)

type BookStock struct {
	Id             int `gorm:"primaryKey;unique"`
	BookId         int `gorm:"foreignKey:BookId;references:Id"`
	TotalStock     int `gorm:"unique"`
	AvailableStock int
	CreatedAt      time.Time `gorm:"default:current_timestamp"`
	UpdatedAt      time.Time `gorm:"default:current_timestamp;on_update:current_timestamp"`
}

func (book BookStock) ToUsecase() booksStocks.BookStock {
	return booksStocks.BookStock{
		Id:             book.Id,
		BookId:         book.BookId,
		TotalStock:     book.TotalStock,
		AvailableStock: book.AvailableStock,
		CreatedAt:      book.CreatedAt,
		UpdatedAt:      book.UpdatedAt,
	}
}

func ToUsecaseList(book []BookStock) []booksStocks.BookStock {
	var newBooksStockss []booksStocks.BookStock

	for _, v := range book {
		newBooksStockss = append(newBooksStockss, v.ToUsecase())
	}
	return newBooksStockss
}

func FromUsecase(book booksStocks.BookStock) BookStock {
	return BookStock{
		Id:             book.Id,
		BookId:         book.BookId,
		TotalStock:     book.TotalStock,
		AvailableStock: book.AvailableStock,
		CreatedAt:      book.CreatedAt,
		UpdatedAt:      book.UpdatedAt,
	}
}
