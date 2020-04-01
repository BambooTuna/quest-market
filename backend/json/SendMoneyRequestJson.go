package json

type SendMoneyRequestJson struct {
	ReceiverAccountId string `json:"to"`
	Money             int64  `json:"amount"`
}
