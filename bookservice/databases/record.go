package books

import (
	books "bookservice/business"
	"time"
)

type Book struct {
	Id              int `gorm:"primaryKey;unique"`
	Title           string
	Isbn            string `gorm:"unique"`
	PublicationYear string
	Description     string
	CreatedAt       time.Time `gorm:"default:current_timestamp"`
	UpdatedAt       time.Time `gorm:"default:current_timestamp;on_update:current_timestamp"`
}

func (book Book) ToUsecase() books.Book {
	return books.Book{
		Id:              book.Id,
		Title:           book.Title,
		Isbn:            book.Isbn,
		PublicationYear: book.PublicationYear,
		Description:     book.Description,
		CreatedAt:       book.CreatedAt,
		UpdatedAt:       book.UpdatedAt,
	}
}

func ToUsecaseList(book []Book) []books.Book {
	var newBooks []books.Book

	for _, v := range book {
		newBooks = append(newBooks, v.ToUsecase())
	}
	return newBooks
}

func FromUsecase(book books.Book) Book {
	return Book{
		Id:              book.Id,
		Title:           book.Title,
		Isbn:            book.Isbn,
		PublicationYear: book.PublicationYear,
		Description:     book.Description,
		CreatedAt:       book.CreatedAt,
		UpdatedAt:       book.UpdatedAt,
	}
}
