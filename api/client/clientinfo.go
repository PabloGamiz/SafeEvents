package client

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/PabloGamiz/SafeEvents-Backend/transactions/client"
)

//Com fer lo de obtindre l'email i fer un dto
//Com funciona lo del context i el timeout o si s'ha de canviar a la tansaction

func HandleClientInfoRequest(w http.ResposeWriter, r *http.Request) {
	log.Printf("Handlering a Client Info request")

	//Expected a email in the URI or a header?
	var clientInfoDTO clientDTO.clientInfoRequestDTO

	//npi de com es fa :(

	//From the uri
	string uri = r.GetRequestURI();
	int lastindex = uri.LastindexOf("=")
	string email = uri.Substring(uri.length(),lastIndex)

	//From the header
	string email = r.Header.Get("email")


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
