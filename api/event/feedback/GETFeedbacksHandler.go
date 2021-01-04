package feedback

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	feedbackDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/feedback"
	"github.com/PabloGamiz/SafeEvents-Backend/transactions/feedback"
)

// HandleGETFeedbacksRequest attends a GET feedback request
func HandleGETFeedbacksRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handlering a GET Feedbacks request")

	// Expected data for a POST feedback request
	var feedbackRequestDTO feedbackDTO.RequestDTO
	if err := json.NewDecoder(r.Body).Decode(&feedbackRequestDTO); err != nil {
		// If some error just happened it means the provided Json does not match with the expected DTO

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Setting up TxPurchase with the required values
	txGETFeedbacks := feedback.NewTxGETFeedbacks(feedbackRequestDTO)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel() // ensures the context is canceled, at least once, at the end of this function

	txGETFeedbacks.Execute(ctx)
	result, err := txGETFeedbacks.Result()

	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	// Sending response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
