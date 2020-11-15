package ticket

import (
	"context"
	"log"

	ticketDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/ticket"
	clientGW "github.com/PabloGamiz/SafeEvents-Backend/gateway/client"
	eventGW "github.com/PabloGamiz/SafeEvents-Backend/gateway/event"
)

// txPurchase represents an
type txPurchase struct {
	request ticketDTO.PurchaseRequestDTO
}

func (tx *txPurchase) buildPurchaseResponseDTO() *ticketDTO.PurchaseResponseDTO {
	return &ticketDTO.PurchaseResponseDTO{}
}

// Precondition validates the transaction is ready to run
func (tx *txPurchase) Precondition() (err error) {
	return
}

// Postcondition creates new user and a opens its first session
func (tx *txPurchase) Postcondition(ctx context.Context) (v interface{}, err error) {
	log.Printf("Got a Purchase request from client %v", tx.request.ClientID)
	var client clientGW.Gateway
	if client, err = clientGW.FindClientByID(ctx, tx.request.ClientID); err != nil {
		return
	}

	var event eventGW.Gateway
	if event, err = eventGW.FindEventByID(ctx, int(tx.request.EventID)); err != nil {
		return
	}

	response := tx.buildPurchaseResponseDTO()
	return response, nil
}

// Commit commits the transaction result
func (tx *txPurchase) Commit() error {
	return nil
}

// Rollback rollbacks any change caused while the transaction
func (tx *txPurchase) Rollback() {

}
