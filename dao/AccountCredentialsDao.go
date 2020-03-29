package dao

import "context"
import "github.com/BambooTuna/quest-market/model/account"

type AccountCredentialsDao interface {
	ResolveByAccountId(ctx context.Context, accountId string) (*account.AccountCredentials, error)
	ResolveByMail(ctx context.Context, mail string) (*account.AccountCredentials, error)
	Insert(ctx context.Context, record *account.AccountCredentials) error
	UpdateMail(ctx context.Context, newMail string) error
	UpdatePassword(ctx context.Context, newPassword string) error
	DeleteByAccountId(ctx context.Context, accountId string) error
}
