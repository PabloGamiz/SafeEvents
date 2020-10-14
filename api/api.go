package api

import (
	"net/http"
)

type api struct {
	router http.Handler
}

func (api *api) signup(w http.ResponseWriter, r *http.Request) {

}

func (api *api) Router() http.Handler {
	return api.router
}
