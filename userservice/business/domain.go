package users

import (
	"time"
)

type Role struct {
	Id   int
	Name string
}

type User struct {
	Id        int
	Name      string
	Email     string
	Password  string
	Token     string
	RoleId    int
	Role      Role
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserUseCaseInterface interface {
	SignUp(user User) (User, error)
	Login(user User) (User, error)
	EditUser(user User, id int) (User, error)
	DeleteUser(id int) (User, error)
	GetUser(id int) (User, error)
}

type UserRepoInterface interface {
	SignUp(user User) (User, error)
	Login(user User) (User, error)
	EditUser(user User, id int) (User, error)
	DeleteUser(id int) (User, error)
	GetUser(id int) (User, error)
}
