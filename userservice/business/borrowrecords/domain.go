package borrowrecords

import (
	"time"
)

type BorrowRecord struct {
	Id         int
	UserId     int
	BookId     int
	BorrowDate time.Time
	ReturnDate time.Time
	IsReturn   bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type BorrowRecordUseCaseInterface interface {
	AddBorrowRecord(borrowRecord BorrowRecord) (BorrowRecord, error)
	EditBorrowRecord(borrowRecord BorrowRecord, id int) (BorrowRecord, error)
	DeleteBorrowRecord(id int) (BorrowRecord, error)
	GetBorrowRecord(id int) (BorrowRecord, error)
}

type BorrowRecordRepoInterface interface {
	AddBorrowRecord(borrowRecord BorrowRecord) (BorrowRecord, error)
	EditBorrowRecord(borrowRecord BorrowRecord, id int) (BorrowRecord, error)
	DeleteBorrowRecord(id int) (BorrowRecord, error)
	GetBorrowRecord(id int) (BorrowRecord, error)
}
