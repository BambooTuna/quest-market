package goods

import (
	error2 "github.com/BambooTuna/quest-market/backend/error"
	"github.com/BambooTuna/quest-market/backend/settings"
	"github.com/go-playground/validator/v10"
)

type ProductDetails struct {
	ProductId   string `db:"product_id"`
	Title       string `validate:"required" db:"title"`
	Detail      string `validate:"required" db:"detail"`
	Price       int64  `validate:"min=1" db:"price"`
	PresenterId string `db:"presenter_id" json:"presenterId"`
	State       string `validate:"required" db:"state"`
}

func (p *ProductDetails) Validate() (*ProductDetails, error) {
	validate := validator.New()
	var errorMessages []error2.CustomError
	if err := validate.Struct(p); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errorMessages = append(errorMessages, error2.ValidateError(err.Field(), err.Tag()))
		}
		return nil, error2.Errors(errorMessages)
	}
	return p, nil
}

func GenerateProductDetails(title, detail, presenterId, state string, price int64) (*ProductDetails, error) {
	uuid, err := settings.GenerateUUID()
	if err != nil {
		return nil, err
	}
	productDetails := &ProductDetails{
		ProductId:   uuid,
		Title:       title,
		Detail:      detail,
		Price:       price,
		PresenterId: presenterId,
		State:       state,
	}
	return productDetails.Validate()
}
