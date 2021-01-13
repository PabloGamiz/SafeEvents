package buyerdata

// Controller represents a Product and it's main data
type Controller interface {
	GetID() uint
	GetTicketID() uint
	GetTotalAmount() float32
	SetTotalAmount(totalamount float32)
	GetFirstName() string
	SetFirstName(firstname string)
	GetLastName() string
	SetLastName(lastname string)
	GetAddressCity() string
	SetAddressCity(addresscity string)
	GetAddressStreet() string
	SetAddressStreet(addressstreet string)
	GetAddressZipCode() string
	SetAddressZipCode(addresszipcode string)
	GetAddressCountry() string
	SetAddressCountry(addresscountry string)
	GetAddressState() string
	SetAddressState(addressstate string)
	GetPhoneNumber() string
	SetPhoneNumber(phonenumber string)
}
