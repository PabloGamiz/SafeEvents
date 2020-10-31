package event

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/PabloGamiz/SafeEvents-Backend/transactions/event"
)

// HandleListEventsRequest attends a list events request
func HandleListEventsRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handlering a List Events request")

	// Setting up TxSignin with the required values
	txListEvents := event.NewTxListEvents()
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel() // ensures the context is canceled, at least once, at the end of this function

	txListEvents.Execute(ctx)
	result, err := txListEvents.Result()

	if err != nil {
		// If err != nil it means the transaction has failed
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	// Sending response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
