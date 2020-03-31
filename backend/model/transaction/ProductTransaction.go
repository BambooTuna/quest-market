package transaction

import "time"

type ProductTransaction struct {
	TransactionId     int64                  `db:"transaction_id"`
	TransactionType   ProductTransactionType `db:"transaction_type"`
	ProductId         string                 `db:"product_id"`
	SenderAccountId   string                 `db:"sender_account_id"`
	ReceiverAccountId string                 `db:"receiver_account_id"`
	CreatedAt         time.Time              `db:"created_at"`
}

func (l *ProductTransaction) Override(r *ProductTransaction) *ProductTransaction {
	if l.TransactionType.IsAfterThan(r.TransactionType) {
		return l
	} else {
		return r
	}
}

type ProductTransactionType string

const (
	WaitingForPayment ProductTransactionType = "waiting_for_payment"
	WaitingToReceive  ProductTransactionType = "waiting_to_receive"
	Complete          ProductTransactionType = "complete"
)

func (l ProductTransactionType) IsAfterThan(r ProductTransactionType) bool {
	switch l {
	case WaitingToReceive:
		if r == WaitingForPayment {
			return true
		}
	case Complete:
		if r != Complete {
			return true
		}
	}
	return false
}
