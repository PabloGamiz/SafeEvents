package event

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	eventDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/event"
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
	txPublicaEvent := event.NewTxPublicaEvent(eventDTO)
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

func HandleAllEventsRequest(w http.ResponseWriter, r *http.Request) {
	events, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, events)
}
