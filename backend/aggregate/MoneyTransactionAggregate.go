package aggregate

import (
	error2 "github.com/BambooTuna/quest-market/backend/error"
	"github.com/BambooTuna/quest-market/backend/model/transaction"
)

type MoneyTransactionAggregate struct {
	AccountId string
	Balance   int64
}

func (p *MoneyTransactionAggregate) ReceiveRecover(data []*transaction.MoneyTransaction) {
	for _, v := range data {
		p.SendTransaction(v)
	}
}

func (p *MoneyTransactionAggregate) SendTransaction(t *transaction.MoneyTransaction) error {
	if t.AccountId != p.AccountId {
		return error2.Error("AggregateId !=")
	}
	moneyDelta := t.RealPart * t.TransactionType.ToNumber()
	if p.Balance+moneyDelta >= 0 {
		p.Balance += moneyDelta
		return nil
	}
	return error2.Error(error2.LackOfMoney)
}
