package client

import (
	clientDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/client"
	"github.com/alvidir/util/pattern/transaction"
)

// NewTxSignin builds a brand new transaction for Signin
func NewTxSignin(request clientDTO.SigninRequestDTO) transaction.Tx {
	body := &txSignin{request: request}
	return transaction.NewTransaction(body)
}

// NewTxLogout builds a brand new transaction for Signin
func NewTxLogout(request clientDTO.LogoutRequestDTO) transaction.Tx {
	body := &txLogout{request: request}
	return transaction.NewTransaction(body)
}

// NewTxClientInfo builds a brand new transaction for client info consulting
func NewTxClientInfo(request clientDTO.ClientInfoRequestDTO) transaction.Tx {
	body := &txClientInfo{request: request}
	return transaction.NewTransaction(body)
}

// NewTxAddFav builds a brand new transaction for Adding a Favorite Event
func NewTxAddFav(request clientDTO.FavDTO) transaction.Tx {
	body := &txAddFav{request: request}
	return transaction.NewTransaction(body)
}

// NewTxDelFav builds a brand new transaction for Deleting a Favorite Event
func NewTxDelFav(request clientDTO.FavDTO) transaction.Tx {
	body := &txDelFav{request: request}
	return transaction.NewTransaction(body)
}
