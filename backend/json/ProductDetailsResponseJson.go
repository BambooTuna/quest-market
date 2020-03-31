package json

import "github.com/BambooTuna/quest-market/backend/model/goods"

type ProductDetailsResponseJson struct {
	ProductId   string `json:"id"`
	Title       string `json:"productTitle"`
	Detail      string `json:"productDetail"`
	Price       int64  `json:"requestPrice"`
	PresenterId string `json:"presenterId"`
	State       string `json:"state"`
}

func ConvertToProductDetailsResponseJson(p goods.ProductDetails) ProductDetailsResponseJson {
	return ProductDetailsResponseJson{
		ProductId:   p.ProductId,
		Title:       p.Title,
		Detail:      p.Detail,
		Price:       p.Price,
		PresenterId: p.PresenterId,
		State:       p.State,
	}
}

func ConvertToProductDetailsListResponseJson(p []goods.ProductDetails) []ProductDetailsResponseJson {
	r := make([]ProductDetailsResponseJson, len(p))
	for i, e := range p {
		r[i] = ConvertToProductDetailsResponseJson(e)
	}
	return r
}
