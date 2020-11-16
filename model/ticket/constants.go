package ticket

// Option represents a purchase option
type Option int

// Option values
const (
	BOOKED Option = iota
	BOUGHT
	CHECKED
)
