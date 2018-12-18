package payloads

type Charge struct {
	Amount      int64   `json:"amount"`
	Currency    string  `json:"currency"`
	Description string  `json:"description"`
	CustomerID  *string `json:"customer_id"`
}

func (c Charge) Valid() bool {
	return c.Currency == "btc" &&
		c.Amount != 0 &&
		len(c.Description) <= 256 &&
		(c.CustomerID == nil || len(*c.CustomerID) <= 64)
}
