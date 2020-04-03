package controller

import (
	error2 "github.com/BambooTuna/quest-market/backend/error"
	"github.com/BambooTuna/quest-market/backend/json"
	"github.com/BambooTuna/quest-market/backend/lib/session"
	"github.com/BambooTuna/quest-market/backend/model"
	"github.com/BambooTuna/quest-market/backend/model/item"
	"github.com/BambooTuna/quest-market/backend/settings"
	"github.com/BambooTuna/quest-market/backend/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ItemContractController struct {
	Session             session.Session
	ItemContractUseCase usecase.ItemContractUseCase
}

func (i *ItemContractController) GetPublicItemContractRoute(paramKey string) func(*gin.Context) {
	return i.Session.OptionalRequiredSession(func(ctx *gin.Context, token *string) {
		itemId := ctx.Param(paramKey)
		var contractDetails *item.ContractDetails
		var err error
		var acquiredBy string
		if token != nil {
			accountSessionToken := model.DecodeToAccountSessionToken(*token)
			acquiredBy = accountSessionToken.AccountId
		}
		contractDetails, err = i.ItemContractUseCase.GetPublicItemContract(itemId, "")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, json.ConvertToContractDetailsResponseJson(contractDetails, acquiredBy))
	})
}

func (i *ItemContractController) GetPublicItemContractsRoute() func(*gin.Context) {
	return func(ctx *gin.Context) {
		page, e1 := strconv.ParseInt(ctx.DefaultQuery("page", "1"), 10, 64)
		limit, e2 := strconv.ParseInt(ctx.DefaultQuery("limit", "10"), 10, 64)
		if e1 != nil || e2 != nil {
			ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: "不正なクエリーパラメータです"})
			return
		}
		quantityLimit := settings.QuantityLimit{Page: page, Limit: limit}
		ctx.JSON(http.StatusOK, json.ConvertToContractDetailsListResponseJson(i.ItemContractUseCase.GetPublicItemContracts(quantityLimit), ""))
	}
}

func (i *ItemContractController) GetMyItemContractsRoute() func(*gin.Context) {
	return i.Session.RequiredSession(func(ctx *gin.Context, token *string) {
		accountSessionToken := model.DecodeToAccountSessionToken(*token)
		page, e1 := strconv.ParseInt(ctx.DefaultQuery("page", "1"), 10, 64)
		limit, e2 := strconv.ParseInt(ctx.DefaultQuery("limit", "10"), 10, 64)
		if e1 != nil || e2 != nil {
			ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: "不正なクエリーパラメータです"})
			return
		}
		quantityLimit := settings.QuantityLimit{Page: page, Limit: limit}
		ctx.JSON(http.StatusBadRequest, json.ConvertToContractDetailsListResponseJson(i.ItemContractUseCase.GetMyItemContracts(quantityLimit, accountSessionToken.AccountId), accountSessionToken.AccountId))
	})
}

func (i *ItemContractController) ExhibitionRoute() func(*gin.Context) {
	return i.Session.RequiredSession(func(ctx *gin.Context, token *string) {
		accountSessionToken := model.DecodeToAccountSessionToken(*token)
		var exhibitionRequestJson json.ExhibitionRequestJson
		if err := ctx.BindJSON(&exhibitionRequestJson); err != nil {
			ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: error2.Error(error2.BindJSONFailed).Error()})
			return
		}
		command := exhibitionRequestJson.GenerateExhibitionCommand(accountSessionToken.AccountId)
		if contractDetails, err := i.ItemContractUseCase.CreateNewItemContract(&command); err != nil {
			ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: err.Error()})
		} else {
			ctx.JSON(http.StatusOK, json.ConvertToContractDetailsResponseJson(contractDetails, accountSessionToken.AccountId))
		}
	})
}

func (i *ItemContractController) PurchaseItemRoute(paramKey string) func(*gin.Context) {
	return i.Session.RequiredSession(func(ctx *gin.Context, token *string) {
		itemId := ctx.Param(paramKey)
		accountSessionToken := model.DecodeToAccountSessionToken(*token)
		if err := i.ItemContractUseCase.PurchaseItem(itemId, accountSessionToken.AccountId); err != nil {
			ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: err.Error()})
		} else {
			ctx.Status(http.StatusOK)
		}
	})
}

func (i *ItemContractController) PaymentOfItemPriceRoute(paramKey string) func(*gin.Context) {
	return i.Session.RequiredSession(func(ctx *gin.Context, token *string) {
		itemId := ctx.Param(paramKey)
		accountSessionToken := model.DecodeToAccountSessionToken(*token)
		if err := i.ItemContractUseCase.PaymentOfItemPrice(itemId, accountSessionToken.AccountId); err != nil {
			ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: err.Error()})
		} else {
			ctx.Status(http.StatusOK)
		}
	})
}

func (i *ItemContractController) ReceiptConfirmationRoute(paramKey string) func(*gin.Context) {
	return i.Session.RequiredSession(func(ctx *gin.Context, token *string) {
		itemId := ctx.Param(paramKey)
		accountSessionToken := model.DecodeToAccountSessionToken(*token)
		if err := i.ItemContractUseCase.ReceiptConfirmation(itemId, accountSessionToken.AccountId); err != nil {
			ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: err.Error()})
		} else {
			ctx.Status(http.StatusOK)
		}
	})
}
