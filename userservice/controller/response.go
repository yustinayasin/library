package controller

import (
	usersProto "shared/proto/users"
	users "userservice/business"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func FromUsecase(user users.User) *usersProto.UserResponse {
	return &usersProto.UserResponse{
		Id:        int32(user.Id),
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Token:     user.Token,
		RoleId:    int32(user.RoleId),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}

func FromUsecaseList(user []users.User) []*usersProto.UserResponse {
	var userResponse []*usersProto.UserResponse

	for _, v := range user {
		userResponse = append(userResponse, FromUsecase(v))
	}

	return userResponse
}
