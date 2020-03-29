package dao

import "context"
import model "github.com/BambooTuna/quest-market/model"

type AccountCredentialsDao interface {
	ResolveByAccountId(ctx context.Context, accountId string) (*model.AccountCredentials, error)
	ResolveByMail(ctx context.Context, mail string) (*model.AccountCredentials, error)
	Insert(ctx context.Context, record *model.AccountCredentials) error
	UpdateMail(ctx context.Context, newMail string) error
	UpdatePassword(ctx context.Context, newPassword string) error
	DeleteByAccountId(ctx context.Context, accountId string) error
}
