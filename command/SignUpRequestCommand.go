package command

import (
	"github.com/BambooTuna/quest-market/model/account"
)

type SignUpRequestCommand struct {
	Mail     string
	Password string
}

func (c *SignUpRequestCommand) CreateAccountCredentials() (*account.AccountCredentials, error) {
	return account.GenerateAccountCredentials(c.Mail, c.Password)
}
