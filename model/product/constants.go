package product

// Status represents the ProductStatus enum
type Status int

// Status possible values
const (
	AVAILABLE Status = iota
	RUNNINGOUT
	SOLDOUT
)
