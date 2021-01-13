package event

const (
	errNotFoundByID       = "Not found event with ID %d"
	errNoEventsOnDatabase = "There are no events initialized on the database."
	queryFindAll          = "date(closure_date) >= date(?)"
	queryFilterByType     = "tipus = ? AND date(closure_date) >= date(?)"
)
