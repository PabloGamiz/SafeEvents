package product

import (
	product "github.com/PabloGamiz/SafeEvents-Backend/model/product"
	factory "github.com/PabloGamiz/SafeEvents-Backend/mysql"
)

// Repository represents a
type Repository interface {
	AddProduct(product product.Product)
}

// RepositoryImpl represents a
type RepositoryImpl struct {
	productRepository Repository
}

// AddProduct adds a product to the DB
func (productRepository *RepositoryImpl) AddProduct(product product.Product) (result interface{}, err error) {
	db, err := factory.OpenStream()
	result = db.Create(&product)
	return
}
