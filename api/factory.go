package api

import (
	"net/http"

	"github.com/PabloGamiz/SafeEvents-Backend/api/client"
	"github.com/PabloGamiz/SafeEvents-Backend/api/event"
	"github.com/gorilla/mux"
)

// NewServer builds a brand new server
func NewServer() Server {
	api := &api{}

	router := mux.NewRouter()

	// Client router Handlers
	router.HandleFunc(client.APISigninPath, client.HandleSigninRequest).Methods(http.MethodPost)
	router.HandleFunc(client.APILogoutPath, client.HandleLogoutRequest).Methods(http.MethodPut)
	router.HandleFunc(client.APIClientInfoPath, client.HandleClientInfoRequest).Methods(http.MethodPut)
	router.HandleFunc(client.APIAddFavPath, client.HandleClientAddFavRequest).Methods(http.MethodPost)
	router.HandleFunc(client.APIDelFavPath, client.HandleClientDelFavRequest).Methods(http.MethodPost)

	// Events router Handlers
	router.HandleFunc(event.APIListEvents, event.HandleListEventsRequest).Methods(http.MethodGet)
	router.HandleFunc(event.APIPubliEvent, event.HandlePublicaEventRequest).Methods(http.MethodPost)
	router.HandleFunc(event.APIGetEvent, event.HandleGetEventRequest).Methods(http.MethodPost)

	api.router = router
	return api
}
