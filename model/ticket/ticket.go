package ticket

import (
	"github.com/PabloGamiz/SafeEvents-Backend/model/client"
	"github.com/PabloGamiz/SafeEvents-Backend/model/event"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"gorm.io/gorm"
)

// Ticket its the main data object fro a Ticket
type Ticket struct {
	gorm.Model
	ID          uint           `json:"id" gorm:"primaryKey; autoIncrement:true"`
	Description string         `json:"description" gorm:"not null"`
	Price       float32        `json:"price" gorm:"not null"`
	Client      *client.Client `json:"-" gorm:"foreingkey:ClientID"`
	Event       *event.Event   `json:"-" gorm:"foreingkey:EventID"`
	Option      Option         `json:"option" gorm:"not null"`
	ClientID    uint
	EventID     uint
	QrCode      *barcode.Barcode
}

func (ticket *Ticket) generateQrCode() {
	data := "dataString"
	qrCode, _ := qr.Encode(data, qr.L, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 512, 512)

	ticket.QrCode = &qrCode
	//png.Encode(w, qrCode)
}

// GetID return the id of the ticket
func (ticket *Ticket) GetID() uint {
	return ticket.ID
}
