package product

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Controller represents a Product and it's main data
type Controller interface {
	GetID() primitive.ObjectID
	SetID(id primitive.ObjectID)
	GetName() string
	SetName(name string)
	GetDescription() string
	SetDescription(description string)
	GetPrice() int
	SetPrice(price int)
	GetStatus() Status
	SetStatus(status Status)
}
