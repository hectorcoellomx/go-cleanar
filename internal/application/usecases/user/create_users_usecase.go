package user

import (
	"github.com/hectorcoellomx/go-cleanar/internal/application/services"
	"github.com/hectorcoellomx/go-cleanar/internal/domain/entities"
)

type CreateUser struct {
	userService services.UserService
}

func NewCreateUsers(userService services.UserService) *CreateUser {
	return &CreateUser{
		userService: userService,
	}
}

func (uc *CreateUser) CreateUser(id int, name string, email string, password string, status int) (*entities.User, error) {
	return uc.userService.CreateUser(id, name, email, password, status)
}
