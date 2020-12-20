package product

import (
	"time"
)

// Product represents the product class from UML
type Product struct {
	ID          uint   `json:"id" gorm:"primaryKey; autoIncrement:true"`
	Name        string `json:"name" gorm:"not null;unique"`
	Description string `json:"description" gorm:"not null"`
	Price       int    `json:"price" gorm:"not null"`
	Status      string `json:"status" gorm:"not null"`
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
func (product *Product) GetStatus() string {
	return product.Status
}

// SetStatus sets the Status of the Product.
func (product *Product) SetStatus(status string) {
	product.Status = status
}
