package users

import (
	"time"
	users "userservice/business"
)

type Role struct {
	Id   int `gorm:"primaryKey;unique"`
	Name string
}

func (role Role) ToUseCase() users.Role {
	return users.Role{
		Id:   role.Id,
		Name: role.Name,
	}
}

func FromUsecaseRole(role users.Role) Role {
	return Role{
		Id:   role.Id,
		Name: role.Name,
	}
}

type User struct {
	Id          int `gorm:"primaryKey;unique"`
	Name        string
	Email       string `gorm:"unique"`
	Password    string
	Address     string
	PhoneNumber string
	PostalCode  string
	RoleId      int
	Role        Role `gorm:"foreignKey:RoleId"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (user User) ToUsecase() users.User {
	newRole := Role.ToUseCase(user.Role)

	return users.User{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		RoleId:    user.RoleId,
		Role:      newRole,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUsecaseList(user []User) []users.User {
	var newUsers []users.User

	for _, v := range user {
		newUsers = append(newUsers, v.ToUsecase())
	}
	return newUsers
}

func FromUsecase(user users.User) User {
	newRole := FromUsecaseRole(user.Role)

	return User{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		RoleId:    user.RoleId,
		Role:      newRole,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
