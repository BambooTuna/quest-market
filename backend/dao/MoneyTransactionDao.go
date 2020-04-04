package dao

import (
	"github.com/BambooTuna/quest-market/backend/model/transaction"
)

type MoneyTransactionDao interface {
	Insert(record *transaction.MoneyTransaction) error
	ResolveAllByAccountId(accountId string) ([]*transaction.MoneyTransaction, error)
}
