package api

import (
	"net/http"

	"github.com/PabloGamiz/SafeEvents-Backend/api/users"
	"github.com/gorilla/mux"
)

// NewServer builds a brand new server
func NewServer() Server {
	api := &api{}

	router := mux.NewRouter()
	router.HandleFunc("/signin", users.HandleSigninRequest).Methods(http.MethodPost)
	//router.HandleFunc("/client/{ID:[a-zA-Z0-9_]+}", api.getClient).Methods(http.MethodGet)

	api.router = router
	return api
}
