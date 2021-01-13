package radar

import "time"

// InteractionRequestDTO represents the expected data from a activate request
type InteractionRequestDTO struct {
	Cookie  string   `json:"cookie"`
	CloseTo []string `json:"close_to"`
	Instant int64    `json:"instant"`
	Unix    time.Time
}

// InteractionResponseDTO represents the expected data from a activate request
type InteractionResponseDTO struct {
	HowMany uint `json:"how_many"`
}
