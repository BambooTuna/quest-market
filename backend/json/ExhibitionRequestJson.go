package json

import (
	"github.com/BambooTuna/quest-market/backend/command"
	"strconv"
)

type ExhibitionRequestJson struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Price  string `json:"price"`
	State  string `json:"state"`
}

func (e ExhibitionRequestJson) GenerateProductDetailCommand() command.ProductDetailCommand {
	price, err := strconv.ParseInt(e.Price, 10, 64)
	if err != nil {
		price = 0
	}
	return command.ProductDetailCommand{Title: e.Title, Detail: e.Detail, Price: price, State: e.State}
}

func (e ExhibitionRequestJson) GenerateExhibitionCommand(presenterId string) command.ExhibitionCommand {
	return command.ExhibitionCommand{ProductDetailCommand: e.GenerateProductDetailCommand(), PresenterId: presenterId}
}
