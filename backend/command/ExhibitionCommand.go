package command

import "github.com/BambooTuna/quest-market/backend/model/goods"

type ExhibitionCommand struct {
	ProductDetailCommand
	PresenterId string
}

func (e ExhibitionCommand) ToProductDetails() (*goods.ProductDetails, error) {
	return goods.GenerateProductDetails(e.Title, e.Detail, e.PresenterId, e.State, e.Price)
}
