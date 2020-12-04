package feedback

import (
	"context"

	"github.com/PabloGamiz/SafeEvents-Backend/model/feedback"
)

// NewFeedbackGW builds a gateway for the provided feedback
func NewFeedbackGateway(ctx context.Context, feedback feedback.Controller) Gateway {
	return &feedbackGateway{Controller: feedback, ctx: ctx}
}
