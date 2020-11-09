package dtos

// ClientInfoRequestDTO represents the expected data from a Client Info request
type ClientInfoRequestDTO struct {
	Email string 'json:"email"'
}

// ClientInfoResponseDTO represents the provided data in front of a Client Info response
type ClientInfoResponse struct {
	Username string 'json:"username"'
	Email string 'json:"email"'
	Verified string 'json:"verified"'
	Events map[string]string 'json:"events"'
}