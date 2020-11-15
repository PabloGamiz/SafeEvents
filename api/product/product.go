package product

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	productDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/product"
	"github.com/PabloGamiz/SafeEvents-Backend/transactions/product"
)

// HandleAddProductRequest attends an add product request
func HandleAddProductRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handlering a Signin request")

	// Expected data for a Signup request
	var productDTO productDTO.DTO
	if err := json.NewDecoder(r.Body).Decode(&productDTO); err != nil {
		// If some error just happened it means the provided Json does not match with the expected DTO
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Setting up TxSignin with the required values
	txAddProduct := product.NewTxAddProduct(productDTO)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel() // ensures the context is canceled, at least once, at the end of this function

	txAddProduct.Execute(ctx)
	result, err := txAddProduct.Result()

	if err != nil {
		// If err != nil it means the transaction has failed
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	// Sending response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
