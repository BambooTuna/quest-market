package json

type MyItemContractsResponseJson struct {
	ContractType                string                       `json:"id"`
	ContractDetailsResponseJson *ContractDetailsResponseJson `json:"id"`
}
