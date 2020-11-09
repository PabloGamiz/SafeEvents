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

// HandleLogoutRequest attends a signin request
func HandleLogoutRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handlering a Logout request")

	// Expected data for a Signup request
	var logoutDTO clientDTO.LogoutRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&logoutDTO); err != nil {
		// If some error just happened it means the provided Json does not match with the expected DTO
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Setting up TxLogout with the required values
	txLogout := client.NewTxLogout(logoutDTO)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel() // ensures the context is canceled, at least once, at the end of this function

	txLogout.Execute(ctx)
	result, err := txLogout.Result()

	if err != nil {
		// If err != nil it means the transaction has failed
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	// Sending response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
