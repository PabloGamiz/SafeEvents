package event

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	eventDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/event"
	"github.com/PabloGamiz/SafeEvents-Backend/transactions/event"
)

func HandlePublicaEventRequest(w http.ResponseWriter, r *http.Request) {
	// Expected data for a Event request
	var eventDTO eventDTO.EventRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&eventDTO); err != nil {
		// If some error just happened it means the provided Json does not match with the expected DTO
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Setting up TxPublicaevent with the required values
	txPublicaEvent := event.NewtxPublicaEvent(eventDTO)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel() // ensures the context is canceled, at least once, at the end of this function

	txPublicaEvent.Execute(ctx)
	result, err := txPublicaEvent.Result()

	if err != nil {
		// If err != nil it means the transaction has failed
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	// Sending response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func HandleGetEventRequest(w http.ResponseWriter, r *http.Request) {
	// Expected data for a Event request
	log.Printf("Handlering a single event request")

	txGetEvent := event.NewtxGetEvent()
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel() // ensures the context is canceled, at least once, at the end of this function

	txPublicaEvent.Execute(ctx)
	result, err := txPublicaEvent.Result()

	if err != nil {
		http.Error(w, err.Error(), htttp.StatusConflict)
		return
	}
	// Sending response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
