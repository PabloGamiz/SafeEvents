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

// HandlePublicaEventRequest attends a Publica Esdeveniment request
func HandlePublicaEventRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handlering a Publica Esdeveniment request")

	// Expected data for a Publica request
	var publicaDTO eventDTO.DTO
	if err := json.NewDecoder(r.Body).Decode(&publicaDTO); err != nil {
		// If some error just happened it means the provided Json does not match with the expected DTO
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Setting up TxPublicaEvent with the required values
	txPublicaEvent := event.NewTxPublicaEvent(publicaDTO)

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

// HandleGetEventRequest attends a Get a single Esdeveniment request
func HandleGetEventRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handlering a single event request")

	//Obte el id passat com a parametre a la url
	/*idR, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || idR < 1 {
		http.NotFound(w, r)
		return
	}

	geteventDTO := eventDTO.DTO{
		ID: uint(idR),
	}*/

	var getDTO eventDTO.DTO
	if err := json.NewDecoder(r.Body).Decode(&getDTO); err != nil {
		// If some error just happened it means the provided Json does not match with the expected DTO
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Setting up txGetEvent with the required values
	txGetEvent := event.NewTxGetEvent(getDTO)

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel() // ensures the context is canceled, at least once, at the end of this function

	txGetEvent.Execute(ctx)
	result, err := txGetEvent.Result()

	if err != nil {
		// If err != nil it means the transaction has failed
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	// Sending response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// HandleListFavoritesRequest attends a list of favorites events request
func HandleListFavoritesRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handlering a List Favorites request")

	//Get the cookie from the request body
	var req eventDTO.ListFavoritesRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// If some error just happened it means the provided Json does not match with the expected DTO
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Setting uo TxListFavorites with the required values
	txListFavorites := event.NewTxListFavorites(req)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	txListFavorites.Execute(ctx)
	result, err := txListFavorites.Result()

	if err != nil {
		//Transaction has failed
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	//sending response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
