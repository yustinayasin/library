package booksstocks

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

func (repo *BooksStocksRepository) AddBooksStocks(book booksStocks.BooksStocks) (booksStocks.BooksStocks, error) {
	bookDB := FromUsecase(book)

	result := repo.Db.Create(&bookDB)

	if result.Error != nil {
		return booksStocks.BooksStocks{}, result.Error
	}

	return bookDB.ToUsecase(), nil
}

func (repo *BooksStocksRepository) EditBooksStocks(book booksStocks.BooksStocks, id int) (booksStocks.BooksStocks, error) {
	bookDb := FromUsecase(book)

	var newBooksStocks BooksStocks

	result := repo.Db.First(&newBooksStocks, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return booksStocks.BooksStocks{}, errors.New("BooksStocks not found")
		}
		return booksStocks.BooksStocks{}, errors.New("error in database")
	}

	newBooksStocks.BookId = bookDb.BookId
	newBooksStocks.TotalStock = bookDb.TotalStock
	newBooksStocks.AvailableStock = bookDb.AvailableStock

	repo.Db.Save(&newBooksStocks)
	return newBooksStocks.ToUsecase(), nil
}

func (repo *BooksStocksRepository) DeleteBooksStocks(id int) (booksStocks.BooksStocks, error) {
	var bookDb BooksStocks

	resultFind := repo.Db.First(&bookDb, id)

	if resultFind.Error != nil {
		return booksStocks.BooksStocks{}, errors.New("book not found")
	}

	result := repo.Db.Delete(&bookDb, id)

	if result.Error != nil {
		return booksStocks.BooksStocks{}, errors.New("Delete failed")
	}

	return bookDb.ToUsecase(), nil
}

func (repo *BooksStocksRepository) GetBooksStocks(id int) (booksStocks.BooksStocks, error) {
	var bookDb BooksStocks

	resultFind := repo.Db.First(&bookDb, id)

	if resultFind.Error != nil {
		return booksStocks.BooksStocks{}, errors.New("book not found")
	}

	return bookDb.ToUsecase(), nil
}
