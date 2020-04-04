package item

type State string

const (
	Open State = "open"
	Sold State = "sold"

	Draft State = "draft"

	Unpaid   State = "unpaid"
	Sent     State = "sent"
	Complete State = "complete"

	Deleted State = "deleted"
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
