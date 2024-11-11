package booksstocks

import (
	booksStocks "bookservice/business/booksstocks"
	"time"
)

type BooksStocks struct {
	Id             int `gorm:"primaryKey;unique"`
	BookId         int `gorm:"foreignKey:BookId;references:Id"`
	TotalStock     int `gorm:"unique"`
	AvailableStock int
	CreatedAt      time.Time `gorm:"default:current_timestamp"`
	UpdatedAt      time.Time `gorm:"default:current_timestamp;on_update:current_timestamp"`
}

func (book BooksStocks) ToUsecase() booksStocks.BooksStocks {
	return booksStocks.BooksStocks{
		Id:             book.Id,
		BookId:         book.BookId,
		TotalStock:     book.TotalStock,
		AvailableStock: book.AvailableStock,
		CreatedAt:      book.CreatedAt,
		UpdatedAt:      book.UpdatedAt,
	}
}

func ToUsecaseList(book []BooksStocks) []booksStocks.BooksStocks {
	var newBooksStockss []booksStocks.BooksStocks

	for _, v := range book {
		newBooksStockss = append(newBooksStockss, v.ToUsecase())
	}
	return newBooksStockss
}

func FromUsecase(book booksStocks.BooksStocks) BooksStocks {
	return BooksStocks{
		Id:             book.Id,
		BookId:         book.BookId,
		TotalStock:     book.TotalStock,
		AvailableStock: book.AvailableStock,
		CreatedAt:      book.CreatedAt,
		UpdatedAt:      book.UpdatedAt,
	}
}
