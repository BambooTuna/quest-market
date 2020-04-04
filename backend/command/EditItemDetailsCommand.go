package command

import (
	"github.com/BambooTuna/quest-market/backend/model/item"
)

type EditItemDetailsCommand struct {
	Title                 string
	Detail                string
	Price                 int64
	State                 item.State
	ItemId                string
	PractitionerAccountId string
}
