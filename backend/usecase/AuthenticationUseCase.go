package usecase

import (
	"context"
	"github.com/BambooTuna/quest-market/backend/command"
	"github.com/BambooTuna/quest-market/backend/dao"
	error2 "github.com/BambooTuna/quest-market/backend/error"
	"github.com/BambooTuna/quest-market/backend/model/account"
)

type AuthenticationUseCase struct {
	AccountCredentialsDao dao.AccountCredentialsDao
}

func (authenticationUseCase *AuthenticationUseCase) SignUp(ctx context.Context, c *command.SignUpRequestCommand) (*account.AccountCredentials, error) {
	accountCredentials, err := c.CreateAccountCredentials()
	if err != nil {
		return nil, err
	}

	if err := authenticationUseCase.AccountCredentialsDao.Insert(ctx, accountCredentials); err != nil {
		return nil, error2.Error(error2.DuplicateRegistration)
	}
	return accountCredentials, nil
}

func (authenticationUseCase *AuthenticationUseCase) SignIn(ctx context.Context, c *command.SignInRequestCommand) (*account.AccountCredentials, error) {
	record, err := authenticationUseCase.AccountCredentialsDao.ResolveByMail(ctx, c.Mail)
	if err != nil {
		return nil, error2.Error(error2.AccountNotFound)
	}

	accountCredentials, err := record.Authentication(c.Password)
	if err != nil {
		return nil, error2.Error(error2.AuthenticationFailed)
	}
	return accountCredentials, nil
}
