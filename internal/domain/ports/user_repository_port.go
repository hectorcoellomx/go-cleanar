package ports

import (
	"github.com/hectorcoellomx/go-cleanar/internal/domain/entities"
)

type UserRepositoryPort interface {
	FindAll() (*[]entities.User, error)
	Create(user *entities.User) error
	FindByID(id uint) (*entities.User, error)
}
