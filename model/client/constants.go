package client

const (
	errAssertionFailed = "Assertion has failed, expecting *Client type"
	errClientNotExists = "Client for %d does not exists"

	errNotFoundByID    = "Not found client with ID %v"
	errNotFoundByEmail = "Not found client with email %s"

	queryFindByEmail = "email = ?"
	queryFindByID    = "ID = ?"
)
