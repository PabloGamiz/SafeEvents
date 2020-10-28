package service

import (
	"github.com/PabloGamiz/SafeEvents-Backend/model/location"
	"github.com/PabloGamiz/SafeEvents-Backend/model/product"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Service represents the product class from UML
type Service struct {
	ID          primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Name        string               `json:"name" bson:"name,omitempty"`
	Description string               `json:"description" bson:"description,omitempty"`
	Kind        Kind                 `json:"kind" bson:"kind,omitempty"`
	Location    location.Controller  `json:"location"`
	Products    []product.Controller `json:"product"`
}

// GetID return the ID of the Service.
func (service *Service) GetID() string {
	return service.ID.String()
}

// SetID sets a new id to the model
func (service *Service) SetID(id primitive.ObjectID) {
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
func (service *Service) GetKind() Kind {
	return service.Kind
}

// SetKind sets the Kind of the Service.
func (service *Service) SetKind(kind Kind) {
	service.Kind = kind
}

// GetLocation return the Location of the Service.
func (service *Service) GetLocation() (loc location.Controller) {
	return service.Location
}

// GetProducts return the Products of the Service.
func (service *Service) GetProducts() []product.Controller {
	return service.Products
}

// SetProducts sets the Description of the Service.
func (service *Service) SetProducts(products []product.Controller) {
	service.Products = products
}
