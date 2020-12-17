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

// NewTxActivate builds a brand new transaction for Signin
func NewTxActivate(request ticketDTO.ActivateRequestDTO) transaction.Tx {
	body := &txActivate{request: request}
	return transaction.NewTransaction(body)
}

// NewTxGetTickets builds a brand new transaction for Signin
func NewTxGetTickets(request ticketDTO.GetTicketsRequestDTO) transaction.Tx {
	body := &txGetTickets{request: request}
	return transaction.NewTransaction(body)
}

// NewTxCheck builds a brand new transaction for Signin
func NewTxCheck(request ticketDTO.CheckRequestDTO) transaction.Tx {
	body := &txCheck{request: request}
	return transaction.NewTransaction(body)
}
