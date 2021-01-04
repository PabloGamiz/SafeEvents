package ticket

// PayPalPurchaseRequestDTO is the expected struct for a purchase request
type PayPalPurchaseRequestDTO struct {
	Cookie               string  `json:"cookie"`
	TicketID             uint    `json:"ticket_id"`
	TotalAmount          float32 `json:"totalamount"`
	SubTotalAmount       string  `json:"subtotalamount"`
	ShippingCost         string  `json:"shippingcost"`
	ShippingDiscountCost string  `json:"shippingdiscountcost"`
	UserFirstName        string  `json:"userfirstname"`
	UserLastName         string  `json:"userlastname"`
	AddressCity          string  `json:"addresscity"`
	AddressStreet        string  `json:"addressstreet"`
	AddressZipCode       string  `json:"addresszipcode"`
	AddressCountry       string  `json:"addresscountry"`
	AddressState         string  `json:"addressstate"`
	AddressPhoneNumber   string  `json:"addressphonenumber"`
}

// PayPalPurchaseResponseDTO is the response for a purchase request
type PayPalPurchaseResponseDTO struct {
	Error string `json:"error"`
}
