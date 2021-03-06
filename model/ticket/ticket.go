package ticket

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"time"
)

// Ticket its the main data object fro a Ticket
type Ticket struct {
	ID          uint       `json:"id" gorm:"primaryKey; autoIncrement:true"`
	Description string     `json:"description" gorm:"not null"`
	EventID     uint       `json:"event_id"`
	AssistantID uint       `json:"assistant_id"`
	Option      Option     `json:"option" gorm:"not null"`
	QrCode      *string    `json:"qr_code" gorm:"unique"`
	CreatedAt   time.Time  `json:"createdAt"`
	CheckIn     *time.Time `json:"check_in"`
	ClientID    uint       `json:"client_id"`
}

func (ticket *Ticket) generateQrCode() (err error) {
	data := make([]byte, 32)
	if _, err = io.ReadFull(rand.Reader, data); err != nil {
		return
	}

	qr := base64.URLEncoding.EncodeToString(data)
	ticket.QrCode = &qr
	//qrCode, _ := qr.Encode(data, qr.L, qr.Auto)
	//qrCode, _ = barcode.Scale(qrCode, 512, 512)
	//png.Encode(w, qrCode)
	return
}

// Activate checks the ticket as bought, elsewhere its just booked
func (ticket *Ticket) Activate() (err error) {
	if ticket.Option != BOOKED {
		return fmt.Errorf("Ticket already activated or checked")
	}

	if err = ticket.generateQrCode(); err != nil {
		return
	}

	ticket.Option = BOUGHT
	return
}

// Check checks the tickets if it isn't already
func (ticket *Ticket) Check() (err error) {
	switch ticket.Option {
	case CHECKED:
		err = fmt.Errorf("This ticket has been already checked")
	case BOOKED:
		err = fmt.Errorf("This ticket has not been purchased")
	case BOUGHT:
		current := time.Now()
		ticket.Option = CHECKED
		ticket.CheckIn = &current
	}

	return
}

// GetID return the id of the ticket
func (ticket *Ticket) GetID() uint {
	return ticket.ID
}

// GetQR return the id of the ticket
func (ticket *Ticket) GetQR() string {
	if ticket.QrCode != nil {
		return *ticket.QrCode
	}

	return ""
}

// GetOption return the purchase option of the ticket
func (ticket *Ticket) GetOption() Option {
	return ticket.Option
}

// GetCreatedAt return the creation time of the ticket
func (ticket *Ticket) GetCreatedAt() time.Time {
	return ticket.CreatedAt
}

// GetClientID return the client of the ticket
func (ticket *Ticket) GetClientID() uint {
	return ticket.AssistantID
}

// GetEventID return the client of the ticket
func (ticket *Ticket) GetEventID() uint {
	return ticket.EventID
}

// GetInstance return itself as an instance of ticket
func (ticket *Ticket) GetInstance() *Ticket {
	return ticket
}
