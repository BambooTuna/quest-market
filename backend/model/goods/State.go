package goods

type State string

const (
	Open    State = "open"
	Draft   State = "draft"
	Closed  State = "closed"
	Deleted State = "deleted"
)
