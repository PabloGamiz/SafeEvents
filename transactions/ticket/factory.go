package ticket

import (
	ticketDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/ticket"
	"github.com/alvidir/util/pattern/transaction"
)

// NewTxPurchase builds a brand new transaction for Signin
func NewTxPurchase(request ticketDTO.PurchaseRequestDTO) transaction.Tx {
	body := &txPurchase{request: request}
	return transaction.NewTransaction(body)
}
