package aggregate

import (
	"github.com/BambooTuna/quest-market/backend/dao"
	"github.com/BambooTuna/quest-market/backend/model/transaction"
)

type ProductTransactionAggregates struct {
	ProductTransactionDao dao.ProductTransactionDao
	Aggregates            map[string]*ProductTransactionAggregate
}

func (p *ProductTransactionAggregates) RecoverAll() {
	productTransactionsTransactions, err := p.ProductTransactionDao.ResolveAll()
	if err != nil {
		println(err.Error())
		return
	}
	p.Aggregates = map[string]*ProductTransactionAggregate{}
	for _, t := range productTransactionsTransactions {
		if aggregate, ok := p.Aggregates[t.ProductId]; ok {
			aggregate.ReceiveRecover([]*transaction.ProductTransaction{t})
		} else {
			aggregate := ProductTransactionAggregate{ProductId: t.ProductId, Transaction: nil}
			aggregate.ReceiveRecover([]*transaction.ProductTransaction{t})
			p.Aggregates[t.ProductId] = &aggregate
		}
	}
}

func (p *ProductTransactionAggregates) SendTransaction(t *transaction.ProductTransaction) error {
	aggregate, err := p.GetAggregateByProductId(t.ProductId)
	if err != nil {
		return err
	}

	if err := aggregate.SendTransaction(t); err != nil {
		return err
	}

	if err := p.ProductTransactionDao.Insert(t); err != nil {
		return err
	}
	return nil
}

func (p *ProductTransactionAggregates) GetTransaction(productId string) (transaction.ProductTransaction, error) {
	aggregate, err := p.GetAggregateByProductId(productId)
	if err != nil {
		return transaction.ProductTransaction{}, err
	}
	return *aggregate.Transaction, nil
}

func (p *ProductTransactionAggregates) GetTransactionByAccountId(accountId string) []*transaction.ProductTransaction {
	var transactions []*transaction.ProductTransaction
	for _, value := range p.Aggregates {
		t := value.Transaction
		if t != nil {
			if t.PurchaserAccountId == accountId || t.SellerAccountId == accountId {
				transactions = append(transactions, t)
			}
		}
	}
	return transactions
}

func (p *ProductTransactionAggregates) GetAggregateByProductId(productId string) (*ProductTransactionAggregate, error) {
	if value, ok := p.Aggregates[productId]; ok {
		p.Aggregates[productId] = value
	} else {
		productTransactions, err := p.ProductTransactionDao.ResolveAllByProductId(productId)
		if err != nil {
			return nil, err
		}
		aggregate := ProductTransactionAggregate{ProductId: productId, Transaction: nil}
		aggregate.ReceiveRecover(productTransactions)
		p.Aggregates[productId] = &aggregate
	}
	return p.Aggregates[productId], nil
}
