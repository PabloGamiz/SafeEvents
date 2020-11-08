package dtos

// SigninRequestDTO represents the expected data from a Signin request
type SigninRequestDTO struct {
	TokenID string `json:"token_id"`
}

// SigninResponseDTO represents the provided data in front of a Signin response
type SigninResponseDTO struct {
	Cookie   string `json:"cookie"`
	Deadline int64  `json:"deadline"`
}
