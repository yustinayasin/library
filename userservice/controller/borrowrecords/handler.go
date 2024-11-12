package controller

import (
	"context"
	borrowrecords "userservice/business/borrowrecords"

	proto "shared/proto/users"
	"shared/utils"

	"google.golang.org/grpc"
)

type BorrowRecordService struct {
	usecase borrowrecords.BorrowRecordUseCaseInterface
	proto.UnimplementedBorrowRecordsServiceServer
}

func NewServer(grpcServer *grpc.Server, usecase borrowrecords.BorrowRecordUseCaseInterface) {
	userGrpc := &BorrowRecordService{usecase: usecase}
	proto.RegisterBorrowRecordsServiceServer(grpcServer, userGrpc)
}

func NewBorrowRecordService(uc borrowrecords.BorrowRecordUseCaseInterface) *BorrowRecordService {
	return &BorrowRecordService{
		usecase: uc,
	}
}

func (s *BorrowRecordService) AddBorrowRecord(ctx context.Context, req *proto.BorrowRecordsRequest) (*proto.BaseResponseBorrowRecords, error) {
	userUseCase := ToUsecase(req)
	borrowRecord, err := s.usecase.AddBorrowRecord(*userUseCase)
	if err != nil {
		return nil, utils.NewGrpcError(err)
	}

	return &proto.BaseResponseBorrowRecords{
		Message:       "Borrow record successfully created",
		Borrowrecords: FromUsecase(borrowRecord),
	}, nil
}

func (s *BorrowRecordService) EditBorrowRecord(ctx context.Context, req *proto.BorrowRecordsRequest) (*proto.BaseResponseBorrowRecords, error) {
	userId := req.GetId()
	userEdit := ToUsecase(req)
	borrowRecord, err := s.usecase.EditBorrowRecord(*userEdit, int(userId))
	if err != nil {
		return nil, utils.NewGrpcError(err)
	}

	return &proto.BaseResponseBorrowRecords{
		Message:       "Borrow record successfully updated",
		Borrowrecords: FromUsecase(borrowRecord),
	}, nil
}

func (s *BorrowRecordService) DeleteBorrowRecord(ctx context.Context, req *proto.BorrowRecordsRequest) (*proto.BaseResponseBorrowRecords, error) {
	userId := req.Id
	borrowRecord, err := s.usecase.DeleteBorrowRecord(int(userId))
	if err != nil {
		return nil, utils.NewGrpcError(err)
	}

	return &proto.BaseResponseBorrowRecords{
		Message:       "Borrow record successfully deleted",
		Borrowrecords: FromUsecase(borrowRecord),
	}, nil
}

func (s *BorrowRecordService) GetBorrowRecord(ctx context.Context, req *proto.BorrowRecordsRequest) (*proto.BaseResponseBorrowRecords, error) {
	userId := req.GetId()
	borrowRecord, err := s.usecase.GetBorrowRecord(int(userId))
	if err != nil {
		return nil, utils.NewGrpcError(err)
	}

	return &proto.BaseResponseBorrowRecords{
		Message:       "Successfully retrieved borrow record",
		Borrowrecords: FromUsecase(borrowRecord),
	}, nil
}
