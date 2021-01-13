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

// HandleClientInfoRequest ...
func HandleClientInfoRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handlering a Client Info request")

	//Get the cookie and id from the request body
	var req clientDTO.ClientInfoRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// If some error just happened it means the provided Json does not match with the expected DTO
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Setting up TxClientInfo with the required values
	txClientInfo := client.NewTxClientInfo(req)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	txClientInfo.Execute(ctx)
	result, err := txClientInfo.Result()

	if err != nil {
		//Transaction has failed
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	//sending response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
