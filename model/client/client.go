package client

// Client its the main data object fro a client
type Client struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Name     string `json:"name" bson:"_name,omitempty"`
	Email    string `json:"email" bson:"_email,omitempty"`
	Password string `json:"password" bson:"_password,omitempty"`
	Verified bool   `json:"verified" bson:"_id,omitempty"`
}

// SetID sets a new id to the model
func (client *Client) SetID(id string) {
	client.ID = id
}

// GetID return the id of the client
func (client *Client) GetID() string {
	return client.ID
}

// GetName returns the name of the client
func (client *Client) GetName() string {
	return client.Name
}

// GetEmail return the email of the client
func (client *Client) GetEmail() string {
	return client.Email
}
