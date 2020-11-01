package product

import (
	"time"
)

// Product represents the product class from UML
type Product struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Status      Status `json:"status"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// GetID return the ID of the Product.
func (product *Product) GetID() uint {
	return product.ID
}

// GetName return the Name of the Product.
func (product *Product) GetName() string {
	return product.Name
}

// SetName sets the Name of the Product.
func (product *Product) SetName(name string) {
	product.Name = name
}

// GetDescription return the Description of the Product.
func (product *Product) GetDescription() string {
	return product.Description
}

// SetDescription sets the Description of the Product.
func (product *Product) SetDescription(description string) {
	product.Description = description
}

// GetPrice return the Price of the Product.
func (product *Product) GetPrice() int {
	return product.Price
}

// SetPrice sets the Price of the Product.
func (product *Product) SetPrice(price int) {
	product.Price = price
}

// GetStatus return the Status of the Product.
func (product *Product) GetStatus() Status {
	return product.Status
}

// SetStatus sets the Status of the Product.
func (product *Product) SetStatus(status Status) {
	product.Status = status
}
