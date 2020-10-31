package service

// Kind represents the kind enum
type Kind int

// Kind possible values
const (
	GROCERIES Kind = iota
	INFORMATION
	DRINK
	SECURITY
	BATHROOMS
	PHARMACY
	SAFEPOINT
	PURPLEPOINT
	HYDROGELDISPENSER
	PCRs
	NUERSERY
	OTHERS
)
