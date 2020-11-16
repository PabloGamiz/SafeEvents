package service

import (
	"github.com/PabloGamiz/SafeEvents-Backend/model/location"
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
	GetKind() Kind
	SetKind(kind Kind)
	GetLocation() location.Location
	SetLocation(location location.Location)
	GetProducts() []product.Product
	SetProducts(products []product.Product)
}
