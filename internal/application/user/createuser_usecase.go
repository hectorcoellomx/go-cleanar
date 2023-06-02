package user

import (
	"github.com/hectorcoellomx/go-cleanar/internal/domain/user"
)

type CreateUser struct {
	userService UserService
}

func NewCreateUser(userService UserService) *CreateUser {
	return &CreateUser{
		userService: userService,
	}
}

func (uc *CreateUser) Execute(name string, email string) (*user.User, error) {
	return uc.userService.CreateUser(name, email)
}
