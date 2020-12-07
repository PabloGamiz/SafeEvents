package client

// ClientInfoRequestDTO represents the expected data from a Client Info request
type ClientInfoRequestDTO struct {
	ID     uint   `json:"id"`
	Cookie string `json:"cookie"`
}
