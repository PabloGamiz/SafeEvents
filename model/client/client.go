package client

// Client its the main data object fro a client
type Client struct {
	ID    uint   `json:"id" gorm:"primaryKey; autoIncrement:true"`
	Email string `json:"email" gorm:"not null"`
}

// GetID return the id of the client
func (client *Client) GetID() uint {
	return client.ID
}

// GetEmail return the email of the client
func (client *Client) GetEmail() string {
	return client.Email
}
