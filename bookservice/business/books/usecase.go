package books

import (
	"errors"
)

type BookUseCase struct {
	Repo BookRepoInterface
}

func NewUseCase(bookRepo BookRepoInterface) BookUseCaseInterface {
	return &BookUseCase{
		Repo: bookRepo,
	}
}

func (bookUseCase *BookUseCase) AddBook(book Book) (Book, error) {
	if book.Title == "" {
		return Book{}, errors.New("Title cannot be empty")
	}

	if book.Isbn == "" {
		return Book{}, errors.New("ISBN cannot be empty")
	}

	if book.PublicationYear == "" {
		return Book{}, errors.New("Publication year cannot be empty")
	}

	if book.Description == "" {
		return Book{}, errors.New("Description cannot be empty")
	}

	bookRepo, err := bookUseCase.Repo.AddBook(book)

	if err != nil {
		return Book{}, err
	}

	return bookRepo, nil
}

func (bookUseCase *BookUseCase) EditBook(book Book, id int) (Book, error) {
	if id == 0 {
		return Book{}, errors.New("book ID cannot be empty")
	}

	if book.Title == "" {
		return Book{}, errors.New("Title cannot be empty")
	}

	if book.Isbn == "" {
		return Book{}, errors.New("ISBN cannot be empty")
	}

	if book.PublicationYear == "" {
		return Book{}, errors.New("Publication year cannot be empty")
	}

	if book.Description == "" {
		return Book{}, errors.New("Description cannot be empty")
	}

	bookRepo, err := bookUseCase.Repo.EditBook(book, id)

	if err != nil {
		return Book{}, err
	}

	return bookRepo, nil
}

func (bookUseCase *BookUseCase) DeleteBook(id int) (Book, error) {
	if id == 0 {
		return Book{}, errors.New("Book ID cannot be empty")
	}

	bookRepo, err := bookUseCase.Repo.DeleteBook(id)

	if err != nil {
		return Book{}, err
	}

	return bookRepo, nil
}

func (bookUseCase *BookUseCase) GetBook(id int) (Book, error) {
	if id == 0 {
		return Book{}, errors.New("Book ID cannot be empty")
	}

	bookRepo, err := bookUseCase.Repo.GetBook(id)

	if err != nil {
		return Book{}, err
	}

	return bookRepo, nil
}
