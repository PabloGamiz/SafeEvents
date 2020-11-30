package ticket

// Option represents a purchase option
type Option int

// Option values
const (
	BOOKED Option = iota
	BOUGHT
	CHECKED

	queryFindByEventID            = "event_id = ?"
	queryFindByEventIDAndClientID = "event_id = ? and client_id = ?"
)
