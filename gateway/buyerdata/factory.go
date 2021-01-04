package buyerdata

import (
	"context"

	"github.com/PabloGamiz/SafeEvents-Backend/model/buyerdata"
)

// NewBuyerDataGateway builds a gateway for the provided ticket
func NewBuyerDataGateway(ctx context.Context, buyerdata buyerdata.Controller) Gateway {
	return &buyerdataGateway{Controller: buyerdata, ctx: ctx}
}
