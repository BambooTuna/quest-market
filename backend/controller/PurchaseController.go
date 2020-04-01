package controller

import (
	"github.com/BambooTuna/quest-market/backend/command"
	"github.com/BambooTuna/quest-market/backend/json"
	"github.com/BambooTuna/quest-market/backend/lib/session"
	"github.com/BambooTuna/quest-market/backend/model"
	"github.com/BambooTuna/quest-market/backend/model/transaction"
	"github.com/BambooTuna/quest-market/backend/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PurchaseController struct {
	Session         session.Session
	PurchaseUseCase usecase.PurchaseUseCase
}

func (c *PurchaseController) PurchaseFlowRoute() func(*gin.Context) {
	return c.Session.RequiredSession(func(ctx *gin.Context, token *string) {
		productId := ctx.Param("productId")
		transactionType := transaction.ProductTransactionType(ctx.Query("type"))

		if transactionType == "" {
			ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: "transactionTypeが無効です"})
		} else {
			accountSessionToken := model.DecodeToAccountSessionToken(*token)

			var err error
			switch transactionType {
			case transaction.WaitingForPayment:
				command := command.PurchaseApplicationCommand{ProductId: productId, PurchaserAccountId: accountSessionToken.AccountId}
				err = c.PurchaseUseCase.Application(ctx, &command)
			case transaction.WaitingToReceive:
				command := command.PurchasePaymentCommand{ProductId: productId, PurchaserAccountId: accountSessionToken.AccountId}
				err = c.PurchaseUseCase.Payment(ctx, &command)
			case transaction.Complete:
				command := command.PurchaseReceiptConfirmationCommand{ProductId: productId, SellerAccountId: accountSessionToken.AccountId}
				err = c.PurchaseUseCase.ReceiptConfirmation(ctx, &command)
			}

			if err != nil {
				ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: err.Error()})
				return
			}
			ctx.Status(http.StatusOK)
		}
	})
}

func (c *PurchaseController) GetMyProductTransactionRoute() func(*gin.Context) {
	return c.Session.RequiredSession(func(ctx *gin.Context, token *string) {
		accountSessionToken := model.DecodeToAccountSessionToken(*token)
		ctx.JSON(http.StatusOK, json.ConvertToProductTransactionListResponseJson(c.PurchaseUseCase.GetTransactionByAccountId(accountSessionToken.AccountId)))
	})
}
