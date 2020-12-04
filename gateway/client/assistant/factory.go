package assistant

import (
	"context"

	"github.com/PabloGamiz/SafeEvents-Backend/model/client/assistant"
)

// NewAssistantGateway builds a gateway for the provided client
func NewAssistantGateway(ctx context.Context, assistant assistant.Controller) Gateway {
	return &assistantGateway{Controller: assistant, ctx: ctx}
}
