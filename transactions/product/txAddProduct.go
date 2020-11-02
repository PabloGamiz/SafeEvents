package product

import (
	"context"

	productDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/product"
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
	// SESSION //
	// var prodRepository productRepository.Repository

	// if response, err = prodRepository.AddProduct(); err == nil {
	// 	return
	// }
	return
}

// Commit commits the transaction result
func (tx *txAddProduct) Commit() error {
	return nil
}

// Rollback rollbacks any change caused while the transaction
func (tx *txAddProduct) Rollback() {

}
