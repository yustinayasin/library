package authorsbooks

import (
	booksStocks "bookservice/business/booksstocks"
	"errors"

	"gorm.io/gorm"
)

type BooksStocksRepository struct {
	Db *gorm.DB
}

func NewBooksStocksRepository(database *gorm.DB) booksStocks.BooksStocksRepoInterface {
	return &BooksStocksRepository{
		Db: database,
	}
}

func (repo *BooksStocksRepository) AddBooksStocks(book booksStocks.BookStock) (booksStocks.BookStock, error) {
	bookDB := FromUsecase(book)

	result := repo.Db.Create(&bookDB)

	if result.Error != nil {
		return booksStocks.BookStock{}, result.Error
	}

	return bookDB.ToUsecase(), nil
}

func (repo *BooksStocksRepository) EditBooksStocks(book booksStocks.BookStock, id int) (booksStocks.BookStock, error) {
	bookStockdDB := FromUsecase(book)

	var newBooksStocks BookStock

	result := repo.Db.First(&newBooksStocks, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return booksStocks.BookStock{}, errors.New("BookStock not found")
		}
		return booksStocks.BookStock{}, errors.New("error in database")
	}

	newBooksStocks.BookId = bookStockdDB.BookId
	newBooksStocks.TotalStock = bookStockdDB.TotalStock
	newBooksStocks.AvailableStock = bookStockdDB.AvailableStock

	repo.Db.Save(&newBooksStocks)
	return newBooksStocks.ToUsecase(), nil
}

func (repo *BooksStocksRepository) DeleteBooksStocks(id int) (booksStocks.BookStock, error) {
	var bookStockdDB BookStock

	resultFind := repo.Db.First(&bookStockdDB, id)

	if resultFind.Error != nil {
		return booksStocks.BookStock{}, errors.New("book not found")
	}

	result := repo.Db.Delete(&bookStockdDB, id)

	if result.Error != nil {
		return booksStocks.BookStock{}, errors.New("Delete failed")
	}

	return bookStockdDB.ToUsecase(), nil
}

func (repo *BooksStocksRepository) GetBooksStocks(id int) (booksStocks.BookStock, error) {
	var bookStockdDB BookStock

	resultFind := repo.Db.First(&bookStockdDB, id)

	if resultFind.Error != nil {
		return booksStocks.BookStock{}, errors.New("book not found")
	}

	return bookStockdDB.ToUsecase(), nil
}
