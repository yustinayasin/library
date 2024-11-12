package controller

import (
	authors "authorservice/business/authors"
	authorsProto "shared/proto/authors"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func FromUsecase(book authors.Author) *authorsProto.AuthorResponse {
	return &authorsProto.AuthorResponse{
		Id:        int32(book.Id),
		Name:      book.Name,
		Bio:       book.Bio,
		Birthdate: timestamppb.New(book.Birthdate),
		CreatedAt: timestamppb.New(book.CreatedAt),
		UpdatedAt: timestamppb.New(book.UpdatedAt),
	}
}

func FromUsecaseList(book []authors.Author) []*authorsProto.AuthorResponse {
	var bookResponse []*authorsProto.AuthorResponse

	for _, v := range book {
		bookResponse = append(bookResponse, FromUsecase(v))
	}

	return bookResponse
}
