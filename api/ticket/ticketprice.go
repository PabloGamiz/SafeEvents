package ticket

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	ticketDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/ticket"
	"github.com/PabloGamiz/SafeEvents-Backend/transactions/ticket"
)

// HandleTicketPriceRequest attends a signin request
func HandleTicketPriceRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handlering a Paypal purchase")

	// Expected data for a ticket price request
	var tickDTO ticketDTO.PriceTicketRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&tickDTO); err != nil {
		// If some error just happened it means the provided Json does not match with the expected DTO
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Setting up txTicketPrice with the required values
	txTicketPrice := ticket.NewTicketPrice(tickDTO)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel() // ensures the context is canceled, at least once, at the end of this function

	txTicketPrice.Execute(ctx)
	result, err := txTicketPrice.Result()

	if err != nil {
		// If err != nil it means the transaction has failed
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	// Sending response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
