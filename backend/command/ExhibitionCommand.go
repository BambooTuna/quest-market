package command

import (
	"github.com/BambooTuna/quest-market/backend/model/item"
)

type ExhibitionCommand struct {
	Title           string
	Detail          string
	Price           int64
	SellerAccountId string
}

func (e ExhibitionCommand) ToContractDetails() (*item.ContractDetails, error) {
	return item.Generate(e.Title, e.Detail, e.Price, e.SellerAccountId)
}
