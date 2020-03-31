package transaction

import "time"

type MoneyTransaction struct {
	TransactionId   int64                `db:"transaction_id"`
	TransactionType MoneyTransactionType `db:"transaction_type"`
	AccountId       string               `db:"account_id"`
	Currency        Currency             `db:"currency"`
	RealPart        int64                `db:"real_part"`
	ExponentPart    int64                `db:"exponent_part"`
	CreatedAt       time.Time            `db:"created_at"`
}

type MoneyTransactionType string

const (
	Deposit  MoneyTransactionType = "deposit"
	Withdraw MoneyTransactionType = "withdraw"
	//CanceledDeposit MoneyTransactionType = "canceled_deposit"
	//CanceledWithdraw MoneyTransactionType = "canceled_withdraw"
)

func (m MoneyTransactionType) ToNumber() int64 {
	switch m {
	case Deposit:
		return 1
	case Withdraw:
		return -1
		//case CanceledDeposit:
		//	return -1
		//case CanceledWithdraw:
		//	return 1
	}
	return 0
}

type Currency string

const (
	JPY Currency = "jpy"
)
