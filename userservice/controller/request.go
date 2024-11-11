package controller

import (
	usersProto "shared/proto/users"
	users "userservice/business"
)

func ToUsecaseLogin(user *usersProto.UserLoginRequest) *users.User {
	return &users.User{
		Email:    user.Email,
		Password: user.Password,
	}
}

func ToUsecaseEdit(user *usersProto.UserEditRequest) *users.User {
	return &users.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}
