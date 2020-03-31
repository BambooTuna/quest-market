package json

import (
	"github.com/BambooTuna/quest-market/backend/command"
)

type ExhibitionRequestJson struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Price  int64  `json:"price"`
	State  string `json:"state"`
}

func (e ExhibitionRequestJson) GenerateProductDetailCommand() command.ProductDetailCommand {
	return command.ProductDetailCommand{Title: e.Title, Detail: e.Detail, Price: e.Price, State: e.State}
}

func (e ExhibitionRequestJson) GenerateExhibitionCommand(presenterId string) command.ExhibitionCommand {
	return command.ExhibitionCommand{ProductDetailCommand: e.GenerateProductDetailCommand(), PresenterId: presenterId}
}
