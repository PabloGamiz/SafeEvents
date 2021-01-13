package ticket

const (
	errEmailFormat = "The provided email address does not has the expected format"
	errTokenFormat = "The provided token does no match with the expected format"
	errNotStock    = "The selected event has not enought stock for this process"
	errActivate    = "Cannot activate %v of %v booked tickets"
	errOrganizer   = "The client does not organize any event with the provided id %v"
	errBelongs     = "The provided ticket does not belongs to event %v, but %v"
)
