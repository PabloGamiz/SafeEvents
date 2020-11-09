package session

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"sync"

	"github.com/PabloGamiz/SafeEvents-Backend/model/client"
)

var (
	// AllInstancesByID stores all sessions indexed by its ID
	AllInstancesByID = &sync.Map{}
	// AllInstancesByEmail stores all sessions indexed by its email
	AllInstancesByEmail = &sync.Map{}
)

type sID string
type email string

func registerSession(session *Session) (err error) {
	sid := sID(session.Cookie())
	email := email(session.GetEmail())

	if _, exists := AllInstancesByID.Load(sid); exists {
		return fmt.Errorf("Session with the provided ID already exists")
	} else if _, exists = AllInstancesByEmail.Load(email); exists {
		return fmt.Errorf("There is already a session for the provided user email")
	}

	AllInstancesByID.Store(sid, session)
	AllInstancesByEmail.Store(email, session)
	return
}

func removeSession(sid sID) (err error) {
	content, exists := AllInstancesByID.Load(sid)
	if !exists {
		return fmt.Errorf(errSessionNotExists, sid)
	}

	session, exists := content.(*Session)
	if !exists {
		return fmt.Errorf(errAssertionFailed)
	}

	email := email(session.GetEmail())
	AllInstancesByID.Delete(sid)
	AllInstancesByEmail.Delete(email)
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

// GetSessionByID returns the session with the provided cookie, if exists
func GetSessionByID(cookie string) (ctrl Controller, err error) {
	sid := sID(cookie)

	content, exists := AllInstancesByID.Load(sid)
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

// GetSessionByEmail returns the session with the provided email, if exists
func GetSessionByEmail(mail string) (ctrl Controller, err error) {
	email := email(mail)

	content, exists := AllInstancesByEmail.Load(email)
	if !exists {
		err = fmt.Errorf(errSessionNotExists, email)
		return
	}

	var ok bool
	if ctrl, ok = content.(*Session); !ok {
		err = fmt.Errorf(errAssertionFailed)
	}

	return
}

// NewSession returns a brand new session for the provided client
func NewSession(ctx context.Context, cancel context.CancelFunc, client client.Controller) (ctrl Controller, err error) {
	if _, ok := ctx.Deadline(); !ok {
		err = fmt.Errorf(errNoDeadline)
		return
	}

	var sid sID
	if sid, err = newSessionID(); err != nil {
		return
	}

	cookie := string(sid)
	session := &Session{
		Context:    ctx,
		Controller: client,
		cancel:     cancel,
		cookie:     cookie,
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
