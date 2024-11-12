package controller

import (
	books "bookservice/business/books"
	"context"

	proto "shared/proto/books"
	"shared/utils"

	"google.golang.org/grpc"
)

type BookService struct {
	usecase books.BookUseCaseInterface
	proto.UnimplementedBookServiceServer
}

func NewServer(grpcServer *grpc.Server, usecase books.BookUseCaseInterface) {
	bookGrpc := &BookService{usecase: usecase}
	proto.RegisterBookServiceServer(grpcServer, bookGrpc)
}

func NewBookService(uc books.BookUseCaseInterface) *BookService {
	return &BookService{
		usecase: uc,
	}
}

func (s *BookService) AddBook(ctx context.Context, req *proto.BookRequest) (*proto.BaseResponseBook, error) {
	bookUseCase := ToUsecase(req)
	book, err := s.usecase.AddBook(*bookUseCase)
	if err != nil {
		return nil, utils.NewGrpcError(err)
	}

	return &proto.BaseResponseBook{
		Message: "Book successfully created",
		Book:    FromUsecase(book),
	}, nil
}

func (s *BookService) EditBook(ctx context.Context, req *proto.BookRequest) (*proto.BaseResponseBook, error) {
	bookId := req.GetId()
	bookEdit := ToUsecase(req)

	book, err := s.usecase.EditBook(*bookEdit, int(bookId))

	if err != nil {
		return nil, utils.NewGrpcError(err)
	}

	return &proto.BaseResponseBook{
		Message: "Book successfully updated",
		Book:    FromUsecase(book),
	}, nil
}

func (s *BookService) DeleteBook(ctx context.Context, req *proto.BookIdRequest) (*proto.BaseResponseBook, error) {
	bookId := req.Id

	book, err := s.usecase.DeleteBook(int(bookId))

	if err != nil {
		return nil, utils.NewGrpcError(err)
	}

	return &proto.BaseResponseBook{
		Message: "Book successfully deleted",
		Book:    FromUsecase(book),
	}, nil
}

func (s *BookService) GetBook(ctx context.Context, req *proto.BookIdRequest) (*proto.BaseResponseBook, error) {
	bookId := req.GetId()

	book, err := s.usecase.GetBook(int(bookId))

	if err != nil {
		return nil, utils.NewGrpcError(err)
	}

	return &proto.BaseResponseBook{
		Message: "Successfully retrieved book",
		Book:    FromUsecase(book),
	}, nil
}

func (s *BookService) GetBookExist(ctx context.Context, req *proto.BookIdRequest) (*proto.BookResponseExist, error) {
	bookId := req.GetId()

	_, err := s.usecase.GetBookExist(int(bookId))

	if err != nil {
		return &proto.BookResponseExist{Exists: false}, utils.NewGrpcError(err)
	}

	return &proto.BookResponseExist{Exists: true}, nil
}
