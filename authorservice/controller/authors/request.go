package controller

import (
	authors "authorservice/business/authors"
	authorProto "shared/proto/authors"
)

func ToUsecase(author *authorProto.AuthorRequest) *authors.Author {
	return &authors.Author{
		Id:        int(author.Id),
		Name:      author.Name,
		Bio:       author.Bio,
		Birthdate: author.Birthdate.AsTime(),
	}
}
