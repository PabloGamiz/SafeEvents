package radar

import (
	"fmt"
	"sync"

	interactionMOD "github.com/PabloGamiz/SafeEvents-Backend/model/radar/interaction"
)

var (
	// AllInstancesByMAC stores all radars indexed by its MAC
	AllInstancesByMAC = &sync.Map{}
)

type macID string

// FindRadarByMAC returns the radar associated to the provided MAC
func FindRadarByMAC(MAC string) (ctrl Controller, err error) {
	mID := macID(MAC)
	content, ok := AllInstancesByMAC.Load(mID)
	if !ok {
		err = fmt.Errorf("No radar has been found for the provided MAC")
		return
	}

	if ctrl, ok = content.(Controller); !ok {
		err = fmt.Errorf("Got an error while asserting radar to controller")
	}

	return
}

// NewRadar builds a brand new radar for the provided MAC and ClientID
func NewRadar(mac string, clientID uint) (ctrl Controller, err error) {
	mID := macID(mac)
	if _, exists := AllInstancesByMAC.Load(mID); exists {
		err = fmt.Errorf("Already exists a radar for the provided MAC")
		return
	}

	radar := &Radar{
		MAC:          mac,
		ClientID:     clientID,
		interactions: []interactionMOD.Controller{},
	}

	AllInstancesByMAC.Store(mID, radar)
	return radar, nil
}
