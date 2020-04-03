package json

import (
	"github.com/BambooTuna/quest-market/backend/command"
)

type ExhibitionRequestJson struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Price  int64  `json:"price"`
}

func (e ExhibitionRequestJson) GenerateExhibitionCommand(sellerAccountId string) command.ExhibitionCommand {
	return command.ExhibitionCommand{Title: e.Title, Detail: e.Detail, Price: e.Price, SellerAccountId: sellerAccountId}
}
