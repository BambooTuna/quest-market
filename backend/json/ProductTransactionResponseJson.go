package json

import (
	"github.com/BambooTuna/quest-market/backend/model/transaction"
	"time"
)

type ProductTransactionResponseJson struct {
	TransactionId      int64                              `json:"transaction_id"`
	TransactionType    transaction.ProductTransactionType `json:"transaction_type"`
	ProductId          string                             `json:"product_id"`
	PurchaserAccountId string                             `json:"purchaser_account_id"`
	SellerAccountId    string                             `json:"seller_account_id"`
	CreatedAt          time.Time                          `json:"created_at"`
}

func ConvertToProductTransactionResponseJson(p *transaction.ProductTransaction) *ProductTransactionResponseJson {
	return &ProductTransactionResponseJson{
		TransactionId:      p.TransactionId,
		TransactionType:    p.TransactionType,
		ProductId:          p.ProductId,
		PurchaserAccountId: p.PurchaserAccountId,
		SellerAccountId:    p.SellerAccountId,
		CreatedAt:          p.CreatedAt,
	}
}

func ConvertToProductTransactionListResponseJson(p []*transaction.ProductTransaction) []*ProductTransactionResponseJson {
	r := make([]*ProductTransactionResponseJson, len(p))
	for i, e := range p {
		r[i] = ConvertToProductTransactionResponseJson(e)
	}
	return r
}
