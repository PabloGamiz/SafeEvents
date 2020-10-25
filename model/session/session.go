package session

import "github.com/PabloGamiz/SafeEvents-Backend/model/client"

// Session its the main data object fro a client
type Session struct {
	client.Controller
	Cookie   string `json:"cookie" bson:"_cookie,omitempty"`
	Deadline uint64 `json:"deadline" bson:"_deadline,omitempty"`
}

// GetCookie return the id of the client
func (session *Session) GetCookie() string {
	return session.Cookie
}

// GetDeadline return the deadline of the client in unix64
func (session *Session) GetDeadline() uint64 {
	return session.Deadline
}
