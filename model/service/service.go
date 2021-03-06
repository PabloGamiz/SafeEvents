package service

import (
	"time"

	"github.com/PabloGamiz/SafeEvents-Backend/model/product"
)

// Service represents the product class from UML
type Service struct {
	ID          uint              `json:"id" gorm:"primaryKey; autoIncrement:true"`
	Name        string            `json:"name" gorm:"not null;unique"`
	Description string            `json:"description" gorm:"not null"`
	Kind        string            `json:"kind" gorm:"not null"`
	Location    string            `json:"location" gorm:"not null"`
	Products    []product.Product `json:"products" gorm:"foreignkey:ServiceID;constraint:OnDelete:CASCADE"`
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
	return service.Kind
}

// SetKind sets the Kind of the Service.
func (service *Service) SetKind(kind string) {
	service.Kind = kind
}

// GetLocation return the Location of the Service.
func (service *Service) GetLocation() string {
	return service.Location
}

// SetLocation sets the Description of the Service.
func (service *Service) SetLocation(location string) {
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

// // GetService gets a pointer to this Service.
// func (service *Service) GetService() *Service {
// 	return &Service{}
// }
