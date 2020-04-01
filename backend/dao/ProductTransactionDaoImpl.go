package dao

import (
	"fmt"
	"github.com/BambooTuna/quest-market/backend/model/transaction"
	"gopkg.in/gorp.v1"
)

type ProductTransactionDaoImpl struct {
	DBSession *gorp.DbMap
}

func (a ProductTransactionDaoImpl) Insert(record *transaction.ProductTransaction) error {
	return a.DBSession.Insert(record)
}

func (a ProductTransactionDaoImpl) ResolveAllByProductId(productId string) ([]*transaction.ProductTransaction, error) {
	var productTransaction []*transaction.ProductTransaction
	_, err := a.DBSession.Select(&productTransaction, fmt.Sprintf("select * from product_transaction where product_id = '%s' ORDER BY transaction_id ASC", productId))
	return productTransaction, err
}
