package api

import (
	"net/http"

	"github.com/PabloGamiz/SafeEvents-Backend/api/events"
	"github.com/PabloGamiz/SafeEvents-Backend/api/users"
	"github.com/gorilla/mux"
)

// NewServer builds a brand new server
func NewServer() Server {
	api := &api{}

	router := mux.NewRouter()

	// Users router Handlers
	router.HandleFunc("/signup", users.HandleSignupRequest).Methods(http.MethodPost)

	// Events router Handlers
	router.HandleFunc("/events/list", events.HandleListEventsRequest).Methods(http.MethodGet)

	api.router = router
	return api
}
