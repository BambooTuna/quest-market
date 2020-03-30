package dao

import "context"
import "github.com/BambooTuna/quest-market/backend/model/account"

type AccountCredentialsDao interface {
	ResolveByAccountId(ctx context.Context, accountId string) (*account.AccountCredentials, error)
	ResolveByMail(ctx context.Context, mail string) (*account.AccountCredentials, error)
	Insert(ctx context.Context, record *account.AccountCredentials) error
	Update(ctx context.Context, record *account.AccountCredentials) (int64, error)
	DeleteByAccountId(ctx context.Context, accountId string) (int64, error)
}
