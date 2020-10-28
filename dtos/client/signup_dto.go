package client

// SignupRequestDTO represents the expected data from a Signup request
type SignupRequestDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

// SignupResponseDTO represents the provided data in front of a Signup response
type SignupResponseDTO struct {
	Cookie   string `json:"cookie"`
	Deadline uint64 `json:"timeout"`
}
