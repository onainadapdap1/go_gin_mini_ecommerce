package repository

import (
	"go_gin_mini_ecommerce/models"

	"github.com/jinzhu/gorm"
)

type ProductRepository interface {
	GetProduct(int) (models.Product, error)
	AddProduct(models.Product) (models.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository() ProductRepository {
	return &productRepository{db: DB()}
}

func(r *productRepository) GetProduct(id int) (product models.Product, err error) {
	return product, r.db.First(&product, id).Error
}

func (r *productRepository) AddProduct(product models.Product) (models.Product, error) {
	return product, r.db.Create(&product).Error
}