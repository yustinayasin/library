package controller

import (
	borrowRecordsProto "shared/proto/users"
	borrowrecords "userservice/business/borrowrecords"
)

func ToUsecase(borrowrecord *borrowRecordsProto.BorrowRecordsRequest) *borrowrecords.BorrowRecord {
	return &borrowrecords.BorrowRecord{
		UserId:     int(borrowrecord.UserId),
		BookId:     int(borrowrecord.BookId),
		BorrowDate: borrowrecord.BorrowDate.AsTime(),
		ReturnDate: borrowrecord.ReturnDate.AsTime(),
		IsReturn:   borrowrecord.IsReturn,
	}
}
