package input

type DepositRequest struct {
	Amount int `validate:"required" json:"amount"`
}
