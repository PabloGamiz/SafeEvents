package service

import (
	"github.com/PabloGamiz/SafeEvents-Backend/model/location"
	"github.com/PabloGamiz/SafeEvents-Backend/model/product"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Controller represents a Service and it's main data
type Controller interface {
	GetID() primitive.ObjectID
	SetID(id primitive.ObjectID)
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
