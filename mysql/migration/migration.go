package migration

import (
	"log"

	"github.com/PabloGamiz/SafeEvents-Backend/model/client"
	"github.com/PabloGamiz/SafeEvents-Backend/model/client/assistant"
	"github.com/PabloGamiz/SafeEvents-Backend/model/client/organizer"
	"github.com/PabloGamiz/SafeEvents-Backend/model/event"
	"github.com/PabloGamiz/SafeEvents-Backend/model/feedback"
	"github.com/PabloGamiz/SafeEvents-Backend/model/product"
	"github.com/PabloGamiz/SafeEvents-Backend/model/service"
	"github.com/PabloGamiz/SafeEvents-Backend/model/ticket"
	"github.com/PabloGamiz/SafeEvents-Backend/mysql"
	"gorm.io/gorm"
)

func OpenStream() (db *gorm.DB, err error) {
	if db, err = mysql.OpenStream(); err != nil {
		log.Fatalf("Got %v error while OpenStream", err.Error())
		return
	}

	return
}

// MigrateTables migrates the database tables
func MigrateTables() (err error) {
	db, err := OpenStream()
	if err != nil {
		return
	}

	db.AutoMigrate(&client.Client{},
		&organizer.Organizer{},
		&assistant.Assistant{},
		&feedback.Feedback{},
		&ticket.Ticket{},
		&event.Event{},
		&service.Service{},
		&product.Product{})

	return
}
