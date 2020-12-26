package radar

import (
	"fmt"

	interactionMOD "github.com/PabloGamiz/SafeEvents-Backend/model/radar/interaction"
)

// Radar its the main data object from a radar
type Radar struct {
	MAC          string `json:"device" gorm:"not null"`
	ClientID     uint   `json:"client_id"`
	interactions []interactionMOD.Controller
}

// Init initializes the radar
func (radar *Radar) Init() (err error) {
	return
}

// Close deactivates the radar
func (radar *Radar) Close() (err error) {
	return
}

// GetID returns te client id of the radar's owner
func (radar *Radar) GetID() uint {
	return radar.ClientID
}

// GetMAC returns te MAC address of the  logged device
func (radar *Radar) GetMAC() string {
	return radar.MAC
}

// SetMAC sets the MAC for this radar
func (radar *Radar) SetMAC(mac string) (old string, updated bool) {
	if old = radar.MAC; len(old) > 0 {
		updated = true
	}

	radar.MAC = mac
	return
}

// SetInteractions registers a set of new interactions into this radar
func (radar *Radar) SetInteractions(newInteractions []interactionMOD.Controller) int {
	radar.interactions = append(radar.interactions, newInteractions...)
	return len(radar.interactions)
}

// PopInteractions removes the latest n interactions registered in the radar
func (radar *Radar) PopInteractions(n int) (err error) {
	lenght := len(radar.interactions)
	if n > lenght {
		return fmt.Errorf("Not enough interactions to remove")
	}

	radar.interactions = radar.interactions[0 : lenght-n-1]
	return
}
