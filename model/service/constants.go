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

func (k Kind) String() string {
	return [...]string{"Groceries", "Information", "Drink",
		"Security", "Bathrooms", "Pharmacy",
		"Safe point", "Purple point", "Hydrogel dispenser",
		"PCRs", "Nursery", "Others"}[k]
}
