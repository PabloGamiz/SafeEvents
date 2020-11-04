package service

import (
	"time"

	"github.com/PabloGamiz/SafeEvents-Backend/model/location"
	"github.com/PabloGamiz/SafeEvents-Backend/model/product"
	"gorm.io/gorm"
)

// Service represents the product class from UML
type Service struct {
	gorm.Model
	ID          uint              `json:"id" gorm:"primaryKey; autoIncrement:true"`
	Name        string            `json:"name" gorm:"not null;unique"`
	Description string            `json:"description" gorm:"not null"`
	Kind        Kind              `json:"kind" gorm:"not null"`
	Location    location.Location `json:"location" gorm:"foreignkey:LocationID;not null"`
	LocationID  uint64            `json:"-"`
	Products    []product.Product `json:"products" gorm:"many2many:services_products"`
	EventID     uint              `json:"-"`
	CreatedAt   time.Time         `json:"createdAt"`
	UpdatedAt   time.Time         `json:"updatedAt"`
}

// GetID return the ID of the Service.
func (service *Service) GetID() uint {
	return service.ID
}

// SetID sets a new id to the model
func (service *Service) SetID(id uint) {
	service.ID = id
}

// GetName return the Name of the Service.
func (service *Service) GetName() string {
	return service.Name
}

// SetName sets the Name of the Service.
func (service *Service) SetName(name string) {
	service.Name = name
}

// GetDescription return the Description of the Service.
func (service *Service) GetDescription() string {
	return service.Description
}

// SetDescription sets the Description of the Service.
func (service *Service) SetDescription(description string) {
	service.Description = description
}

// GetKind return the Kind of the Service.
func (service *Service) GetKind() string {
	return service.Kind.String()
}

// SetKind sets the Kind of the Service.
func (service *Service) SetKind(kind Kind) {
	service.Kind = kind
}

// GetLocation return the Location of the Service.
func (service *Service) GetLocation() (loc location.Location) {
	return service.Location
}

// SetLocation sets the Description of the Service.
func (service *Service) SetLocation(location location.Location) {
	service.Location = location
}

// GetProducts return the Products of the Service.
func (service *Service) GetProducts() []product.Product {
	return service.Products
}

// SetProducts sets the Description of the Service.
func (service *Service) SetProducts(products []product.Product) {
	service.Products = products
}
