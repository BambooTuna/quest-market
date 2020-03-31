package account

import (
	error2 "github.com/BambooTuna/quest-market/backend/error"
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

	if plainPass == "" {
		return nil, error2.Error(error2.ValidateError("Password", "required"))
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
	var errorMessages []error2.CustomError
	if err := validate.Struct(accountCredentials); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errorMessages = append(errorMessages, error2.ValidateError(err.Field(), err.Tag()))
		}
		return nil, error2.Errors(errorMessages)
	}
	return accountCredentials, nil
}
