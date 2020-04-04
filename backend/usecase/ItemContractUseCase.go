package usecase

import (
	"github.com/BambooTuna/quest-market/backend/command"
	"github.com/BambooTuna/quest-market/backend/dao"
	error2 "github.com/BambooTuna/quest-market/backend/error"
	"github.com/BambooTuna/quest-market/backend/model/item"
	"github.com/BambooTuna/quest-market/backend/settings"
)

type ItemContractUseCase struct {
	ItemContractDao        dao.ItemContractDao
	MoneyManagementUseCase *MoneyManagementUseCase
}

func (i *ItemContractUseCase) GetPublicItemContract(itemId, practitioner string) (*item.ContractDetails, error) {
	var contractDetails *item.ContractDetails
	if practitioner == "" {
		contractDetails = i.ItemContractDao.ResolvePublishableItemByItemId(itemId)
	} else {
		contractDetails = i.ItemContractDao.ResolvePrivateItemByItemId(itemId, practitioner)
	}
	if contractDetails == nil {
		return nil, error2.Error(error2.ItemNotFoundError)
	} else {
		return contractDetails, nil
	}
}

func (i *ItemContractUseCase) GetPublicItemContracts(q settings.QuantityLimit) []*item.ContractDetails {
	return i.ItemContractDao.Publishable(q)
}

func (i *ItemContractUseCase) GetMyItemContracts(q settings.QuantityLimit, accountId string) []*item.ContractDetails {
	return i.ItemContractDao.ResolveByAccountId(q, accountId)
}

func (i *ItemContractUseCase) CreateNewItemContract(c *command.ExhibitionCommand) (*item.ContractDetails, error) {
	if contractDetails, err := c.ToContractDetails(); err != nil {
		return nil, err
	} else if err := i.ItemContractDao.Insert(contractDetails); err != nil {
		return nil, error2.Error(error2.CustomError(err.Error()))
	} else {
		return contractDetails, nil
	}
}

func (i *ItemContractUseCase) EditItemContract(c *command.EditItemDetailsCommand) (*item.ContractDetails, error) {
	if contractDetails := i.ItemContractDao.ResolvePrivateItemByItemId(c.ItemId, c.PractitionerAccountId); contractDetails == nil {
		return nil, error2.Error(error2.ItemNotFoundError)
	} else if newContractDetails, err := contractDetails.Update(c.Title, c.Detail, c.Price, c.State); err != nil {
		return nil, error2.Error(error2.CustomError(err.Error()))
	} else if err := i.ItemContractDao.Update(newContractDetails); err != nil {
		return contractDetails, error2.Error(error2.CustomError(err.Error()))
	} else {
		return contractDetails, nil
	}
}

func (i *ItemContractUseCase) PurchaseItem(itemId, purchaserAccountId string) error {
	if contractDetails := i.ItemContractDao.ResolveOpenItemByItemId(itemId); contractDetails == nil {
		return error2.Error(error2.ItemSoldError)
	} else if contractDetails.SellerAccountId == purchaserAccountId {
		return error2.Error(error2.PurchaseYourself)
	} else if err := i.ItemContractDao.UpdateContractDetails(contractDetails.PurchaseBy(purchaserAccountId)); err != nil {
		return error2.Error(error2.CannotBuy)
	} else {
		return nil
	}
}

func (i *ItemContractUseCase) PaymentOfItemPrice(itemId, purchaserAccountId string) error {
	contractDetails := i.ItemContractDao.ResolveUnpaidItemByItemId(itemId, purchaserAccountId)
	if contractDetails == nil {
		return error2.Error(error2.ItemNotFoundError)
	} else if err := i.MoneyManagementUseCase.ManagementKeeps(contractDetails.PurchaserAccountId, contractDetails.Price); err != nil {
		return error2.Error(error2.LackOfMoney)
	} else if err := i.ItemContractDao.UpdateContractDetails(contractDetails.ChangeState(item.Sent)); err != nil {
		return error2.Error("深刻なエラー: 1001")
	} else {
		return nil
	}
}

func (i *ItemContractUseCase) ReceiptConfirmation(itemId, purchaserAccountId string) error {
	contractDetails := i.ItemContractDao.ResolveSentItemByItemId(itemId, purchaserAccountId)
	if contractDetails == nil {
		return error2.Error(error2.ItemNotFoundError)
	} else if err := i.MoneyManagementUseCase.ManagementPayment(contractDetails.SellerAccountId, contractDetails.Price); err != nil {
		return error2.Error("深刻なエラー: 5001")
	} else if err := i.ItemContractDao.UpdateContractDetails(contractDetails.ChangeState(item.Complete)); err != nil {
		return error2.Error("深刻なエラー: 1002")
	} else {
		return nil
	}
}
