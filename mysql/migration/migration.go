package migration

import (
	"log"

	"github.com/PabloGamiz/SafeEvents-Backend/mysql"
	"gorm.io/gorm"
)

func openMigrationStream() (db *gorm.DB, err error) {
	if db, err = mysql.OpenStream(); err != nil {
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

	db.AutoMigrate()

	return
}
