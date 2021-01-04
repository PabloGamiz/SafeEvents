package event

import (
	"context"
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/PabloGamiz/SafeEvents-Backend/model/feedback"
	"github.com/PabloGamiz/SafeEvents-Backend/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	// AllInstancesByID stores all events indexed by its ID
	AllInstancesByID = &sync.Map{}
	once             sync.Once
)

type sID uint

//
//func LoadOrStoreNewEvent(event Controller) error {
//	sid := event.GetID()
//	if
//}

// FindAll returns the controllers of all the events loaded on the BBDD
func FindAll(ctx context.Context) (ctrl []Controller, err error) {
	var db *gorm.DB
	var cancel mysql.Disconnect
	if db, cancel, err = mysql.OpenStream(); err != nil {
		return
	}

	defer cancel()
	var eventsMOD []*Event
	db.Preload(clause.Associations).Find(&eventsMOD)
	fmt.Println(eventsMOD)
	ctrl = make([]Controller, len(eventsMOD))
	for index, event := range eventsMOD {
		ctrl[index] = event
	}

	return
}

// FindAllByType returns the controllers of all the events loaded on the BBDD
func FindAllByType(ctx context.Context, eventType string) (ctrl []Controller, err error) {
	var db *gorm.DB
	var cancel mysql.Disconnect
	if db, cancel, err = mysql.OpenStream(); err != nil {
		return
	}

	defer cancel()
	var eventsMOD []*Event
	db.Preload(clause.Associations).Where(queryFilterByType, eventType).Find(&eventsMOD)
	fmt.Println(eventsMOD)
	ctrl = make([]Controller, len(eventsMOD))
	for index, event := range eventsMOD {
		ctrl[index] = event
	}

	return
}

// FindEventByID returns the gateway for the event that match the provided name
func FindEventByID(ID uint) (ctrl Controller, err error) {
	var db *gorm.DB
	var cancel mysql.Disconnect
	if db, cancel, err = mysql.OpenStream(); err != nil {
		return
	}
	var events []*Event
	defer cancel()
	db.Preload(clause.Associations).Preload("Services.Products").Where("id = ?", ID).Find(&events)
	if len(events) == 0 {
		err = fmt.Errorf(errNotFoundByID, ID)
		return
	}

	ctrl = events[0]

	return
}

// FindRecomended returns the controllers of all the events recomended for the user
func FindRecomended(ctx context.Context, clientID uint) (ctrl []Controller, err error) {
	var db *gorm.DB
	var cancel mysql.Cancel
	if db, cancel, err = mysql.OpenStream(); err != nil {
		return
	}
	m := map[string]int{}
	var PopularEvts []*Event
	defer cancel()
	var idAssis uint
	db.Select("id").Where("client_id = ?", clientID).Table("assistants").Find(&idAssis)
	db.Raw("SELECT * FROM `events` WHERE capacity <> taken AND closure_date > CURRENT_DATE ORDER BY taken desc").Scan(&PopularEvts) //OBTE TOTS ELS ESDEVENIMENTS DISPONIBLES ORDENATS PER PLACES LLIURES
	type Result struct {
		EventID uint
	}
	var result []Result
	//GUARDA LES ASOCIACIONS TICKET_ID I ASSISTANT_ID DELS ESDEVENIMENTS PASATS
	db.Raw("SELECT DISTINCT tickets.event_id FROM tickets INNER JOIN events on tickets.event_id = events.id WHERE events.closure_date < CURRENT_DATE AND tickets.assistant_id = ?", idAssis).Scan(&result)
	if len(PopularEvts) == 0 {
		err = fmt.Errorf("Not Available Events for client ID %d", clientID)
		return
	}
	//CONSULTA ELS RATINGS DELS FEEDBACKS
	for _, valor := range result {
		var fb feedback.Feedback
		db.Preload(clause.Associations).Table("feedbacks").Where("event_id = ?", valor.EventID).Where("assistant_id = ?", idAssis).Find(&fb)
		var evnt []*Event
		db.Table("events").Where("id = ?", valor.EventID).Find(&evnt)
		if len(evnt) != 0 {
			if fb.ID != 0 {
				if fb.Rating > 3 {
					m[evnt[0].Tipus] += 10
				} else {
					m[evnt[0].Tipus] -= 10
				}
			} else {
				m[evnt[0].Tipus] += 5
			}
		}
	}
	//ES TREUEN ELS ESDEVENIMENTS PROPERS QUE JA TE COMPRATS
	db.Raw("SELECT DISTINCT tickets.event_id FROM tickets INNER JOIN events on tickets.event_id = events.id WHERE events.closure_date >= CURRENT_DATE AND tickets.assistant_id = ?", idAssis).Scan(&result)
	for _, j := range result {
		for x := 0; x < len(PopularEvts); x++ {
			if j.EventID == PopularEvts[x].ID {
				PopularEvts = append(PopularEvts[:x], PopularEvts[x+1:]...)
			}
		}
	}
	if len(PopularEvts) == 0 {
		err = fmt.Errorf("Not Available Events for client ID %d", clientID)
		return
	}
	//CONSULTA ELS FAVS DEL CLIENT
	var eventid []uint
	mapdeMots := map[string]int{}
	db.Raw("SELECT `event_id` FROM `clients_favs` WHERE client_id = ?", clientID).Scan(&eventid)
	if len(eventid) != 0 {
		for _, valor := range eventid {
			var evnt []*Event
			db.Table("events").Where("id = ?", valor).Find(&evnt)
			if len(evnt) != 0 {
				m[evnt[0].Tipus] += 10
			}
			titolev := strings.Split(evnt[0].Title, " ")
			for _, palabra := range titolev {
				mapdeMots[palabra]++
			}
		}
	}
	for key := range mapdeMots {
		if strings.ToLower(key) == key {
			delete(mapdeMots, key)
		}
	}
	var primerInteres []*Event //Els que tenen coincidencia amb el nom
	var segonInteres []*Event  //Els que tenen feedback positiu o estan a preferits
	var tercerInteres []*Event //La resta dels esdeveniments disponibles
	//AFEGEIX TOTS ELS EVENTS CANDIDATS SEGONS LA PREFERENCIA OBTENIDA AMB ELS PASOS ANTERIORS
	for _, aux := range PopularEvts {
		titlesplit := strings.Split(aux.Title, " ")
		trobat := false
		for s := 0; s < len(titlesplit); s++ {
			_, found := mapdeMots[titlesplit[s]]
			if found {
				trobat = true
				primerInteres = append(primerInteres, aux)
				s = len(titlesplit)
			}
		}
		if !trobat {
			valor, found := m[aux.Tipus]
			if found {
				if valor < 0 {
					tercerInteres = append(tercerInteres, aux)
				} else {
					segonInteres = append(segonInteres, aux)
				}
			} else {
				tercerInteres = append(tercerInteres, aux)
			}
		}

	}

	resultsize := int(math.Min(20, float64(len(PopularEvts))))

	ctrl = make([]Controller, resultsize)
	for i := 0; i < resultsize; i++ {
		if len(primerInteres) != 0 {
			ctrl[i] = primerInteres[0]
			primerInteres = append(primerInteres[1:])
		} else if len(segonInteres) != 0 {
			ctrl[i] = segonInteres[0]
			segonInteres = append(segonInteres[1:])
		} else {
			ctrl[i] = tercerInteres[0]
			tercerInteres = append(tercerInteres[1:])
		}
	}
	return
}
