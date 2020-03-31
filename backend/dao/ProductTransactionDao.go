package dao

import (
	"context"
	"github.com/BambooTuna/quest-market/backend/model/transaction"
)

type ProductTransactionDao interface {
	Insert(ctx context.Context, record *transaction.ProductTransaction) error
	ResolveAllByAccountId(ctx context.Context, accountId string) ([]*transaction.ProductTransaction, error)
}
