package event

import (
	eventDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/event"
	"github.com/alvidir/util/pattern/transaction"
)

// NewtxPublicaEvent builds a brand new transaction for Publica_event
func NewtxPublicaEvent(request eventDTO.Publica_eventRequestDTO) transaction.Tx {
	body := &txSignin{request: request}
	return transaction.NewTransaction(body)
}
