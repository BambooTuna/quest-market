package item

import (
	error2 "github.com/BambooTuna/quest-market/backend/error"
	"github.com/BambooTuna/quest-market/backend/settings"
	"github.com/go-playground/validator/v10"
	"time"
)

type ContractDetails struct {
	ItemId             string    `db:"item_id"`
	Title              string    `validate:"required" db:"title"`
	Detail             string    `validate:"required" db:"detail"`
	Price              int64     `validate:"min=1" db:"price"`
	PurchaserAccountId string    `db:"purchaser_account_id"`
	SellerAccountId    string    `db:"seller_account_id"`
	State              State     `db:"state"`
	CreatedAt          time.Time `db:"created_at"`
	UpdatedAt          time.Time `db:"updated_at"`
}

func Generate(title, detail string, price int64, sellerAccountId string) (*ContractDetails, error) {
	uuid, err := settings.GenerateUUID()
	if err != nil {
		return nil, err
	}
	details := ContractDetails{
		ItemId:          uuid,
		Title:           title,
		Detail:          detail,
		Price:           price,
		SellerAccountId: sellerAccountId,
		State:           Open,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	if err := details.Validate(); err != nil {
		return nil, err
	}
	return &details, nil
}

func (d *ContractDetails) Validate() error {
	validate := validator.New()
	var errorMessages []error2.CustomError
	if err := validate.Struct(d); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errorMessages = append(errorMessages, error2.ValidateError(err.Field(), err.Tag()))
		}
		return error2.Errors(errorMessages)
	}
	return nil
}

func (d *ContractDetails) ChangeState(state State) *ContractDetails {
	d.State = state
	return d
}

func (d *ContractDetails) PurchaseBy(purchaserAccountId string) *ContractDetails {
	d.PurchaserAccountId = purchaserAccountId
	d.State = Unpaid
	return d
}
