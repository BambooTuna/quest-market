package json

import (
	"github.com/BambooTuna/quest-market/backend/command"
	"github.com/BambooTuna/quest-market/backend/model/item"
)

type EditItemDetailsRequestJson struct {
	Title  string     `json:"title"`
	Detail string     `json:"detail"`
	Price  int64      `json:"price"`
	State  item.State `json:"state"`
}

func (e EditItemDetailsRequestJson) GenerateEditItemDetailsCommand(itemId, practitionerAccountId string) (command.EditItemDetailsCommand, error) {
	if e.State == item.Open || e.State == item.Draft || e.State == item.Deleted {
		return command.EditItemDetailsCommand{
			Title:                 e.Title,
			Detail:                e.Detail,
			Price:                 e.Price,
			State:                 e.State,
			ItemId:                itemId,
			PractitionerAccountId: practitionerAccountId,
		}, nil
	} else {
		return command.EditItemDetailsCommand{}, nil
	}
}
