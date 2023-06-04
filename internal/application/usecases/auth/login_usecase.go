package auth

import (
	"github.com/hectorcoellomx/go-cleanar/internal/application/services"
	"github.com/hectorcoellomx/go-cleanar/internal/domain/entities"
)

type Login struct {
	userService services.UserService
}

func NewLogin(userService services.UserService) *Login {
	return &Login{
		userService: userService,
	}
}

func (uc *Login) Login() (*[]entities.User, error) {
	return uc.userService.GetUsers()
}
