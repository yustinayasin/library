package controller

import (
	borrowRecordsProto "shared/proto/users"
	borrowrecords "userservice/business/borrowrecords"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func FromUsecase(borrowrecords borrowrecords.BorrowRecord) *borrowRecordsProto.BorrowRecordsResponse {
	return &borrowRecordsProto.BorrowRecordsResponse{
		Id:         int32(borrowrecords.Id),
		UserId:     int32(borrowrecords.UserId),
		BookId:     int32(borrowrecords.BookId),
		BorrowDate: timestamppb.New(borrowrecords.BorrowDate),
		ReturnDate: timestamppb.New(borrowrecords.ReturnDate),
		IsReturn:   borrowrecords.IsReturn,
		CreatedAt:  timestamppb.New(borrowrecords.CreatedAt),
		UpdatedAt:  timestamppb.New(borrowrecords.UpdatedAt),
	}
}

func FromUsecaseList(user []borrowrecords.BorrowRecord) []*borrowRecordsProto.BorrowRecordsResponse {
	var borrowRecordResponse []*borrowRecordsProto.BorrowRecordsResponse

	for _, v := range user {
		borrowRecordResponse = append(borrowRecordResponse, FromUsecase(v))
	}

	return borrowRecordResponse
}
