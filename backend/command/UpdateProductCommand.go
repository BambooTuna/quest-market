package command

type UpdateProductCommand struct {
	ProductId      string
	PractitionerId string
	ProductDetail  ProductDetailCommand
}
