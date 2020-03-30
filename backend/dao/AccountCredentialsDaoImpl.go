package dao

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
)
import "github.com/BambooTuna/quest-market/backend/model/account"

type AccountCredentialsDaoImpl struct {
	DBSession *gorp.DbMap
}

func (a AccountCredentialsDaoImpl) ResolveByAccountId(ctx context.Context, accountId string) (*account.AccountCredentials, error) {
	var accountCredentials *account.AccountCredentials
	err := a.DBSession.SelectOne(&accountCredentials, fmt.Sprintf("select * from account_credentials where account_id = '%s'", accountId))
	return accountCredentials, err
}

func (a AccountCredentialsDaoImpl) ResolveByMail(ctx context.Context, mail string) (*account.AccountCredentials, error) {
	var accountCredentials *account.AccountCredentials
	err := a.DBSession.SelectOne(&accountCredentials, fmt.Sprintf("select * from account_credentials where mail = '%s'", mail))
	return accountCredentials, err
}

func (a AccountCredentialsDaoImpl) Insert(ctx context.Context, record *account.AccountCredentials) error {
	return a.DBSession.Insert(record)
}

func (a AccountCredentialsDaoImpl) Update(ctx context.Context, record *account.AccountCredentials) (int64, error) {
	return a.DBSession.Update(record)
}

func (a AccountCredentialsDaoImpl) DeleteByAccountId(ctx context.Context, accountId string) (int64, error) {
	accountCredentials, err := a.ResolveByAccountId(ctx, accountId)
	if err != nil {
		return 0, nil
	}
	return a.DBSession.Delete(accountCredentials)
}
