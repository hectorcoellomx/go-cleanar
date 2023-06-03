package ports

import (
	"github.com/hectorcoellomx/go-cleanar/internal/domain/user"
)

type UserRepositoryPort interface {
	FindAll() (*[]user.User, error)
	Create(user *user.User) error
	FindByID(id uint) (*user.User, error)
}
