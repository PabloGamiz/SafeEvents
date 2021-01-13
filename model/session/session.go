package session

import (
	"context"
	"fmt"

	"github.com/PabloGamiz/SafeEvents-Backend/model/client"
	radarMOD "github.com/PabloGamiz/SafeEvents-Backend/model/radar"
)

// Session its the main data object fro a client
type Session struct {
	client.Controller                     // client associated to this session
	context.Context                       // session context with deadline
	cancel            context.CancelFunc  // session context canceler
	cookie            string              // session ID
	token             string              // google token for this session
	radar             radarMOD.Controller // session radar
}

// Cookie return the id of the client
func (session *Session) Cookie() string {
	return session.cookie
}

// Client returns the client logged with this session
func (session *Session) Client() client.Controller {
	return session.Controller
}

// GetRadar returns the radar for this session, if any
func (session *Session) GetRadar() radarMOD.Controller {
	return session.radar
}

// InitRadar activates the radar for this session
func (session *Session) InitRadar(MAC string) (err error) {
	if session.radar != nil {
		return fmt.Errorf("This user has the radar already activated")
	}

	session.radar, err = radarMOD.NewRadar(MAC, session.GetID())
	return
}

// FinishRadar deactivates the radar for this session
func (session *Session) FinishRadar() (err error) {
	if session.radar == nil {
		return fmt.Errorf("This session has no radar to deactivate")
	}

	if err = session.radar.Close(); err != nil {
		return
	}

	session.radar = nil
	return
}
