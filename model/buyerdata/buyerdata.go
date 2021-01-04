package buyerdata

import (
	"time"

	"github.com/PabloGamiz/SafeEvents-Backend/model/ticket"
)

// BuyerData represents the product class from UML
type BuyerData struct {
	ID                   uint           `json:"id" gorm:"primaryKey; autoIncrement:true"`
	Ticket               *ticket.Ticket `json:"ticket" gorm:"foreignkey:TicketID"`
	TicketID             uint           `json:"ticket_id" gorm: "index:uc_ticket,unique; not null"`
	TotalAmount          float32        `json:"totalamount" gorm:"not null"`
	SubTotalAmount       string         `json:"subtotalamount" gorm:"not null"`
	ShippingCost         string         `json:"shippingcost" gorm:"not null"`
	ShippingDiscountCost string         `json:"shippingdiscountcost" gorm:"not null"`
	FirstName            string         `json:"firstname" gorm:"not null"`
	LastName             string         `json:"lastname" gorm:"not null"`
	AddressCity          string         `json:"addresscity" gorm:"not null"`
	AddressStreet        string         `json:"addressstreet" gorm:"not null"`
	AddressZipCode       string         `json:"addresszipcode" gorm:"not null"`
	AddressCountry       string         `json:"addresscountry" gorm:"not null"`
	AddressState         string         `json:"addressstate" gorm:"not null"`
	AddressPhoneNumber   string         `json:"addressphonenumber" gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

// GetID return the ID of the buyerdata.
func (buyerdata *BuyerData) GetID() uint {
	return buyerdata.ID
}

// GetTicketID return the ID of the buyerdata.
func (buyerdata *BuyerData) GetTicketID() uint {
	return buyerdata.TicketID
}

// GetTotalAmount return the Name of the buyerdata.
func (buyerdata *BuyerData) GetTotalAmount() float32 {
	return buyerdata.TotalAmount
}

// SetTotalAmount sets the TotalAmount of the buyerdata.
func (buyerdata *BuyerData) SetTotalAmount(totalamount float32) {
	buyerdata.TotalAmount = totalamount
}

// GetSubTotalAmount return the Description of the buyerdata.
func (buyerdata *BuyerData) GetSubTotalAmount() string {
	return buyerdata.SubTotalAmount
}

// SetSubTotalAmount sets the Description of the buyerdata.
func (buyerdata *BuyerData) SetSubTotalAmount(subtotalamount string) {
	buyerdata.SubTotalAmount = subtotalamount
}

// GetShippingCost return the Description of the buyerdata.
func (buyerdata *BuyerData) GetShippingCost() string {
	return buyerdata.ShippingCost
}

// SetShippingCost sets the Description of the buyerdata.
func (buyerdata *BuyerData) SetShippingCost(shippingcost string) {
	buyerdata.SubTotalAmount = shippingcost
}

// GetShippingDiscountCost return the Description of the buyerdata.
func (buyerdata *BuyerData) GetShippingDiscountCost() string {
	return buyerdata.ShippingDiscountCost
}

// SetShippingDiscountCost sets the Description of the buyerdata.
func (buyerdata *BuyerData) SetShippingDiscountCost(shippingdiscountcost string) {
	buyerdata.ShippingDiscountCost = shippingdiscountcost
}

// GetFirstName return the Description of the buyerdata.
func (buyerdata *BuyerData) GetFirstName() string {
	return buyerdata.FirstName
}

// SetFirstName sets the Description of the buyerdata.
func (buyerdata *BuyerData) SetFirstName(firstname string) {
	buyerdata.FirstName = firstname
}

// GetLastName return the Price of the buyerdata.
func (buyerdata *BuyerData) GetLastName() string {
	return buyerdata.LastName
}

// SetLastName sets the Price of the buyerdata.
func (buyerdata *BuyerData) SetLastName(lastname string) {
	buyerdata.LastName = lastname
}

// GetAddressCity return the Price of the buyerdata.
func (buyerdata *BuyerData) GetAddressCity() string {
	return buyerdata.AddressCity
}

// SetAddressCity sets the Price of the buyerdata.
func (buyerdata *BuyerData) SetAddressCity(addressCity string) {
	buyerdata.AddressCity = addressCity
}

// GetAddressStreet return the Price of the buyerdata.
func (buyerdata *BuyerData) GetAddressStreet() string {
	return buyerdata.AddressStreet
}

// SetAddressStreet sets the Price of the buyerdata.
func (buyerdata *BuyerData) SetAddressStreet(addressstreet string) {
	buyerdata.AddressStreet = addressstreet
}

// GetAddressZipCode return the Price of the buyerdata.
func (buyerdata *BuyerData) GetAddressZipCode() string {
	return buyerdata.AddressZipCode
}

// SetAddressZipCode sets the Price of the buyerdata.
func (buyerdata *BuyerData) SetAddressZipCode(addresszipcode string) {
	buyerdata.AddressZipCode = addresszipcode
}

// GetAddressCountry return the Price of the buyerdata.
func (buyerdata *BuyerData) GetAddressCountry() string {
	return buyerdata.AddressCountry
}

// SetAddressCountry sets the Price of the buyerdata.
func (buyerdata *BuyerData) SetAddressCountry(addresscountry string) {
	buyerdata.AddressCountry = addresscountry
}

// GetAddressState return the Price of the buyerdata.
func (buyerdata *BuyerData) GetAddressState() string {
	return buyerdata.AddressState
}

// SetAddressState sets the Price of the buyerdata.
func (buyerdata *BuyerData) SetAddressState(addressstate string) {
	buyerdata.AddressState = addressstate
}

// GetPhoneNumber return the Price of the buyerdata.
func (buyerdata *BuyerData) GetPhoneNumber() string {
	return buyerdata.AddressPhoneNumber
}

// SetPhoneNumber sets the Price of the buyerdata.
func (buyerdata *BuyerData) SetPhoneNumber(phonenumber string) {
	buyerdata.AddressPhoneNumber = phonenumber
}
