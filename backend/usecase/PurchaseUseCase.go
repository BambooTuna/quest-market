package usecase

import (
	"context"
	"github.com/BambooTuna/quest-market/backend/aggregate"
	"github.com/BambooTuna/quest-market/backend/command"
	"github.com/BambooTuna/quest-market/backend/dao"
	error2 "github.com/BambooTuna/quest-market/backend/error"
	"github.com/BambooTuna/quest-market/backend/model/goods"
	"github.com/BambooTuna/quest-market/backend/model/transaction"
)

type PurchaseUseCase struct {
	ProductDetailsDao            dao.ProductDetailsDao
	MoneyManagementUseCase       *MoneyManagementUseCase
	ProductTransactionAggregates *aggregate.ProductTransactionAggregates
}

func (p *PurchaseUseCase) Application(ctx context.Context, c *command.PurchaseApplicationCommand) error {
	productDetails, err := p.ProductDetailsDao.ResolveByProductId(ctx, c.ProductId, "")
	if err != nil {
		return error2.Error(error2.ProductSold)
	}
	if productDetails.PresenterId == c.PurchaserAccountId {
		return error2.Error(error2.PurchaseYourself)
	}
	productDetails.State = goods.Closed
	if count, err := p.ProductDetailsDao.Update(ctx, productDetails); count == 0 {
		return error2.Error(error2.NoUpdatesWereFound)
	} else if err != nil {
		return error2.Error(error2.SqlRequestFailed)
	}

	productTransaction := transaction.ApplyProductTransaction(transaction.WaitingForPayment, c.ProductId, c.PurchaserAccountId, productDetails.PresenterId)

	if err := p.ProductTransactionAggregates.SendTransaction(productTransaction); err != nil {
		return error2.Error("Application: 深刻なエラー")
	}
	return nil
}

func (p *PurchaseUseCase) Payment(ctx context.Context, c *command.PurchasePaymentCommand) error {
	productDetails, err := p.ProductDetailsDao.ResolveByPurchasedProductId(ctx, c.ProductId)
	if err != nil {
		return err
	}
	if err := p.MoneyManagementUseCase.ManagementKeeps(c.PurchaserAccountId, productDetails.Price); err != nil {
		return error2.Error(error2.LackOfMoney)
	}
	productTransaction := transaction.ApplyProductTransaction(transaction.WaitingToReceive, c.ProductId, c.PurchaserAccountId, productDetails.PresenterId)
	return p.ProductTransactionAggregates.SendTransaction(productTransaction)
}

func (p *PurchaseUseCase) ReceiptConfirmation(ctx context.Context, c *command.PurchaseReceiptConfirmationCommand) error {
	productDetails, err := p.ProductDetailsDao.ResolveByPurchasedProductId(ctx, c.ProductId)
	if err != nil {
		return err
	}
	if err := p.MoneyManagementUseCase.ManagementPayment(c.SellerAccountId, productDetails.Price); err != nil {
		return error2.Error("ReceiptConfirmation: 致命的なエラー")
	}
	productTransaction := transaction.ApplyProductTransaction(transaction.Complete, c.ProductId, "", c.SellerAccountId)
	return p.ProductTransactionAggregates.SendTransaction(productTransaction)
}

func (p *PurchaseUseCase) GetTransactionByAccountId(accountId string) []*transaction.ProductTransaction {
	return p.ProductTransactionAggregates.GetTransactionByAccountId(accountId)
}
