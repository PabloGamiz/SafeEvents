package client

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	clientDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/client"
	"github.com/PabloGamiz/SafeEvents-Backend/transactions/client"
)

// HandleStatusRequest attends a signin request
func HandleStatusRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handlering a status update request")

	// Expected data for a Signup request
	var statusDTO clientDTO.StatusRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&statusDTO); err != nil {
		// If some error just happened it means the provided Json does not match with the expected DTO
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	statusDTO.Unix = time.Unix(statusDTO.Date, 0)

	// Setting up TxSignin with the required values
	txSignin := client.NewTxStatus(statusDTO)
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
