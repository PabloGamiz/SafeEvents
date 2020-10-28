package users

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	clientDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/client"
	"github.com/PabloGamiz/SafeEvents-Backend/transactions/users"
)

// HandleSigninRequest attends a signin request
func HandleSigninRequest(w http.ResponseWriter, r *http.Request) {
	// Expected data for a Signup request
	var signinDTO clientDTO.SigninRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&signinDTO); err != nil {
		// If some error just happened it means the provided Json does not match with the expected DTO
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Setting up TxSignin with the required values
	txSignin := users.NewTxSignin(signinDTO)
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
