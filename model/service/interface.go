package service

import (
	"github.com/PabloGamiz/SafeEvents-Backend/model/product"
)

// Controller represents a Service and it's main data
type Controller interface {
	GetID() uint
	SetID(id uint)
	GetName() string
	SetName(name string)
	GetDescription() string
	SetDescription(description string)
	GetKind() string
	// SetKind(kind Kind)
	GetLocation() string
	SetLocation(location string)
	GetProducts() []product.Product
	SetProducts(products []product.Product)
	// GetService() *Service
}
