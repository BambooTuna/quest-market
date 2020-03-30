package command

import (
	"github.com/BambooTuna/quest-market/backend/model/account"
)

type SignUpRequestCommand struct {
	Mail     string `json:"mail"`
	Password string `json:"pass"`
}

func (c *SignUpRequestCommand) CreateAccountCredentials() (*account.AccountCredentials, error) {
	return account.GenerateAccountCredentials(c.Mail, c.Password)
}
