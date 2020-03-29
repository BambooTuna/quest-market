package usecase

import (
	"context"
	"github.com/BambooTuna/quest-market/dao"
	"github.com/BambooTuna/quest-market/model"
)

type AuthenticationUseCase struct {
	AccountCredentialsDao dao.AccountCredentialsDao
}

func (authenticationUseCase *AuthenticationUseCase) SinUp(ctx context.Context, u model.AccountCredentials) (err error) {
	_, err := authenticationUseCase.AccountCredentialsDao.ResolveByAccountId(ctx, "")
	return
}
