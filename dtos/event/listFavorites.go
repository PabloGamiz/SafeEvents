package event

// ListFavoritesRequestDTO represents the expected data from a List Favorites request
type ListFavoritesRequestDTO struct {
	ID     uint   `json:"id"`
	Cookie string `json:"cookie"`
}
