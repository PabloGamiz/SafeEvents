package api

import (
	"net/http"

	"github.com/PabloGamiz/SafeEvents-Backend/api/client"
	"github.com/PabloGamiz/SafeEvents-Backend/api/events"

	"github.com/gorilla/mux"
)

// NewServer builds a brand new server
func NewServer() Server {
	api := &api{}

	router := mux.NewRouter()
	router.HandleFunc(client.APIPath, client.HandleSigninRequest).Methods(http.MethodPost)
	router.HandleFunc(event.APILISTEVENTS, event.HandleListEventsRequest).Methods(http.MethodGet)
	router.HandleFunc(event.APILISTEVENTS, event.HandlePublicaEventRequest).Methods(http.MethodPost)
	router.HandleFunc(event.APILISTEVENTS, event.HandleGetEventRequest).Methods(http.MethodGet)
	//router.HandleFunc("/client/{ID:[a-zA-Z0-9_]+}", api.getClient).Methods(http.MethodGet)

	api.router = router
	return api
}
