package dao

import "context"
import "github.com/BambooTuna/quest-market/model/account"

type AccountCredentialsDaoImpl struct {
}

func (a AccountCredentialsDaoImpl) ResolveByAccountId(ctx context.Context, accountId string) (*account.AccountCredentials, error) {
	return nil, nil
}

func (a AccountCredentialsDaoImpl) ResolveByMail(ctx context.Context, mail string) (*account.AccountCredentials, error) {
	return nil, nil
}
func (a AccountCredentialsDaoImpl) Insert(ctx context.Context, record *account.AccountCredentials) error {
	return nil
}
func (a AccountCredentialsDaoImpl) UpdateMail(ctx context.Context, newMail string) error {
	return nil
}
func (a AccountCredentialsDaoImpl) UpdatePassword(ctx context.Context, newPassword string) error {
	return nil
}
func (a AccountCredentialsDaoImpl) DeleteByAccountId(ctx context.Context, accountId string) error {
	return nil
}
