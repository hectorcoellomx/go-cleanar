package services

import (
	"github.com/hectorcoellomx/go-cleanar/internal/domain/user"
)

type UserService struct {
	UserRepository user.UserRepositoryPort
}

func NewUserService(userRepository user.UserRepositoryPort) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (us *UserService) GetUsers() (*[]user.User, error) {
	foundUser, err := us.UserRepository.FindAll()
	if err != nil {
		// Manejar el error de búsqueda del usuario por ID
		return nil, err
	}

	return foundUser, nil
}

func (us *UserService) GetUserByID(id uint) (*user.User, error) {
	foundUser, err := us.UserRepository.FindByID(id)
	if err != nil {
		// Manejar el error de búsqueda del usuario por ID
		return nil, err
	}

	return foundUser, nil
}

func (us *UserService) CreateUser(username string, email string) (*user.User, error) {
	newUser := &user.User{
		Username: username,
		Email:    email,
	}

	err := us.UserRepository.Create(newUser)
	if err != nil {
		// Manejar el error de creación del usuario
		return nil, err
	}

	return newUser, nil
}
