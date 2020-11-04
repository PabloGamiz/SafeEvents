package product

import (
	product "github.com/PabloGamiz/SafeEvents-Backend/model/product"
)

// Repository represents a
type Repository interface {
	AddProduct(product product.Product, err error)
}
