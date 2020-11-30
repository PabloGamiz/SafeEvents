package mysql

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
	"gorm.io/gorm"
)

func openMigrationStream() (db *gorm.DB, err error) {
	if db, err = OpenStream(); err != nil {
		log.Fatalf("Got %v error while opening stream", err.Error())
		return
	}

	return
}

func migrateTables() (err error) {
	db, err := openMigrationStream()
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
