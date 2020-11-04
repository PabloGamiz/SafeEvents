package product

import (
	"log"

	product "github.com/PabloGamiz/SafeEvents-Backend/model/product"
	mysql "github.com/PabloGamiz/SafeEvents-Backend/mysql"
)

// RepositoryImpl represents a
type RepositoryImpl struct {
	productRepository Repository
}

// AddProduct adds a product to the DB
func (productRepository *RepositoryImpl) AddProduct(product product.Product) (result interface{}, err error) {
	db, err := mysql.OpenStream()
	if err != nil {
		log.Printf("Got %v error while opening stream", err.Error())
		return
	}
	result = db.Create(&product)
	return
}
