package resources

type Charge struct {
	ID             string `json:"id"`
	Object         string `json:"object"`
	Amount         int64  `json:"amount"`
	AmountSatoshi  string `json:"amount_satoshi"`
	Currency       string `json:"currency"`
	Description    string `json:"description"`
	Paid           bool   `json:"paid"`
	PaymentHash    string `json:"payment_hash"`
	PaymentRequest string `json:"payment_request"`
	Created        int64  `json:"created"`
	Updated        int64  `json:"updated"`
}
