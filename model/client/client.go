package client

type client struct {
	id       string
	name     string
	email    string
	password string
	verified bool
}

func (client *client) GetID() string {
	return client.id
}

func (client *client) GetName() string {
	return client.name
}

func (client *client) GetEmail() string {
	return client.email
}
