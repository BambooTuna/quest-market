package dao

import "context"
import model "github.com/BambooTuna/quest-market/model"

type AccountCredentialsDao interface {
	ResolveByAccountId(ctx context.Context, accountId string) model.AccountCredentials
}
