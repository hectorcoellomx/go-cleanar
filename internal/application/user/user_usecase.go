package user

import (
	"github.com/hectorcoellomx/go-cleanar/internal/domain/user"
)

type UserUseCase interface {
	CreateUser(name string, email string) (*user.User, error)
	GetUserByID(id uint) (*user.User, error)
}

type userUseCase struct {
	userService UserService
}

func NewUserUseCase(userService UserService) UserUseCase {
	return &userUseCase{
		userService: userService,
	}
}

func (uc *userUseCase) CreateUser(name string, email string) (*user.User, error) {
	return uc.userService.CreateUser(name, email)
}

func (uc *userUseCase) GetUserByID(id uint) (*user.User, error) {
	return uc.userService.GetUserByID(id)
}
