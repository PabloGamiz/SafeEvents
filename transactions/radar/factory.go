package radar

import (
	radarDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/radar"
	"github.com/alvidir/util/pattern/transaction"
)

// NewTxActivate builds a brand new transaction for radar activation
func NewTxActivate(request radarDTO.ActivateRequestDTO) transaction.Tx {
	body := &txActivate{request: request}
	return transaction.NewTransaction(body)
}

// NewTxInteraction builds a brand new transaction for interaction registry
func NewTxInteraction(request radarDTO.InteractionRequestDTO) transaction.Tx {
	body := &txInteraction{request: request}
	return transaction.NewTransaction(body)
}
