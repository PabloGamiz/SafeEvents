package product

import (
	productDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/product"
	"github.com/alvidir/util/pattern/transaction"
)

// NewTxAddProduct builds a brand new transaction for Add product
func NewTxAddProduct(product productDTO.DTO) transaction.Tx {
	body := &txAddProduct{product: product}
	return transaction.NewTransaction(body)
}
