package location

import "go.mongodb.org/mongo-driver/bson/primitive"

// Controller represents a Location and it's main data
type Controller interface {
	GetID() primitive.ObjectID
	SetID(id primitive.ObjectID)
	GetName() string
	SetName(name string)
	GetAddress() string
	SetAddress(address string)
	GetCoordinates() string
	SetCoordinates(coordinates string)
	GetExtension() int
	SetExtension(extension int)
}
