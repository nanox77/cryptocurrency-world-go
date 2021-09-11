package input

type SellRequest struct {
	Coin   string `validate:"required" json:"coin"`
	Amount int    `validate:"required" json:"amount"`
}
