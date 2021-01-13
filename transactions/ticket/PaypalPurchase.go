package ticket

import (
	"context"
	"log"

	ticketDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/ticket"
	buyerdataGW "github.com/PabloGamiz/SafeEvents-Backend/gateway/buyerdata"
	"github.com/PabloGamiz/SafeEvents-Backend/model/buyerdata"
	buyerdataMOD "github.com/PabloGamiz/SafeEvents-Backend/model/buyerdata"
	"github.com/PabloGamiz/SafeEvents-Backend/model/session"
	"github.com/PabloGamiz/SafeEvents-Backend/model/ticket"
	ticketMOD "github.com/PabloGamiz/SafeEvents-Backend/model/ticket"
)

// txPaypalPurchase represents an
type txPaypalPurchase struct {
	request       ticketDTO.PayPalPurchaseRequestDTO
	sessCtrl      session.Controller
	buyerdataCtrl buyerdata.Controller
	ticketCtrl    ticket.Controller
	ctx           context.Context
}

func (tx *txPaypalPurchase) uncoderequest(ctx context.Context) (err error) {
	var buyd = &buyerdataMOD.BuyerData{
		TicketID:             tx.request.TicketID,
		TotalAmount:          tx.request.TotalAmount,
		SubTotalAmount:       tx.request.SubTotalAmount,
		ShippingCost:         tx.request.ShippingCost,
		ShippingDiscountCost: tx.request.ShippingDiscountCost,
		FirstName:            tx.request.UserFirstName,
		LastName:             tx.request.UserLastName,
		AddressCity:          tx.request.AddressCity,
		AddressStreet:        tx.request.AddressStreet,
		AddressZipCode:       tx.request.AddressZipCode,
		AddressCountry:       tx.request.AddressCountry,
		AddressState:         tx.request.AddressState,
		AddressPhoneNumber:   tx.request.AddressPhoneNumber,
	}
	gw := buyerdataGW.NewBuyerDataGateway(ctx, buyd)
	if err = gw.Insert(); err != nil {
		return
	}
	return nil
}

// Precondition validates the transaction is ready to run
func (tx *txPaypalPurchase) Precondition() (err error) {
	// make sure the session exists
	tx.sessCtrl, err = session.GetSessionByID(tx.request.Cookie)
	return
}

// Postcondition creates new user and a opens its first session
func (tx *txPaypalPurchase) Postcondition(ctx context.Context) (v interface{}, err error) {
	log.Printf("Got a Paypal Purchase request from client %v", tx.sessCtrl.GetID())

	//make sure the ticket exists
	if tx.ticketCtrl, err = ticketMOD.GetTicketByID(tx.request.TicketID); err != nil {
		return
	}
	log.Print("Mostra DTO", tx.request)
	err = tx.uncoderequest(ctx)
	var response string
	if err == nil {
		response = "Everything Fine"
	}
	return response, nil
}

// Commit commits the transaction result
func (tx *txPaypalPurchase) Commit() (err error) {
	return err
}

// Rollback rollbacks any change caused while the transaction
func (tx *txPaypalPurchase) Rollback() {

}
