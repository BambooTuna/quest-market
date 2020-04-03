package dao

import (
	"github.com/BambooTuna/quest-market/backend/model/item"
	"github.com/BambooTuna/quest-market/backend/settings"
	"time"
)

type ItemContractDao interface {

	//state=Deleted, Draft以外
	Publishable(q settings.QuantityLimit) []*item.ContractDetails

	//state=Deleted, Draft以外
	ResolvePublishableItemByItemId(itemId string) *item.ContractDetails

	//state=Deleted, Draft以外 (自分のみDraftを取得可能)
	ResolvePrivateItemByItemId(itemId, practitioner string) *item.ContractDetails

	ResolveByAccountId(q settings.QuantityLimit, accountId string) []*item.ContractDetails

	//state=Open
	ResolveOpenItemByItemId(itemId string) *item.ContractDetails
	//state=Unpaid && 購入者のみ取得可能
	ResolveUnpaidItemByItemId(itemId, purchaserAccountId string) *item.ContractDetails
	//state=Sent && 購入者のみ取得可能
	ResolveSentItemByItemId(itemId, purchaserAccountId string) *item.ContractDetails

	Insert(record *item.ContractDetails) error

	UpdateItemDetails(record *item.ContractDetails) error
	UpdateContractDetails(record *item.ContractDetails) error
}

type ItemContractDetails struct {
	ItemDetails
	ContractDetails
}

type ItemDetails struct {
	ItemId string `db:"item_id"`
	Title  string `db:"title"`
	Detail string `db:"detail"`
	Price  int64  `db:"price"`
}

func (ItemDetails) FromContractDetails(contractDetails *item.ContractDetails) *ItemDetails {
	return &ItemDetails{
		ItemId: contractDetails.ItemId,
		Title:  contractDetails.Title,
		Detail: contractDetails.Detail,
		Price:  contractDetails.Price,
	}
}

type ContractDetails struct {
	ItemId             string     `db:"item_id"`
	PurchaserAccountId string     `db:"purchaser_account_id"`
	SellerAccountId    string     `db:"seller_account_id"`
	State              item.State `db:"state"`
	CreatedAt          time.Time  `db:"created_at"`
	UpdatedAt          time.Time  `db:"updated_at"`
}

func (ContractDetails) FromContractDetails(contractDetails *item.ContractDetails) *ContractDetails {
	return &ContractDetails{
		ItemId:             contractDetails.ItemId,
		PurchaserAccountId: contractDetails.PurchaserAccountId,
		SellerAccountId:    contractDetails.SellerAccountId,
		State:              contractDetails.State,
		CreatedAt:          contractDetails.CreatedAt,
		UpdatedAt:          contractDetails.UpdatedAt,
	}
}
