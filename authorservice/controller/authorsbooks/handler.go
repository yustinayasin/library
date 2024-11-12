package controller

import (
	authorsBooks "authorservice/business/authorsbooks"
	"context"

	proto "shared/proto/authors"
	"shared/utils"

	"google.golang.org/grpc"
)

type AuthorBooksService struct {
	usecase authorsBooks.AuthorsBooksUseCaseInterface
	proto.UnimplementedAuthorsBooksServiceServer
}

func NewServer(grpcServer *grpc.Server, usecase authorsBooks.AuthorsBooksUseCaseInterface) {
	bookGrpc := &AuthorBooksService{usecase: usecase}
	proto.RegisterAuthorsBooksServiceServer(grpcServer, bookGrpc)
}

func NewBookService(uc authorsBooks.AuthorsBooksUseCaseInterface) *AuthorBooksService {
	return &AuthorBooksService{
		usecase: uc,
	}
}

func (s *AuthorBooksService) AddAuthorBook(ctx context.Context, req *proto.AuthorsBooksRequest) (*proto.BaseResponseAuthorBooks, error) {
	authorsBooksUseCase := ToUsecase(req)

	authorbook, err := s.usecase.AddAuthorBook(*authorsBooksUseCase)

	if err != nil {
		return nil, utils.NewGrpcError(err)
	}

	return &proto.BaseResponseAuthorBooks{
		Message:      "AuthorBooks successfully created",
		Authorsbooks: FromUsecase(authorbook),
	}, nil
}

func (s *AuthorBooksService) EditAuthorBook(ctx context.Context, req *proto.AuthorsBooksRequest) (*proto.BaseResponseAuthorBooks, error) {
	bookId := req.GetBookId()
	authorId := req.GetAuthorId()
	bookEdit := ToUsecase(req)

	authorbook, err := s.usecase.EditAuthorBook(*bookEdit, int(bookId))

	if err != nil {
		return nil, utils.NewGrpcError(err)
	}

	return &proto.BaseResponseAuthorBooks{
		Message:      "AuthorBooks successfully updated",
		Authorsbooks: FromUsecase(authorbook),
	}, nil
}

func (s *AuthorBooksService) DeleteADeleteAuthorBookuthorBooks(ctx context.Context, req *proto.AuthorsBooksRequest) (*proto.BaseResponseAuthorBooks, error) {
	bookId := req.Id

	authorbook, err := s.usecase.DeleteAuthorBook(int(bookId))

	if err != nil {
		return nil, utils.NewGrpcError(err)
	}

	return &proto.BaseResponseAuthorBooks{
		Message:      "AuthorBooks successfully deleted",
		Authorsbooks: FromUsecase(authorbook),
	}, nil
}

func (s *AuthorBooksService) GetAuthorBooks(ctx context.Context, req *proto.AuthorsBooksRequest) (*proto.BaseResponseAuthorBooks, error) {
	bookId := req.GetId()

	authorbook, err := s.usecase.GetAuthorBook(int(bookId))

	if err != nil {
		return nil, utils.NewGrpcError(err)
	}

	return &proto.BaseResponseAuthorBooks{
		Message:      "Successfully retrieved authorbook",
		Authorsbooks: FromUsecase(authorbook),
	}, nil
}
