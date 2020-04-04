package item

type State string

type Accessor string

const (
	General Accessor = "general"
	Seller  Accessor = "seller"
	Buyer   Accessor = "buyer"
)

const (
	//Public
	Open    State = "open"
	Draft   State = "draft"
	Deleted State = "deleted"

	//Seller, Buyerのみ
	Unpaid   State = "unpaid"
	Sent     State = "sent"
	Complete State = "complete"
	//Public向け
	Sold State = "sold"
)

func (s State) Secret(hide bool) State {
	if hide {
		switch s {
		case Open:
			return Open
		case Draft:
			return Draft
		case Deleted:
			return Deleted
		default:
			return Sold
		}
	} else {
		return s
	}
}
