package command

type PurchaseApplicationCommand struct {
	ProductId          string
	PurchaserAccountId string
}

type PurchasePaymentCommand struct {
	ProductId          string
	PurchaserAccountId string
}

type PurchaseReceiptConfirmationCommand struct {
	ProductId       string
	SellerAccountId string
}
