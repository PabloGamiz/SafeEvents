package client

import "go.mongodb.org/mongo-driver/bson/primitive"

// Client its the main data object fro a client
type Client struct {
	ID    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email string             `json:"email" bson:"email"`
}

// SetID sets a new id to the model
func (client *Client) SetID(id primitive.ObjectID) {
	client.ID = id
}

// GetID return the id of the client
func (client *Client) GetID() *primitive.ObjectID {
	return &client.ID
}

// GetEmail return the email of the client
func (client *Client) GetEmail() string {
	return client.Email
}
