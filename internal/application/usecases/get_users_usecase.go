package usecases

import (
	"github.com/hectorcoellomx/go-cleanar/internal/application/services"
	"github.com/hectorcoellomx/go-cleanar/internal/domain/entities"
)

type GetUsers struct {
	userService services.UserService
}

func NewGetUsers(userService services.UserService) *GetUsers {
	return &GetUsers{
		userService: userService,
	}
}

func (uc *GetUsers) GetUsers() (*[]entities.User, error) {
	return uc.userService.GetUsers()
}
