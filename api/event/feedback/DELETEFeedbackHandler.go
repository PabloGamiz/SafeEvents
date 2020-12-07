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

// HandleDELETEFeedbackRequest attends a DELETE feedback request
func HandleDELETEFeedbackRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handlering a DELETE Feedback request")

	// Expected data for a POST feedback request
	var feedbackRequestDTO feedbackDTO.RequestDTO
	if err := json.NewDecoder(r.Body).Decode(&feedbackRequestDTO); err != nil {
		// If some error just happened it means the provided Json does not match with the expected DTO
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Setting up TxPurchase with the required values
	txDELETEFeedback := feedback.NewTxDELETEFeedback(feedbackRequestDTO)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel() // ensures the context is canceled, at least once, at the end of this function

	txDELETEFeedback.Execute(ctx)
	result, err := txDELETEFeedback.Result()

	if err != nil {
		// If err != nil it means the transaction has failed
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	// Sending response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
