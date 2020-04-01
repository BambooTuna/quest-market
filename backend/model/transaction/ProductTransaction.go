package transaction

import "time"

type ProductTransaction struct {
	TransactionId      int64                  `db:"transaction_id"`
	TransactionType    ProductTransactionType `db:"transaction_type"`
	ProductId          string                 `db:"product_id"`
	PurchaserAccountId string                 `db:"purchaser_account_id"`
	SellerAccountId    string                 `db:"seller_account_id"`
	CreatedAt          time.Time              `db:"created_at"`
}

func ApplyProductTransaction(transactionType ProductTransactionType, productId, purchaserAccountId, sellerAccountId string) *ProductTransaction {
	transaction := ProductTransaction{
		TransactionType:    transactionType,
		ProductId:          productId,
		PurchaserAccountId: purchaserAccountId,
		SellerAccountId:    sellerAccountId,
		CreatedAt:          time.Now(),
	}
	return &transaction
}

type ProductTransactionType string

const (
	WaitingForPayment ProductTransactionType = "waiting_for_payment"
	WaitingToReceive  ProductTransactionType = "waiting_to_receive"
	Complete          ProductTransactionType = "complete"
)

func (after ProductTransactionType) CanOverwrite(before *ProductTransactionType) bool {
	switch after {
	case WaitingForPayment:
		return before == nil
	case WaitingToReceive:
		if before != nil {
			return *before == WaitingForPayment
		}
	case Complete:
		if before != nil {
			return *before == WaitingToReceive
		}
	}
	return false
}
