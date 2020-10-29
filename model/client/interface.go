package client

import "go.mongodb.org/mongo-driver/bson/primitive"

// Controller represents a client and it's main data
type Controller interface {
	SetID(primitive.ObjectID)
	GetID() *primitive.ObjectID
	GetEmail() string
}
