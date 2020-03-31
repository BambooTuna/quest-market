package aggregate

import "github.com/BambooTuna/quest-market/backend/model/transaction"

type MoneyTransactionAggregate struct {
}

func (p *MoneyTransactionAggregate) Recovery(data []*transaction.MoneyTransaction) int64 {
	var money int64 = 0
	for _, v := range data {
		//TODO
		moneyDelta := v.RealPart * v.TransactionType.ToNumber()
		if money+moneyDelta >= 0 {
			money += moneyDelta
		}
	}
	return money
}
