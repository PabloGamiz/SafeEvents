package client

// Client represents a client and it's main data
type Client interface {
	GetID() string
	GetName() string
	GetEmail() string
}
