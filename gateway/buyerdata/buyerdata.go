package buyerdata

import (
	"context"

	"github.com/PabloGamiz/SafeEvents-Backend/model/buyerdata"
	"github.com/PabloGamiz/SafeEvents-Backend/mysql"
	"gorm.io/gorm"
)

type buyerdataGateway struct {
	buyerdata.Controller
	ctx context.Context
}

func (gw *buyerdataGateway) Insert() (err error) {
	var db *gorm.DB
	var cancel mysql.Cancel
	if db, cancel, err = mysql.OpenStream(); err != nil {
		return
	}

	defer cancel()
	db.Create(gw.Controller)
	return
}

func (gw *buyerdataGateway) Update() (err error) {
	var db *gorm.DB
	var cancel mysql.Cancel
	if db, cancel, err = mysql.OpenStream(); err != nil {
		return
	}

	defer cancel()
	db.Table("buyer_data").Updates(gw.Controller)
	return
}

func (gw *buyerdataGateway) Remove() (err error) {
	var db *gorm.DB
	var cancel mysql.Cancel
	if db, cancel, err = mysql.OpenStream(); err != nil {
		return
	}

	defer cancel()
	db.Table("buyer_data").Delete(gw.Controller)
	return
}
