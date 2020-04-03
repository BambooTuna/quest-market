package command

import (
	"github.com/BambooTuna/quest-market/backend/model/item"
)

type ItemDetailsCommand struct {
	Title  string
	Detail string
	Price  int64
	State  item.State
}
