package service

import "database/sql/driver"

// Kind represents the kind enum
type Kind int64

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

var kinds = [...]string{"Groceries", "Information", "Drink",
	"Security", "Bathrooms", "Pharmacy",
	"Safe point", "Purple point", "Hydrogel dispenser",
	"PCRs", "Nursery", "Others"}

func (k Kind) String() string {
	return kinds[k]
}

// Scan ...
func (k *Kind) Scan(value interface{}) error { *k = Kind(value.(int64)); return nil }

// Value ...
func (k Kind) Value() (driver.Value, error) { return int64(k), nil }
