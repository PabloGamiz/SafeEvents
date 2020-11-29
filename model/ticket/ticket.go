package ticket

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"time"
)

// Ticket its the main data object fro a Ticket
type Ticket struct {
	ID          uint      `json:"id" gorm:"primaryKey; autoIncrement:true"`
	Description string    `json:"description" gorm:"not null"`
	EventID     uint      `json:"event_id"`
	AssistantID uint      `json:"assistant_id"`
	Option      Option    `json:"option" gorm:"not null"`
	QrCode      string    `json:"qr_code" gorm:"unique"`
	CreatedAt   time.Time `json:"createdAt"`
}

func (ticket *Ticket) generateQrCode() (err error) {
	data := make([]byte, 32)
	if _, err = io.ReadFull(rand.Reader, data); err != nil {
		return
	}

	ticket.QrCode = base64.URLEncoding.EncodeToString(data)
	//qrCode, _ := qr.Encode(data, qr.L, qr.Auto)
	//qrCode, _ = barcode.Scale(qrCode, 512, 512)
	//png.Encode(w, qrCode)
	return
}

// Activate checks the ticket as bought, elsewhere its just booked
func (ticket *Ticket) Activate() (err error) {
	if err = ticket.generateQrCode(); err != nil {
		return
	}

	ticket.Option = BOUGHT
	return
}

// GetID return the id of the ticket
func (ticket *Ticket) GetID() uint {
	return ticket.ID
}

// GetOption return the purchase option of the ticket
func (ticket *Ticket) GetOption() Option {
	return ticket.Option
}

// GetCreatedAt return the creation time of the ticket
func (ticket *Ticket) GetCreatedAt() time.Time {
	return ticket.CreatedAt
}
