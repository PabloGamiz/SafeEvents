package client

import (
	"fmt"
	"sync"
)

var (
	//AllInstancesByID stores all clients indexed by its ID
	AllInstancesByID = &sync.Map{}
)

type cID uint

func GetClientByID(id uint) (ctrl Controller, err error) {
	cid := cID(id)

	content, exists := AllInstancesByID.Load(cid)
	if !exists {
		err = fmt.Errorf(errClientNotExists, id)
		return
	}

	var ok bool
	if ctrl, ok = content.(*Client); !ok {
		err = fmt.Errorf(errAssertionFailed)
	}

	return
}
