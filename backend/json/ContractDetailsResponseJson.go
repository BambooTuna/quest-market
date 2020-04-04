package json

import (
	"github.com/BambooTuna/quest-market/backend/model/item"
	"time"
)

type ContractDetailsResponseJson struct {
	ItemId          string        `json:"item_id"`
	Title           string        `json:"title"`
	Detail          string        `json:"detail"`
	Price           int64         `json:"price"`
	SellerAccountId string        `json:"seller_account_id"`
	State           item.State    `json:"state"`
	UpdatedAt       time.Time     `json:"updated_at"`
	Accessor        item.Accessor `json:"accessor"`
}

type Side string

const (
	Buy  Side = "buy"
	Sell Side = "sell"
)

func ConvertToContractDetailsResponseJson(p *item.ContractDetails, acquiredBy string) *ContractDetailsResponseJson {
	var accessor item.Accessor
	if p.SellerAccountId == acquiredBy {
		accessor = item.Seller
	} else if p.PurchaserAccountId == acquiredBy {
		accessor = item.Buyer
	} else {
		accessor = item.General
	}

	return &ContractDetailsResponseJson{
		ItemId:          p.ItemId,
		Title:           p.Title,
		Detail:          p.Detail,
		Price:           p.Price,
		SellerAccountId: p.SellerAccountId,
		State:           p.State.Secret(!(accessor == item.Seller || accessor == item.Buyer)),
		UpdatedAt:       p.UpdatedAt,
		Accessor:        accessor,
	}
}

func ConvertToContractDetailsListResponseJson(p []*item.ContractDetails, acquiredBy string) []*ContractDetailsResponseJson {
	r := make([]*ContractDetailsResponseJson, len(p))
	for i, e := range p {
		r[i] = ConvertToContractDetailsResponseJson(e, acquiredBy)
	}
	return r
}
