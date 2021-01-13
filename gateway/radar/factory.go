package radar

import (
	"context"

	"github.com/PabloGamiz/SafeEvents-Backend/model/radar/interaction"
)

// NewInteractionGateway builds a gateway for the provided client
func NewInteractionGateway(ctx context.Context, interaction interaction.Controller) Gateway {
	return &interactionGateway{Controller: interaction, ctx: ctx}
}
