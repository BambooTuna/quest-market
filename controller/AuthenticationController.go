package controller

import (
	"github.com/BambooTuna/quest-market/command"
	"github.com/BambooTuna/quest-market/lib/session"
	"github.com/BambooTuna/quest-market/usecase"
	"github.com/gin-gonic/gin"
)

type AuthenticationController struct {
	Session               session.Session
	AuthenticationUseCase usecase.AuthenticationUseCase
}

func (c *AuthenticationController) SignUp() func(*gin.Context) {
	return c.Session.SetSession(func(ctx *gin.Context) (*string, error) {
		var signUpRequestCommand command.SignUpRequestCommand
		if err := ctx.BindJSON(&signUpRequestCommand); err != nil {
			return nil, err
		}
		accountCredentials, err := c.AuthenticationUseCase.SignUp(ctx, &signUpRequestCommand)
		if err != nil {
			return nil, err
		}
		return &accountCredentials.AccountId, nil
	})
}
