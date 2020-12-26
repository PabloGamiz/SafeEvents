package radar

import "time"

// ActivateRequestDTO represents the expected data from a activate request
type ActivateRequestDTO struct {
	Cookie string `json:"cookie"`
	MAC    string `json:"mac"`
}

// ActivateResponseDTO represents the expected data from a activate request
type ActivateResponseDTO struct {
	StartAt time.Time `json:"start_at"`
}
