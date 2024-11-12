package controller

import (
	authors "authorservice/business/authors"
	"context"

	proto "shared/proto/authors"
	"shared/utils"

	"google.golang.org/grpc"
)

type AuthorService struct {
	usecase authors.AuthorUseCaseInterface
	proto.UnimplementedAuthorServiceServer
}

func NewServer(grpcServer *grpc.Server, usecase authors.AuthorUseCaseInterface) {
	bookGrpc := &AuthorService{usecase: usecase}
	proto.RegisterAuthorServiceServer(grpcServer, bookGrpc)
}

func NewAuthorService(uc authors.AuthorUseCaseInterface) *AuthorService {
	return &AuthorService{
		usecase: uc,
	}
}

func (s *AuthorService) AddAuthor(ctx context.Context, req *proto.AuthorRequest) (*proto.BaseResponseAuthor, error) {
	bookUseCase := ToUsecase(req)
	book, err := s.usecase.AddAuthor(*bookUseCase)
	if err != nil {
		return nil, utils.NewGrpcError(err)
	}

	return &proto.BaseResponseAuthor{
		Message: "Author successfully created",
		Author:  FromUsecase(book),
	}, nil
}

func (s *AuthorService) EditAuthor(ctx context.Context, req *proto.AuthorRequest) (*proto.BaseResponseAuthor, error) {
	bookId := req.GetId()
	bookEdit := ToUsecase(req)

	book, err := s.usecase.EditAuthor(*bookEdit, int(bookId))

	if err != nil {
		return nil, utils.NewGrpcError(err)
	}

	return &proto.BaseResponseAuthor{
		Message: "Author successfully updated",
		Author:  FromUsecase(book),
	}, nil
}

func (s *AuthorService) DeleteAuthor(ctx context.Context, req *proto.AuthorIdRequest) (*proto.BaseResponseAuthor, error) {
	bookId := req.Id

	book, err := s.usecase.DeleteAuthor(int(bookId))

	if err != nil {
		return nil, utils.NewGrpcError(err)
	}

	return &proto.BaseResponseAuthor{
		Message: "Author successfully deleted",
		Author:  FromUsecase(book),
	}, nil
}

func (s *AuthorService) GetAuthor(ctx context.Context, req *proto.AuthorIdRequest) (*proto.BaseResponseAuthor, error) {
	bookId := req.GetId()

	book, err := s.usecase.GetAuthor(int(bookId))

	if err != nil {
		return nil, utils.NewGrpcError(err)
	}

	return &proto.BaseResponseAuthor{
		Message: "Successfully retrieved book",
		Author:  FromUsecase(book),
	}, nil
}

func (s *AuthorService) GetAuthorExist(ctx context.Context, req *proto.AuthorIdRequest) (*proto.AuthorResponseExist, error) {
	bookId := req.GetId()

	_, err := s.usecase.GetAuthorExist(int(bookId))

	if err != nil {
		return &proto.AuthorResponseExist{Exists: false}, utils.NewGrpcError(err)
	}

	return &proto.AuthorResponseExist{Exists: true}, nil
}
