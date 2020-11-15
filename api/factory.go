package api

import (
	"net/http"

	"github.com/PabloGamiz/SafeEvents-Backend/api/client"
<<<<<<< HEAD
=======
	"github.com/PabloGamiz/SafeEvents-Backend/api/event"

>>>>>>> feature/event-management/publica_i_consulta_especifica
	"github.com/gorilla/mux"
)

// NewServer builds a brand new server
func NewServer() Server {
	api := &api{}

	router := mux.NewRouter()
<<<<<<< HEAD
	router.HandleFunc(client.APISigninPath, client.HandleSigninRequest).Methods(http.MethodPost)
	router.HandleFunc(client.APILogoutPath, client.HandleLogoutRequest).Methods(http.MethodPut)
	//router.HandleFunc("/client/{ID:[a-zA-Z0-9_]+}", api.getClient).Methods(http.MethodGet)

	// Events router Handlers
	// router.HandleFunc("/events/list", events.HandleListEventsRequest).Methods(http.MethodGet)
=======
	router.HandleFunc(client.APIPath, client.HandleSigninRequest).Methods(http.MethodPost)
	router.HandleFunc(event.APILISTEVENTS, event.HandleListEventsRequest).Methods(http.MethodGet)
	router.HandleFunc(event.APIPubliEvent, event.HandlePublicaEventRequest).Methods(http.MethodPost)
	router.HandleFunc(event.APIGetEvent, event.HandleGetEventRequest).Methods(http.MethodGet)

	//router.HandleFunc("/client/{ID:[a-zA-Z0-9_]+}", api.getClient).Methods(http.MethodGet)
>>>>>>> feature/event-management/publica_i_consulta_especifica

	api.router = router
	return api
}
