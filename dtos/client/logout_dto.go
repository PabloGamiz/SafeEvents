package dtos

// LogoutRequestDTO represents the expected data from a Signin request
type LogoutRequestDTO struct {
	Cookie string `json:"cookie"`
}

// LogoutResponseDTO represents the provided data in front of a Signin response
type LogoutResponseDTO struct {
	Cookie   string `json:"cookie"`
	Deadline int64  `json:"deadline"`
}
