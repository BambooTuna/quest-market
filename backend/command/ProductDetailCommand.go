package command

import "github.com/BambooTuna/quest-market/backend/model/goods"

type ProductDetailCommand struct {
	Title  string
	Detail string
	Price  int64
	State  goods.State
}
