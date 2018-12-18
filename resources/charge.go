package resources

import (
	"time"

	"github.com/edouardparis/spark/payloads"
)

type Charge struct {
	ID             string `json:"id"`
	Object         string `json:"object"`
	Amount         int64  `json:"amount"`
	AmountSatoshi  int64  `json:"amount_satoshi"`
	Currency       string `json:"currency"`
	Description    string `json:"description"`
	Paid           bool   `json:"paid"`
	PaymentHash    string `json:"payment_hash"`
	PaymentRequest string `json:"payment_request"`
	Created        int64  `json:"created"`
	Updated        int64  `json:"updated"`
}

func NewCharge(payload *payloads.Charge) *Charge {
	t := time.Now().Unix()
	return &Charge{
		Object:         "charge",
		Amount:         payload.Amount,
		AmountSatoshi:  payload.Amount,
		Currency:       payload.Currency,
		Description:    payload.Description,
		PaymentHash:    "3dd10b44194224936b57349915daca4ec05b5603b0342adf45ec639e4e11189d",
		PaymentRequest: "lntb420u1pdsdxfepp58hgsk3qeggjfx66hxjv3tkk2fmq9k4srkq6z4h69a33euns3rzwsdq4xysyymr0vd4kzcmrd9hx7cqp2fs6hglhgfax7depekep53kmlgkswhcvxlmju3e4k0cdex6ml4xpygcxzt93julus09hj30fruzw9l65n5uktqe9khmlk8uh8pvl3f7sp3026hz",
		Created:        t,
		Updated:        t,
	}
}
