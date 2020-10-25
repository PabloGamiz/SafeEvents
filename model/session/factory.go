package session

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"sync"

	"github.com/PabloGamiz/SafeEvents-Backend/model/client"
)

var (
	allInstancesByID    = &sync.Map{}
	allInstancesByEmail = &sync.Map{}
)

type sID string
type email string

func registerSession(session *Session) (err error) {
	sid := sID(session.Cookie)
	email := email(session.GetEmail())

	if _, exists := allInstancesByID.Load(sid); exists {
		return fmt.Errorf("Session with the provided ID already exists")
	} else if _, exists = allInstancesByEmail.Load(email); exists {
		return fmt.Errorf("There is already a session for the provided user email")
	}

	allInstancesByID.Store(sid, session)
	allInstancesByEmail.Store(email, session)
	return
}

func removeSession(sid sID) (err error) {
	content, exists := allInstancesByID.Load(sid)
	if !exists {
		return fmt.Errorf(errSessionNotExists, sid)
	}

	session, exists := content.(*Session)
	if !exists {
		return fmt.Errorf(errAssertionFailed)
	}

	email := email(session.GetEmail())
	allInstancesByID.Delete(sid)
	allInstancesByEmail.Delete(email)
	return
}

func newSessionID() (id sID, err error) {
	b := make([]byte, 32)
	if _, err = io.ReadFull(rand.Reader, b); err != nil {
		return
	}

	raw := base64.URLEncoding.EncodeToString(b)
	id = sID(raw)
	return
}

// GetSession returns the session with the provided cookie, if exists
func GetSession(cookie string) (ctrl Controller, err error) {
	sid := sID(cookie)

	content, exists := allInstancesByID.Load(sid)
	if !exists {
		err = fmt.Errorf(errSessionNotExists, cookie)
		return
	}

	var ok bool
	if ctrl, ok = content.(*Session); !ok {
		err = fmt.Errorf(errAssertionFailed)
	}

	return
}

// NewSession returns a brand new session for the provided client
func NewSession(client client.Controller) (ctrl Controller, err error) {
	var sid sID
	if sid, err = newSessionID(); err != nil {
		return
	}

	cookie := string(sid)
	session := &Session{
		Controller: client,
		Cookie:     cookie,
	}

	err = registerSession(session)
	ctrl = session
	return
}

// KillSession logs out the session with the provided cookie
func KillSession(cookie string) error {
	sid := sID(cookie)
	return removeSession(sid)
}
