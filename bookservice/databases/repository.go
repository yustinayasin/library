package books

import (
	books "bookservice/business"
	"errors"

	"gorm.io/gorm"
)

type BookRepository struct {
	Db *gorm.DB
}

func NewBookRepository(database *gorm.DB) books.BookRepoInterface {
	return &BookRepository{
		Db: database,
	}
}

func (repo *BookRepository) AddBook(book books.Book) (books.Book, error) {
	bookDB := FromUsecase(book)

	result := repo.Db.Create(&bookDB)

	if result.Error != nil {
		return books.Book{}, result.Error
	}

	return bookDB.ToUsecase(), nil
}

func (repo *BookRepository) EditBook(book books.Book, id int) (books.Book, error) {
	bookDb := FromUsecase(book)

	var newBook Book

	result := repo.Db.First(&newBook, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return books.Book{}, errors.New("Book not found")
		}
		return books.Book{}, errors.New("error in database")
	}

	newBook.Title = bookDb.Title
	newBook.Isbn = bookDb.Isbn
	newBook.PublicationYear = bookDb.PublicationYear
	newBook.Description = bookDb.Description

	repo.Db.Save(&newBook)
	return newBook.ToUsecase(), nil
}

func (repo *BookRepository) DeleteBook(id int) (books.Book, error) {
	var bookDb Book

	resultFind := repo.Db.First(&bookDb, id)

	if resultFind.Error != nil {
		return books.Book{}, errors.New("book not found")
	}

	result := repo.Db.Delete(&bookDb, id)

	if result.Error != nil {
		return books.Book{}, errors.New("Delete failed")
	}

	return bookDb.ToUsecase(), nil
}

func (repo *BookRepository) GetBook(id int) (books.Book, error) {
	var bookDb Book

	resultFind := repo.Db.First(&bookDb, id)

	if resultFind.Error != nil {
		return books.Book{}, errors.New("book not found")
	}

	return bookDb.ToUsecase(), nil
}
