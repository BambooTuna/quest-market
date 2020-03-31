package account

import (
	"github.com/BambooTuna/quest-market/backend/settings"
	"github.com/go-playground/validator/v10"
)

type AccountCredentials struct {
	AccountId string `db:"account_id"`
	Mail      string `validate:"required,email" db:"mail"`
	Password  string `validate:"max=255,min=1" db:"password"`
}

func (a *AccountCredentials) Authentication(plainPass string) (*AccountCredentials, error) {
	return a, settings.VerifyPassword(a.Password, plainPass)
}

func GenerateAccountCredentials(mail, plainPass string) (*AccountCredentials, error) {
	uuid, err := settings.GenerateUUID()
	if err != nil {
		return nil, err
	}

	encryptedPass, err := settings.PasswordHash(plainPass)
	if err != nil {
		return nil, err
	}

	accountCredentials := &AccountCredentials{
		AccountId: uuid,
		Mail:      mail,
		Password:  encryptedPass,
	}
	validate := validator.New()
	if err := validate.Struct(accountCredentials); err != nil {
		return nil, err
	}
	return accountCredentials, nil
}
