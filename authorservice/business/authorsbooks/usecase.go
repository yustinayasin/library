package authorsbooks

import (
	"errors"
)

type AuthorsBooksUseCase struct {
	Repo AuthorsBooksRepoInterface
}

func NewUseCase(authorsBooksRepo AuthorsBooksRepoInterface) AuthorsBooksUseCaseInterface {
	return &AuthorsBooksUseCase{
		Repo: authorsBooksRepo,
	}
}

func (booksStocksUseCase *AuthorsBooksUseCase) AddAuthorBook(booksStocks AuthorBook) (AuthorBook, error) {
	if booksStocks.BookId == 0 {
		return AuthorBook{}, errors.New("Book ID cannot be empty")
	}

	if booksStocks.AuthorId == 0 {
		return AuthorBook{}, errors.New("Author ID cannot be empty")
	}

	authorsBooksRepo, err := booksStocksUseCase.Repo.AddAuthorBook(booksStocks)

	if err != nil {
		return AuthorBook{}, err
	}

	return authorsBooksRepo, nil
}

func (booksStocksUseCase *AuthorsBooksUseCase) EditAuthorBook(booksStocks AuthorBook, id int) (AuthorBook, error) {
	if id == 0 {
		return AuthorBook{}, errors.New("AuthorBook ID cannot be empty")
	}

	if booksStocks.BookId == 0 {
		return AuthorBook{}, errors.New("Book ID cannot be empty")
	}

	if booksStocks.AuthorId == 0 {
		return AuthorBook{}, errors.New("Author ID cannot be empty")
	}

	authorsBooksRepo, err := booksStocksUseCase.Repo.EditAuthorBook(booksStocks, id)

	if err != nil {
		return AuthorBook{}, err
	}

	return authorsBooksRepo, nil
}

func (booksStocksUseCase *AuthorsBooksUseCase) DeleteAuthorBook(id int) (AuthorBook, error) {
	if id == 0 {
		return AuthorBook{}, errors.New("AuthorBook ID cannot be empty")
	}

	authorsBooksRepo, err := booksStocksUseCase.Repo.DeleteAuthorBook(id)

	if err != nil {
		return AuthorBook{}, err
	}

	return authorsBooksRepo, nil
}

func (booksStocksUseCase *AuthorsBooksUseCase) GetAuthorBook(id int) (AuthorBook, error) {
	if id == 0 {
		return AuthorBook{}, errors.New("AuthorBook ID cannot be empty")
	}

	authorsBooksRepo, err := booksStocksUseCase.Repo.GetAuthorBook(id)

	if err != nil {
		return AuthorBook{}, err
	}

	return authorsBooksRepo, nil
}
