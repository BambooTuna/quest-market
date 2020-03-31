package dao

import (
	"context"
	"github.com/BambooTuna/quest-market/backend/model/transaction"
)

type MoneyTransactionDao interface {
	Insert(ctx context.Context, record *transaction.MoneyTransaction) error
	ResolveAllByAccountId(ctx context.Context, accountId string) ([]transaction.MoneyTransaction, error)
}
