package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/PabloGamiz/SafeEvents-Backend/api/users"
)

// NewServer builds a brand new server
func NewServer() Server {
	api := &api{}

	router := mux.NewRouter()
	router.HandleFunc("/signup", users.HandleSignupRequest).Methods(http.MethodPost)
	//router.HandleFunc("/client/{ID:[a-zA-Z0-9_]+}", api.getClient).Methods(http.MethodGet)

	api.router = router
	return api
}
