package controller

import (
	booksStocks "bookservice/business/booksstocks"
	"context"

	proto "shared/proto/booksstocks"
	"shared/utils"

	"google.golang.org/grpc"
)

type BooksStocksService struct {
	usecase booksStocks.BooksStocksUseCaseInterface
	proto.UnimplementedBooksStocksServiceServer
}

func NewServer(grpcServer *grpc.Server, usecase booksStocks.BooksStocksUseCaseInterface) {
	bookGrpc := &BooksStocksService{usecase: usecase}
	proto.RegisterBooksStocksServiceServer(grpcServer, bookGrpc)
}

func NewBookService(uc booksStocks.BooksStocksUseCaseInterface) *BooksStocksService {
	return &BooksStocksService{
		usecase: uc,
	}
}

func (s *BooksStocksService) AddBooksStocks(ctx context.Context, req *proto.BooksStocksRequest) (*proto.BaseResponse, error) {
	booksStocksUseCase := ToUsecase(req)

	book, err := s.usecase.AddBooksStocks(*booksStocksUseCase)

	if err != nil {
		return nil, utils.NewGrpcError(err)
	}

	return &proto.BaseResponse{
		Message:     "BooksStocks successfully created",
		Booksstocks: FromUsecase(book),
	}, nil
}

func (s *BooksStocksService) EditBooksStocks(ctx context.Context, req *proto.BooksStocksRequest) (*proto.BaseResponse, error) {
	bookId := req.GetId()
	bookEdit := ToUsecase(req)

	book, err := s.usecase.EditBooksStocks(*bookEdit, int(bookId))

	if err != nil {
		return nil, utils.NewGrpcError(err)
	}

	return &proto.BaseResponse{
		Message:     "BooksStocks successfully updated",
		Booksstocks: FromUsecase(book),
	}, nil
}

func (s *BooksStocksService) DeleteBooksStocks(ctx context.Context, req *proto.BooksStocksIdRequest) (*proto.BaseResponse, error) {
	bookId := req.Id

	book, err := s.usecase.DeleteBooksStocks(int(bookId))

	if err != nil {
		return nil, utils.NewGrpcError(err)
	}

	return &proto.BaseResponse{
		Message:     "BooksStocks successfully deleted",
		Booksstocks: FromUsecase(book),
	}, nil
}

func (s *BooksStocksService) GetBooksStocks(ctx context.Context, req *proto.BooksStocksIdRequest) (*proto.BaseResponse, error) {
	bookId := req.GetId()

	book, err := s.usecase.GetBooksStocks(int(bookId))

	if err != nil {
		return nil, utils.NewGrpcError(err)
	}

	return &proto.BaseResponse{
		Message:     "Successfully retrieved book",
		Booksstocks: FromUsecase(book),
	}, nil
}
