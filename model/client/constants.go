package client

// Status represents the ClientStatus enum
type Status int

// Status possible values
const (
	NEGATIVE Status = iota
	POSITIVE
	UNKNOWN
)

const (
	errAssertionFailed = "Assertion has failed, expecting *Client type"
	errClientNotExists = "Client for %v does not exists"

	errNotFoundByID    = "Got %v, while looking up for client with ID %v"
	errNotFoundByEmail = "Got %v while looking up for client with email %v"

	queryFindByEmail = "email = ?"
	queryFindByID    = "id = ?"
)
