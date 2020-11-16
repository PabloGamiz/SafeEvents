package product

import (
	"context"

	productDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/product"
	productRepo "github.com/PabloGamiz/SafeEvents-Backend/repository/product"
)

// txSignin represents an
type txAddProduct struct {
	product productDTO.DTO
}

// Precondition validates the transaction is ready to run
func (tx *txAddProduct) Precondition() (err error) {
	return
}

// Postcondition creates new user and a opens its first session
func (tx *txAddProduct) Postcondition(ctx context.Context) (v interface{}, err error) {
	// TODO: Convertir productDTO a entity i passar-lo a AddProduct

	v, err = productRepo.Repository.AddProduct(tx.product)
	return
}

// Commit commits the transaction result
func (tx *txAddProduct) Commit() error {
	return nil
}

// Rollback rollbacks any change caused while the transaction
func (tx *txAddProduct) Rollback() {

}
