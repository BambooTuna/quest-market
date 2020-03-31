package controller

import (
	"github.com/BambooTuna/quest-market/backend/command"
	"github.com/BambooTuna/quest-market/backend/lib/session"
	"github.com/BambooTuna/quest-market/backend/model"
	"github.com/BambooTuna/quest-market/backend/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthenticationController struct {
	Session               session.Session
	AuthenticationUseCase usecase.AuthenticationUseCase
}

func (c *AuthenticationController) SignUpRoute() func(*gin.Context) {
	return c.Session.SetSession(func(ctx *gin.Context) (*string, error) {
		var signUpRequestCommand command.SignUpRequestCommand
		if err := ctx.BindJSON(&signUpRequestCommand); err != nil {
			return nil, err
		}
		accountCredentials, err := c.AuthenticationUseCase.SignUp(ctx, &signUpRequestCommand)
		if err != nil {
			return nil, err
		}
		r := model.AccountSessionToken{AccountId: accountCredentials.AccountId, Cooperation: ""}.ToString()
		ctx.Status(http.StatusOK)
		return &r, nil
	})
}

func (c *AuthenticationController) SignInRoute() func(*gin.Context) {
	return c.Session.SetSession(func(ctx *gin.Context) (*string, error) {
		var signInRequestCommand command.SignInRequestCommand
		if err := ctx.BindJSON(&signInRequestCommand); err != nil {
			return nil, err
		}
		accountCredentials, err := c.AuthenticationUseCase.SignIn(ctx, &signInRequestCommand)
		if err != nil {
			return nil, err
		}
		r := model.AccountSessionToken{AccountId: accountCredentials.AccountId, Cooperation: ""}.ToString()
		ctx.Status(http.StatusOK)
		return &r, nil
	})
}

func (c *AuthenticationController) HealthRoute() func(*gin.Context) {
	return c.Session.RequiredSession(func(ctx *gin.Context, token *string) {
		//accountSessionToken := model.DecodeToAccountSessionToken(*token)
		//println("AccountId: " + accountSessionToken.AccountId)
		ctx.Status(http.StatusOK)
	})
}

func (c *AuthenticationController) SignOutRoute() func(*gin.Context) {
	return c.Session.InvalidateSession(func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})
}
