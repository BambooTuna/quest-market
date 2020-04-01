package aggregate

import (
	"github.com/BambooTuna/quest-market/backend/dao"
	"github.com/BambooTuna/quest-market/backend/model/transaction"
)

type MoneyTransactionAggregates struct {
	MoneyTransactionDao dao.MoneyTransactionDao
	Aggregates          map[string]*MoneyTransactionAggregate
}

func (m *MoneyTransactionAggregates) SendTransaction(t *transaction.MoneyTransaction) error {
	aggregate, err := m.Init(t.AccountId)
	if err != nil {
		return err
	}

	if err := aggregate.SendTransaction(t); err != nil {
		return err
	}

	if err := m.MoneyTransactionDao.Insert(t); err != nil {
		return err
	}
	return nil
}

func (m *MoneyTransactionAggregates) GetBalance(accountId string) (int64, error) {
	println(accountId)
	aggregate, err := m.Init(accountId)
	println(aggregate)
	if err != nil {
		return 0, err
	}
	return aggregate.Balance, nil
}

func (m *MoneyTransactionAggregates) Init(accountId string) (*MoneyTransactionAggregate, error) {
	if value, ok := m.Aggregates[accountId]; ok {
		m.Aggregates[accountId] = value
	} else {
		moneyTransactions, err := m.MoneyTransactionDao.ResolveAllByAccountId(accountId)
		if err != nil {
			return nil, err
		}
		aggregate := MoneyTransactionAggregate{AccountId: accountId, Balance: 0}
		aggregate.ReceiveRecover(moneyTransactions)
		m.Aggregates[accountId] = &aggregate
	}
	return m.Aggregates[accountId], nil
}
