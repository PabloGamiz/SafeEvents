package radar

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	radarDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/radar"
	"github.com/PabloGamiz/SafeEvents-Backend/transactions/radar"
)

// HandleActivateRequest attends an activation request
func HandleActivateRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handlering a Radar activation request")

	// Expected data for a radar activation request
	var activateDTO radarDTO.ActivateRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&activateDTO); err != nil {
		// If some error just happened it means the provided Json does not match with the expected DTO
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Setting up TxSignin with the required values
	txSignin := radar.NewTxActivate(activateDTO)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel() // ensures the context is canceled, at least once, at the end of this function

	txSignin.Execute(ctx)
	result, err := txSignin.Result()

	if err != nil {
		// If err != nil it means the transaction has failed
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	// Sending response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
