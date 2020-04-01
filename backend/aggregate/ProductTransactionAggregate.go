package aggregate

import (
	"errors"
	"github.com/BambooTuna/quest-market/backend/model/transaction"
)

type ProductTransactionAggregate struct {
	ProductId   string
	Transaction *transaction.ProductTransaction
}

func (p *ProductTransactionAggregate) ReceiveRecover(data []*transaction.ProductTransaction) {
	for _, v := range data {
		p.SendTransaction(v)
	}
}

func (p *ProductTransactionAggregate) SendTransaction(t *transaction.ProductTransaction) error {
	var transactionType *transaction.ProductTransactionType
	if p.Transaction != nil {
		transactionType = &p.Transaction.TransactionType
	}

	if t.TransactionType.CanOverwrite(transactionType) && p.Transaction.PurchaserAccountId == t.PurchaserAccountId {
		p.Transaction = t
		return nil
	} else {
		return errors.New("SendTransaction: エラー")
	}
}
