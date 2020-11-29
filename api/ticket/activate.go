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

// HandleActivateRequest attends a signin request
func HandleActivateRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handlering an Actiavtion request")

	// Expected data for an activation request
	var activateDTO ticketDTO.ActivateRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&activateDTO); err != nil {
		// If some error just happened it means the provided Json does not match with the expected DTO
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Setting up TxActivate with the required values
	txActivate := ticket.NewTxActivate(activateDTO)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel() // ensures the context is canceled, at least once, at the end of this function

	txActivate.Execute(ctx)
	result, err := txActivate.Result()

	if err != nil {
		// If err != nil it means the transaction has failed
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	// Sending response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
