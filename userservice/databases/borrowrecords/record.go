package borrowrecords

import (
	"time"
	borrowrecords "userservice/business/borrowrecords"
)

type BorrowRecord struct {
	Id         int       `gorm:"primaryKey;unique"`
	UserId     int       `gorm:"foreignKey:UserId;references:Id"`
	BookId     int       `gorm:"foreignKey:BookId;references:Id"`
	BorrowDate time.Time `gorm:"type:date"`
	ReturnDate time.Time `gorm:"type:date"`
	IsReturn   bool
	CreatedAt  time.Time `gorm:"default:current_timestamp"`
	UpdatedAt  time.Time `gorm:"default:current_timestamp;on_update:current_timestamp"`
}

func (borrowRecord BorrowRecord) ToUsecase() borrowrecords.BorrowRecord {
	return borrowrecords.BorrowRecord{
		Id:         borrowRecord.Id,
		UserId:     borrowRecord.UserId,
		BookId:     borrowRecord.BookId,
		BorrowDate: borrowRecord.BorrowDate,
		ReturnDate: borrowRecord.ReturnDate,
		IsReturn:   borrowRecord.IsReturn,
		CreatedAt:  borrowRecord.CreatedAt,
		UpdatedAt:  borrowRecord.UpdatedAt,
	}
}

func ToUsecaseList(borrowRecord []BorrowRecord) []borrowrecords.BorrowRecord {
	var newUsers []borrowrecords.BorrowRecord

	for _, v := range borrowRecord {
		newUsers = append(newUsers, v.ToUsecase())
	}
	return newUsers
}

func FromUsecase(borrowRecord borrowrecords.BorrowRecord) BorrowRecord {
	return BorrowRecord{
		Id:         borrowRecord.Id,
		UserId:     borrowRecord.UserId,
		BookId:     borrowRecord.BookId,
		BorrowDate: borrowRecord.BorrowDate,
		ReturnDate: borrowRecord.ReturnDate,
		IsReturn:   borrowRecord.IsReturn,
		CreatedAt:  borrowRecord.CreatedAt,
		UpdatedAt:  borrowRecord.UpdatedAt,
	}
}
