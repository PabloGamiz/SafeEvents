package ticket

import (
	"crypto/rand"
	"encoding/base64"
	"io"

	"github.com/PabloGamiz/SafeEvents-Backend/model/client"
	"github.com/PabloGamiz/SafeEvents-Backend/model/event"
	"gorm.io/gorm"
)

// Ticket its the main data object fro a Ticket
type Ticket struct {
	gorm.Model
	ID          uint           `json:"id" gorm:"primaryKey; autoIncrement:true"`
	Description string         `json:"description" gorm:"not null"`
	Client      *client.Client `json:"-" gorm:"foreingkey:ClientID"`
	Event       *event.Event   `json:"-" gorm:"foreingkey:EventID"`
	Option      Option         `json:"option" gorm:"not null"`
	QrCode      string         `json:"qr_code" gorm:"unique"`
	ClientID    uint           `json:"client_id"`
	EventID     uint           `json:"event_id"`
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

// Buy checks the ticket as bought, elsewhere its just booked
func (ticket *Ticket) Buy() (err error) {
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
