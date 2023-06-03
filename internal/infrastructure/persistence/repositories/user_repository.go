package repositories

import (
	"github.com/hectorcoellomx/go-cleanar/internal/domain/user"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (ur *UserRepository) FindAll() (*[]user.User, error) {

	var users []user.User
	err := ur.DB.Find(&users).Error
	if err != nil {
		// Manejar el error de creación del usuario
		return nil, err
	}

	return &users, nil
}

func (ur *UserRepository) Create(user *user.User) error {
	err := ur.DB.Create(user).Error
	if err != nil {
		// Manejar el error de creación del usuario
		return err
	}

	return nil
}

func (ur *UserRepository) FindByID(id uint) (*user.User, error) {
	var user user.User
	err := ur.DB.First(&user, id).Error
	if err != nil {
		// Manejar el error de búsqueda del usuario por ID
		return nil, err
	}

	return &user, nil
}
