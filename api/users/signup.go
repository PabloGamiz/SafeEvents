package users

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	clientDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/client"
	tx "github.com/PabloGamiz/SafeEvents-Backend/transactions"
)

func HandleSignupRequest(w http.ResponseWriter, r *http.Request) {
	// Expected data for a Signup request
	var signupRequest clientDTO.SignupRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&signupRequest); err != nil {
		// If some error just happened it means the provided Json does not match with the expected DTO
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Setting up TxSignup with the required values
	txSignup := tx.NewTxSignup(signupRequest)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel() // ensures the context is canceled, at least once, at the end of this function

	txSignup.Execute(ctx)
	result, err := txSignup.Result()

	if err != nil {
		// If err != nil it means the transaction has failed
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	// Sending response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
