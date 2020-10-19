package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewServer builds a brand new server
func NewServer() Server {
	api := &api{}

	router := mux.NewRouter()
	router.HandleFunc("/signup", api.handleSignupRequest).Methods(http.MethodPost)
	//router.HandleFunc("/client/{ID:[a-zA-Z0-9_]+}", api.getClient).Methods(http.MethodGet)

	api.router = router
	return api
}
