package input

type WithdrawRequest struct {
	Amount int `validate:"required" json:"amount"`
}
