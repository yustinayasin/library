package booksstocks

import (
	"time"
)

type BooksStocks struct {
	Id             int
	BookId         int
	TotalStock     int
	AvailableStock int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type BooksStocksUseCaseInterface interface {
	AddBooksStocks(booksStocks BooksStocks) (BooksStocks, error)
	EditBooksStocks(booksStocks BooksStocks, id int) (BooksStocks, error)
	DeleteBooksStocks(id int) (BooksStocks, error)
	GetBooksStocks(id int) (BooksStocks, error)
}

type BooksStocksRepoInterface interface {
	AddBooksStocks(booksStocks BooksStocks) (BooksStocks, error)
	EditBooksStocks(booksStocks BooksStocks, id int) (BooksStocks, error)
	DeleteBooksStocks(id int) (BooksStocks, error)
	GetBooksStocks(id int) (BooksStocks, error)
}
