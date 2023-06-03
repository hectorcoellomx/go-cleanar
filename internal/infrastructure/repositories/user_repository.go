package repositories

import (
	"github.com/hectorcoellomx/go-cleanar/internal/domain/entities"
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

func (ur *UserRepository) FindAll() (*[]entities.User, error) {

	var users []entities.User
	err := ur.DB.Find(&users).Error
	if err != nil {
		// Manejar el error de creación del usuario
		return nil, err
	}

	return &users, nil
}

func (ur *UserRepository) Create(user *entities.User) error {
	err := ur.DB.Create(user).Error
	if err != nil {
		// Manejar el error de creación del usuario
		return err
	}

	return nil
}

func (ur *UserRepository) FindByID(id uint) (*entities.User, error) {
	var user entities.User
	err := ur.DB.First(&user, id).Error
	if err != nil {
		// Manejar el error de búsqueda del usuario por ID
		return nil, err
	}

	return &user, nil
}
