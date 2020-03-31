package aggregate

import (
	"github.com/BambooTuna/quest-market/backend/model/transaction"
)

type ProductTransactionAggregate struct {
	ProductTransactions map[string]*transaction.ProductTransaction
}

func (p *ProductTransactionAggregate) Recovery(data []*transaction.ProductTransaction) []*transaction.ProductTransaction {
	for _, v := range data {
		if value, ok := p.ProductTransactions[v.ProductId]; ok {
			p.ProductTransactions[v.ProductId] = value.Override(v)
		} else {
			p.ProductTransactions[v.ProductId] = v
		}
	}
	var result []*transaction.ProductTransaction
	for _, value := range p.ProductTransactions {
		result = append(result, value)
	}
	return result
}
