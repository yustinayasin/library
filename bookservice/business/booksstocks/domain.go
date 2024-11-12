package booksstocks

import (
	"time"
)

type BookStock struct {
	Id             int
	BookId         int
	TotalStock     int
	AvailableStock int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type BooksStocksUseCaseInterface interface {
	AddBooksStocks(booksStocks BookStock) (BookStock, error)
	EditBooksStocks(booksStocks BookStock, id int) (BookStock, error)
	DeleteBooksStocks(id int) (BookStock, error)
	GetBooksStocks(id int) (BookStock, error)
}

type BooksStocksRepoInterface interface {
	AddBooksStocks(booksStocks BookStock) (BookStock, error)
	EditBooksStocks(booksStocks BookStock, id int) (BookStock, error)
	DeleteBooksStocks(id int) (BookStock, error)
	GetBooksStocks(id int) (BookStock, error)
}
