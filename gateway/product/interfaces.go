package product

import "github.com/PabloGamiz/SafeEvents-Backend/model/product"

// A Gateway represents the way between a model's object and the database
type Gateway interface {
	product.Controller
	Insert() error
}
