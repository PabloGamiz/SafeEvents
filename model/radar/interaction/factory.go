package interaction

import (
	"time"

	"github.com/PabloGamiz/SafeEvents-Backend/mysql"
	"gorm.io/gorm"
)

// New builds a brand new interaction between two clients
func New(owner, closeTo uint, instant time.Time) Controller {
	return &Interaction{
		ClientID: owner,
		CloseTo:  closeTo,
		DoneAt:   instant,
	}
}

// FindCloseToByClientIDAndTime returns all these interactions done by a client from an instant till now
func FindCloseToByClientIDAndTime(id uint, from time.Time) (closeTo []uint, err error) {
	var db *gorm.DB
	var cancel mysql.Disconnect
	if db, cancel, err = mysql.OpenStream(); err != nil {
		return
	}

	defer cancel()
	var interactions []*Interaction
	if err = db.Where(queryFilterByClientIDAndTime, id, from).Find(&interactions).Error; err != nil {
		return
	}

	closeTo = make([]uint, len(interactions))
	for index, interaction := range interactions {
		closeTo[index] = interaction.GetCloseTo()
	}

	return
}
