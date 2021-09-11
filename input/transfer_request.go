package input

type TransferRequest struct {
	ToUser string `validate:"required" json:"to_user"`
	Coin   string `validate:"required" json:"coin"`
	Amount int    `validate:"required" json:"amount"`
}