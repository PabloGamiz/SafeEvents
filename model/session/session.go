package session

import (
	"context"

	"github.com/PabloGamiz/SafeEvents-Backend/model/client"
)

// Session its the main data object fro a client
type Session struct {
	context.Context                      // session context with deadline
	client.Controller                    // client associated to this session
	cancel            context.CancelFunc // session context canceler
	cookie            string             // session ID
	token             string             // google token for this session
}

// Cookie return the id of the client
func (session *Session) Cookie() string {
	return session.cookie
}
