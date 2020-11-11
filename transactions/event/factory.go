package event

import (
	eventDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/event"
	"github.com/alvidir/util/pattern/transaction"
)

// NewTxListEvents builds a brand new transaction for List Events
func NewTxListEvents() transaction.Tx {
	body := &txListEvents{}
	return transaction.NewTransaction(body)
}

// NewtxPublicaEvent builds a brand new transaction for Publica_event
func NewtxPublicaEvent(request eventDTO.DTO) transaction.Tx {
	body := &txPublicaEvent{request: request}
	return transaction.NewTransaction(body)
}

// NewtxGetEvent builds a brand new transaction for Publica_event
func NewtxGetEvent(request eventDTO.DTO) transaction.Tx {
	body := &txGetEvent{request: request}
	return transaction.NewTransaction(body)
}
