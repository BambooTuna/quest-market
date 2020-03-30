package controller

import (
	"github.com/BambooTuna/quest-market/command"
	"github.com/BambooTuna/quest-market/lib/session"
	"github.com/BambooTuna/quest-market/model"
	"github.com/BambooTuna/quest-market/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
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

func (c *AuthenticationController) SignIn() func(*gin.Context) {
	return c.Session.SetSession(func(ctx *gin.Context) (*string, error) {
		var signInRequestCommand command.SignInRequestCommand
		if err := ctx.BindJSON(&signInRequestCommand); err != nil {
			return nil, err
		}
		accountCredentials, err := c.AuthenticationUseCase.SignIn(ctx, &signInRequestCommand)
		if err != nil {
			return nil, err
		}
		return &accountCredentials.AccountId, nil
	})
}

func (c *AuthenticationController) Health() func(*gin.Context) {
	return c.Session.RequiredSession(func(ctx *gin.Context, token *string) {
		accountSessionToken := model.DecodeToAccountSessionToken(*token)
		println("AccountId: " + accountSessionToken.AccountId)
		ctx.Status(http.StatusOK)
	})
}
