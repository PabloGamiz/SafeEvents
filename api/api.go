package api

import (
	"context"
	"net/http"

	"github.com/PabloGamiz/SafeEvents-Backend/transactions"
)

type api struct {
	router http.Handler
}

func (api *api) handleSignupRequest(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.TODO(), Timeout)
	defer cancel()

	txSignup := transactions.NewTxSignup()
	txSignup.Execute(ctx)

	_, err := txSignup.Result()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
}

func (api *api) Router() http.Handler {
	return api.router
}
