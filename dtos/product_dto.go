package dtos

// ProductDTO represents the expected data from a Product.
type ProductDTO struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Status      string `json:"status"`
}
