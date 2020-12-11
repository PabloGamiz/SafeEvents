package feedback

import (
	"context"

	"github.com/PabloGamiz/SafeEvents-Backend/model/feedback"
	"github.com/PabloGamiz/SafeEvents-Backend/mysql"
	"gorm.io/gorm"
)

type feedbackGateway struct {
	feedback.Controller
	ctx context.Context
}

func (gw *feedbackGateway) Insert() (err error) {
	var db *gorm.DB
	if db, err = mysql.OpenStream(); err != nil {
		return
	}

	db.Create(gw.Controller)
	return
}

func (gw *feedbackGateway) Update() (err error) {
	var db *gorm.DB
	if db, err = mysql.OpenStream(); err != nil {
		return
	}

	db.Save(gw.Controller)
	return
}

func (gw *feedbackGateway) Remove() (err error) {
	var db *gorm.DB
	if db, err = mysql.OpenStream(); err != nil {
		return
	}

	db.Delete(gw.Controller)
	return
}
