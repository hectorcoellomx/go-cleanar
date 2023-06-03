package usecases

import (
	"github.com/hectorcoellomx/go-cleanar/internal/application/services"
	"github.com/hectorcoellomx/go-cleanar/internal/domain/user"
)

type GetUsers struct {
	userService services.UserService
}

func NewGetUsers(userService services.UserService) *GetUsers {
	return &GetUsers{
		userService: userService,
	}
}

func (uc *GetUsers) Execute() (*[]user.User, error) {
	return uc.userService.GetUsers()
}
