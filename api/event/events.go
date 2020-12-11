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

// HandleListEventsByTypeRequest attends a list events by type request
func HandleListEventsByTypeRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handlering a List Events by type request")

	// Expected data for a Publica request
	var requestDTO eventDTO.ListEventsByTypeRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&requestDTO); err != nil {
		// If some error just happened it means the provided Json does not match with the expected DTO
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Setting up TxSignin with the required values
	TxListEventsByType := event.NewTxListEventsByType(requestDTO)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel() // ensures the context is canceled, at least once, at the end of this function

	TxListEventsByType.Execute(ctx)
	result, err := TxListEventsByType.Result()
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
	var publicaDTO eventDTO.PublicaEvent
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

// HandleModificaEventRequest attends a Modifica Esdeveniment request
func HandleModificaEventRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handlering a Modifica Esdeveniment request")

	// Expected data for a Publica request
	var modificaDTO eventDTO.DTO
	if err := json.NewDecoder(r.Body).Decode(&modificaDTO); err != nil {
		// If some error just happened it means the provided Json does not match with the expected DTO
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Setting up TxPublicaEvent with the required values
	txModificaEvent := event.NewTxModificaEvent(modificaDTO)

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel() // ensures the context is canceled, at least once, at the end of this function

	txModificaEvent.Execute(ctx)
	result, err := txModificaEvent.Result()

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

	var getDTO eventDTO.GetEvent
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

func buildListFavoritesRequestDTO(cookie string) eventDTO.ListFavoritesRequestDTO {
	return eventDTO.ListFavoritesRequestDTO{
		Cookie: cookie,
	}
}

// HandleListFavoritesRequest attends a list of favorites events request
func HandleListFavoritesRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handlering a List Favorites request")

	cookie := r.Header.Get("Authorization")

	req := buildListFavoritesRequestDTO(cookie)

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
