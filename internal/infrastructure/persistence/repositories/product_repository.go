package repositories

import (
	"github.com/hectorcoellomx/go-cleanar/internal/domain/product"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(newProduct *product.Product) error {
	// Lógica para crear un producto en la base de datos utilizando GORM
}

func (r *ProductRepository) FindByID(productID int) (*product.Product, error) {
	// Lógica para buscar un producto por ID en la base de datos utilizando GORM
}
