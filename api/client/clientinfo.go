package client

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	clientDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/client"
	"github.com/PabloGamiz/SafeEvents-Backend/transactions/client"
)

func buildClientInfoRequestDTO(id uint, cookie string) clientDTO.ClientInfoRequestDTO {
	return clientDTO.ClientInfoRequestDTO{
		ID:     id,
		Cookie: cookie,
	}
}

// HandleClientInfoRequest ...
func HandleClientInfoRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handlering a Client Info request")

	cookie := r.Header.Get("Authorization")

	//Get the id from the URL
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		log.Printf("Error no id found")
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	uid := uint(id)

	req := buildClientInfoRequestDTO(uid, cookie)
	log.Printf("Handlering a Client Info request for client %d", uid)

	//Setting up TxClientInfo with the required values
	txClientInfo := client.NewTxClientInfo(req)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	txClientInfo.Execute(ctx)
	result, err := txClientInfo.Result()

	if err != nil {
		//Transaction has failed
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	//sending response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
