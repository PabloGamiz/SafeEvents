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

// HandlePOSTFeedbackRequest attends a POST feedback request
func HandlePOSTFeedbackRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handlering a POST Feedback request")

	// Expected data for a POST feedback request
	var feedbackRequestDTO feedbackDTO.RequestDTO
	if err := json.NewDecoder(r.Body).Decode(&feedbackRequestDTO); err != nil {
		// If some error just happened it means the provided Json does not match with the expected DTO
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Setting up TxPurchase with the required values
	txPOSTFeedback := feedback.NewTxPOSTFeedback(feedbackRequestDTO)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel() // ensures the context is canceled, at least once, at the end of this function

	txPOSTFeedback.Execute(ctx)
	result, err := txPOSTFeedback.Result()

	if err != nil {
		// If err != nil it means the transaction has failed
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	// Sending response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
