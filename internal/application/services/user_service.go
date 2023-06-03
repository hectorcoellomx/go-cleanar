package services

import (
	"github.com/hectorcoellomx/go-cleanar/internal/domain/entities"
	"github.com/hectorcoellomx/go-cleanar/internal/domain/ports"
)

type UserService struct {
	UserRepository ports.UserRepositoryPort
}

func NewUserService(userRepository ports.UserRepositoryPort) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (us *UserService) GetUsers() (*[]entities.User, error) {
	foundUser, err := us.UserRepository.FindAll()
	if err != nil {
		// Manejar el error de búsqueda del usuario por ID
		return nil, err
	}

	return foundUser, nil
}

func (us *UserService) GetUserByID(id uint) (*entities.User, error) {
	foundUser, err := us.UserRepository.FindByID(id)
	if err != nil {
		// Manejar el error de búsqueda del usuario por ID
		return nil, err
	}

	return foundUser, nil
}

func (us *UserService) CreateUser(username string, email string) (*entities.User, error) {
	newUser := &entities.User{
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
