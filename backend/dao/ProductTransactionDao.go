package dao

import (
	"github.com/BambooTuna/quest-market/backend/model/transaction"
)

type ProductTransactionDao interface {
	Insert(record *transaction.ProductTransaction) error
	ResolveAllByProductId(productId string) ([]*transaction.ProductTransaction, error)
	ResolveAll() ([]*transaction.ProductTransaction, error)
}
