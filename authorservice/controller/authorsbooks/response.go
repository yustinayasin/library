package controller

import (
	authorsbooks "authorservice/business/authorsbooks"
	authorProto "shared/proto/authors"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func FromUsecase(authorbook authorsbooks.AuthorBook) *authorProto.AuthorsBooksResponse {
	return &authorProto.AuthorsBooksResponse{
		BookId:    int32(authorbook.BookId),
		AuthorId:  int32(authorbook.AuthorId),
		CreatedAt: timestamppb.New(authorbook.CreatedAt),
		UpdatedAt: timestamppb.New(authorbook.UpdatedAt),
	}
}

func FromUsecaseList(authorbook []authorsbooks.AuthorBook) []*authorProto.AuthorsBooksResponse {
	var bookResponse []*authorProto.AuthorsBooksResponse

	for _, v := range authorbook {
		bookResponse = append(bookResponse, FromUsecase(v))
	}

	return bookResponse
}
