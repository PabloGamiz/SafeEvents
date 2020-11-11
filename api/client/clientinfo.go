package client

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/PabloGamiz/SafeEvents-Backend/transactions/client"
)

func HandleClientInfoRequest(w http.ResposeWriter, r *http.Request) {
	log.Printf("Handlering a Client Info request")

	//Expected a email in the URI or a header?
	var clientInfoDTO clientDTO.clientInfoRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&clientInfoDTO); err != nil {
		//The json sent does not match with the expected DTO
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Setting up TxClientInfo with the required values
	txClientInfo := client.NewTxClientInfo(clientInfoDTO)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	txClientInfo.execute()
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
