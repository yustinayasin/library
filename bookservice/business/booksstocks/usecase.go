package booksstocks

import (
	"errors"
)

type BooksStocksUseCase struct {
	Repo BooksStocksRepoInterface
}

func NewUseCase(booksStocksRepo BooksStocksRepoInterface) BooksStocksUseCaseInterface {
	return &BooksStocksUseCase{
		Repo: booksStocksRepo,
	}
}

func (booksStocksUseCase *BooksStocksUseCase) AddBooksStocks(booksStocks BookStock) (BookStock, error) {
	if booksStocks.BookId == 0 {
		return BookStock{}, errors.New("Book ID cannot be empty")
	}

	if booksStocks.TotalStock == 0 {
		return BookStock{}, errors.New("Total stock cannot be empty")
	}

	if booksStocks.AvailableStock == 0 {
		return BookStock{}, errors.New("Available stock cannot be empty")
	}

	booksStocksRepo, err := booksStocksUseCase.Repo.AddBooksStocks(booksStocks)

	if err != nil {
		return BookStock{}, err
	}

	return booksStocksRepo, nil
}

func (booksStocksUseCase *BooksStocksUseCase) EditBooksStocks(booksStocks BookStock, id int) (BookStock, error) {
	if id == 0 {
		return BookStock{}, errors.New("BookStock ID cannot be empty")
	}

	if booksStocks.BookId == 0 {
		return BookStock{}, errors.New("Book ID cannot be empty")
	}

	if booksStocks.TotalStock == 0 {
		return BookStock{}, errors.New("Total stock cannot be empty")
	}

	if booksStocks.AvailableStock == 0 {
		return BookStock{}, errors.New("Available stock cannot be empty")
	}

	booksStocksRepo, err := booksStocksUseCase.Repo.EditBooksStocks(booksStocks, id)

	if err != nil {
		return BookStock{}, err
	}

	return booksStocksRepo, nil
}

func (booksStocksUseCase *BooksStocksUseCase) DeleteBooksStocks(id int) (BookStock, error) {
	if id == 0 {
		return BookStock{}, errors.New("BookStock ID cannot be empty")
	}

	booksStocksRepo, err := booksStocksUseCase.Repo.DeleteBooksStocks(id)

	if err != nil {
		return BookStock{}, err
	}

	return booksStocksRepo, nil
}

func (booksStocksUseCase *BooksStocksUseCase) GetBooksStocks(id int) (BookStock, error) {
	if id == 0 {
		return BookStock{}, errors.New("BookStock ID cannot be empty")
	}

	booksStocksRepo, err := booksStocksUseCase.Repo.GetBooksStocks(id)

	if err != nil {
		return BookStock{}, err
	}

	return booksStocksRepo, nil
}
