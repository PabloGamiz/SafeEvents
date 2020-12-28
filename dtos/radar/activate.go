package radar

import "time"

// ActivateRequestDTO represents the expected data from a activate request
type ActivateRequestDTO struct {
	Cookie string `json:"cookie"`
	MAC    string `json:"mac"`
}

// ActivateResponseDTO represents the expected data from a activate request
type ActivateResponseDTO struct {
	DoneAt time.Time `json:"done_at"`
}
