package client

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	clientDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/client"
	"github.com/PabloGamiz/SafeEvents-Backend/transactions/client"
)

// HandleClientAddFavRequest ...
func HandleClientAddFavRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handlering an Add Fav Request")

	var FavDTO clientDTO.ClientFavDTO
	if err := json.NewDecoder(r.Body).Decode(&FavDTO); err != nil {
		// If some error just happened it means the provided Json does not match with the expected DTO
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Setting up TxAddFav with the required values
	txAddFav := client.NewTxAddFav(FavDTO)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	txAddFav.Execute(ctx)
	result, err := txAddFav.Result()

	if err != nil {
		//Transaction has failed
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	//sending response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
