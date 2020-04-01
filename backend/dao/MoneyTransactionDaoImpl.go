package dao

import (
	"fmt"
	"github.com/BambooTuna/quest-market/backend/model/transaction"
	"gopkg.in/gorp.v1"
)

type MoneyTransactionDaoImpl struct {
	DBSession *gorp.DbMap
}

func (a MoneyTransactionDaoImpl) Insert(record *transaction.MoneyTransaction) error {
	return a.DBSession.Insert(record)
}

func (a MoneyTransactionDaoImpl) ResolveAllByAccountId(accountId string) ([]*transaction.MoneyTransaction, error) {
	var accountCredentials []*transaction.MoneyTransaction
	_, err := a.DBSession.Select(&accountCredentials, fmt.Sprintf("select * from money_transaction where account_id = '%s' ORDER BY transaction_id ASC", accountId))
	return accountCredentials, err
}
