package migration

import (
	"log"

	"github.com/PabloGamiz/SafeEvents-Backend/model/buyerdata"

	"github.com/PabloGamiz/SafeEvents-Backend/model/client"
	"github.com/PabloGamiz/SafeEvents-Backend/model/client/assistant"
	"github.com/PabloGamiz/SafeEvents-Backend/model/client/organizer"
	"github.com/PabloGamiz/SafeEvents-Backend/model/event"
	"github.com/PabloGamiz/SafeEvents-Backend/model/feedback"
	"github.com/PabloGamiz/SafeEvents-Backend/model/product"
	"github.com/PabloGamiz/SafeEvents-Backend/model/radar/interaction"
	"github.com/PabloGamiz/SafeEvents-Backend/model/service"
	"github.com/PabloGamiz/SafeEvents-Backend/model/ticket"
	"github.com/PabloGamiz/SafeEvents-Backend/mysql"
	"gorm.io/gorm"
)

// MigrateTables migrates the database tables
func MigrateTables() (err error) {
	var db *gorm.DB
	var cancel mysql.Disconnect
	if db, cancel, err = mysql.OpenStream(); err != nil {
		log.Fatalf("Got %v error while OpenStream", err.Error())
		return
	}

	defer cancel()
	if err != nil {
		return
	}

	db.AutoMigrate(&client.Client{},
		&buyerdata.BuyerData{},
		&organizer.Organizer{},
		&assistant.Assistant{},
		&feedback.Feedback{},
		&ticket.Ticket{},
		&event.Event{},
		&service.Service{},
		&product.Product{},
		&interaction.Interaction{})

	return
}
