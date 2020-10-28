package dtos

// ClientDTO represents the expected data from a Client.
type ClientDTO struct {
	ID              int        `json:"id"`
	Email           string     `json:"email"`
	FavouriteEvents []EventDTO `json:"favouriteEvents"`
}
