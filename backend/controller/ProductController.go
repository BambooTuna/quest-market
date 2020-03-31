package controller

import (
	"github.com/BambooTuna/quest-market/backend/command"
	"github.com/BambooTuna/quest-market/backend/json"
	"github.com/BambooTuna/quest-market/backend/lib/session"
	"github.com/BambooTuna/quest-market/backend/model"
	"github.com/BambooTuna/quest-market/backend/model/goods"
	"github.com/BambooTuna/quest-market/backend/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductController struct {
	Session               session.Session
	ProductDetailsUseCase usecase.ProductDetailsUseCase
}

func (c *ProductController) GetProductDetailsRoute() func(*gin.Context) {
	return c.Session.OptionalRequiredSession(func(ctx *gin.Context, token *string) {
		productId := ctx.Param("productId")

		var productDetails *goods.ProductDetails
		var err error
		if token != nil {
			accountSessionToken := model.DecodeToAccountSessionToken(*token)
			productDetails, err = c.ProductDetailsUseCase.GetPrivateProductDetails(ctx, productId, accountSessionToken.AccountId)
		} else {
			productDetails, err = c.ProductDetailsUseCase.GetPublicProductDetails(ctx, productId)
		}

		if err != nil {
			ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, productDetails)
	})
}

func (c *ProductController) GetOpenProductsRoute() func(*gin.Context) {
	return func(ctx *gin.Context) {
		productDetails, err := c.ProductDetailsUseCase.GetOpenProducts(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, productDetails)
	}
}

func (c *ProductController) GetMyProductListRoute() func(*gin.Context) {
	return c.Session.RequiredSession(func(ctx *gin.Context, token *string) {
		accountSessionToken := model.DecodeToAccountSessionToken(*token)
		productDetails, err := c.ProductDetailsUseCase.GetMyExistingProducts(ctx, accountSessionToken.AccountId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, productDetails)
	})
}

func (c *ProductController) ExhibitionRoute() func(*gin.Context) {
	return c.Session.RequiredSession(func(ctx *gin.Context, token *string) {
		var exhibitionRequestJson json.ExhibitionRequestJson

		accountSessionToken := model.DecodeToAccountSessionToken(*token)
		if err := ctx.BindJSON(&exhibitionRequestJson); err != nil {
			ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: err.Error()})
			return
		}
		command := exhibitionRequestJson.GenerateExhibitionCommand(accountSessionToken.AccountId)
		productDetails, err := c.ProductDetailsUseCase.Exhibition(ctx, &command)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, productDetails.PresenterId)
	})
}

func (c *ProductController) UpdateProductDetailsRoute() func(*gin.Context) {
	return c.Session.RequiredSession(func(ctx *gin.Context, token *string) {
		productId := ctx.Param("productId")

		accountSessionToken := model.DecodeToAccountSessionToken(*token)
		var exhibitionRequestJson json.ExhibitionRequestJson
		if err := ctx.BindJSON(&exhibitionRequestJson); err != nil {
			ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: err.Error()})
			return
		}
		command := command.UpdateProductCommand{ProductId: productId, PractitionerId: accountSessionToken.AccountId, ProductDetail: exhibitionRequestJson.GenerateProductDetailCommand()}
		productDetails, err := c.ProductDetailsUseCase.UpdateProductDetails(ctx, &command)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, productDetails)
	})
}
