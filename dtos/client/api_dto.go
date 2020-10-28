package client

// DTO represents the expected data from a Client.
type DTO struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}
