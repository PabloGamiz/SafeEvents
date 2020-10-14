package client

// NewClient builds a bran new client for the provided data
func NewClient(id, name, email string) Client {
	return &client{
		id:    id,
		name:  name,
		email: email,
	}
}
