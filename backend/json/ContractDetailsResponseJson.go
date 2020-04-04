package json

import (
	"github.com/BambooTuna/quest-market/backend/model/item"
)

type ContractDetailsResponseJson struct {
	ItemId          string     `json:"item_id"`
	Title           string     `json:"title"`
	Detail          string     `json:"detail"`
	Price           int64      `json:"price"`
	SellerAccountId string     `json:"seller_account_id"`
	State           item.State `json:"state"`
}

func ConvertToContractDetailsResponseJson(p *item.ContractDetails, acquiredBy string) *ContractDetailsResponseJson {
	return &ContractDetailsResponseJson{
		ItemId:          p.ItemId,
		Title:           p.Title,
		Detail:          p.Detail,
		Price:           p.Price,
		SellerAccountId: p.SellerAccountId,
		State:           p.State.Secret(!(p.SellerAccountId == acquiredBy || p.PurchaserAccountId == acquiredBy)),
	}
}

func ConvertToContractDetailsListResponseJson(p []*item.ContractDetails, acquiredBy string) []*ContractDetailsResponseJson {
	r := make([]*ContractDetailsResponseJson, len(p))
	for i, e := range p {
		r[i] = ConvertToContractDetailsResponseJson(e, acquiredBy)
	}
	return r
}
