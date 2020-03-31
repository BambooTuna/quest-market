package goods

import (
	"github.com/BambooTuna/quest-market/backend/settings"
	"github.com/go-playground/validator/v10"
)

type ProductDetails struct {
	ProductId   string `db:"product_id" json:"id"`
	Title       string `db:"title" json:"productTitle"`
	Detail      string `db:"detail" json:"productDetail"`
	Price       int64  `db:"price" json:"requestPrice"`
	PresenterId string `db:"presenter_id" json:"presenterId"`
	State       string `db:"state" json:"state"`
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
	validate := validator.New()
	if err := validate.Struct(productDetails); err != nil {
		return nil, err
	}
	return productDetails, nil
}
