package client

// Controller represents a client and it's main data
type Controller interface {
	GetID() uint
	GetEmail() string
}
