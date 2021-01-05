package ticket

import (
	"github.com/PabloGamiz/SafeEvents-Backend/mysql"
	"gorm.io/gorm"
)

// GetTicketsByEventID return all current tickets for a given event
func GetTicketsByEventID(id uint) (tickets []Controller, err error) {
	var db *gorm.DB
	var cancel mysql.Disconnect
	if db, cancel, err = mysql.OpenStream(); err != nil {
		return
	}

	defer cancel()
	db.Table("tickets").Where(queryFindByEventID, id).Find(&tickets)
	return
}

// GetTicketsByEventIDAndClientID return all current tickets for a given event and client
func GetTicketsByEventIDAndClientID(eid uint, cid uint) (ctrls []Controller, err error) {
	var db *gorm.DB
	var cancel mysql.Disconnect
	if db, cancel, err = mysql.OpenStream(); err != nil {
		return
	}

	defer cancel()
	var tickets []Ticket
	db.Table("tickets").Where(queryFindByEventIDAndClientID, eid, cid).Find(&tickets)

	ctrls = make([]Controller, len(tickets))
	for index, ticket := range tickets {
		ctrls[index] = &ticket
	}

	return
}

// GetTicketByQR return all current tickets for a given event and client
func GetTicketByQR(qr string) (ctrl Controller, err error) {
	var db *gorm.DB
	var cancel mysql.Disconnect
	if db, cancel, err = mysql.OpenStream(); err != nil {
		return
	}

	defer cancel()
	var ticket Ticket
	db = db.Table("tickets").Where(queryFindByQR, qr).Find(&ticket)

	return &ticket, db.Error
}

// GetTicketByID return all current tickets for a given event and client
func GetTicketByID(id uint) (ctrl Controller, err error) {
	var db *gorm.DB
	var cancel mysql.Disconnect
	if db, cancel, err = mysql.OpenStream(); err != nil {
		return
	}

	defer cancel()
	var ticket Ticket
	db = db.Table("tickets").Where("id = ?", id).Find(&ticket)

	return &ticket, db.Error
}
