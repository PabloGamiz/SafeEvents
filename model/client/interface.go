package client

// Controller represents a client and it's main data
type Controller interface {
	SetID(string)
	GetID() string
	GetName() string
	GetEmail() string
}
