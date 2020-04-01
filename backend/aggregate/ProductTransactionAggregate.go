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
		if !p.CanWriteThisTransaction(t) {
			return errors.New("SendTransaction: AccountIdが一致しない")
		}
	}

	if t.TransactionType.CanOverwrite(transactionType) {
		p.Transaction = t
		return nil
	} else {
		return errors.New("SendTransaction: CanOverwriteエラー")
	}
}

func (p *ProductTransactionAggregate) CanWriteThisTransaction(t *transaction.ProductTransaction) bool {
	if p.Transaction == nil {
		return false
	}
	switch p.Transaction.TransactionType {
	case transaction.WaitingForPayment:
		return p.Transaction.PurchaserAccountId == t.PurchaserAccountId
	case transaction.WaitingToReceive:
		return p.Transaction.PurchaserAccountId == t.PurchaserAccountId
	case transaction.Complete:
		return p.Transaction.SellerAccountId == t.SellerAccountId
	}
	return false
}
