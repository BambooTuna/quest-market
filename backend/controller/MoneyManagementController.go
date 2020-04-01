package controller

import (
	error2 "github.com/BambooTuna/quest-market/backend/error"
	"github.com/BambooTuna/quest-market/backend/json"
	"github.com/BambooTuna/quest-market/backend/lib/session"
	"github.com/BambooTuna/quest-market/backend/model"
	"github.com/BambooTuna/quest-market/backend/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MoneyManagementController struct {
	Session                session.Session
	MoneyManagementUseCase usecase.MoneyManagementUseCase
}

func (c *MoneyManagementController) SendMoneyRoute() func(*gin.Context) {
	return c.Session.RequiredSession(func(ctx *gin.Context, token *string) {
		var sendMoneyRequestJson json.SendMoneyRequestJson
		if err := ctx.BindJSON(&sendMoneyRequestJson); err != nil {
			ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: error2.Error(error2.BindJSONFailed).Error()})
			return
		}

		accountSessionToken := model.DecodeToAccountSessionToken(*token)
		if err := c.MoneyManagementUseCase.SendMoney(accountSessionToken.AccountId, sendMoneyRequestJson.ReceiverAccountId, sendMoneyRequestJson.Money); err != nil {
			ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: err.Error()})
			return
		}
		ctx.Status(http.StatusOK)
	})
}

func (c *MoneyManagementController) GetBalanceRoute() func(*gin.Context) {
	return c.Session.RequiredSession(func(ctx *gin.Context, token *string) {
		accountSessionToken := model.DecodeToAccountSessionToken(*token)
		balance, err := c.MoneyManagementUseCase.GetBalance(accountSessionToken.AccountId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, json.GetBalanceResponseJson{Balance: balance})
	})
}
