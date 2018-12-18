package payloads

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

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

func NewChargePayload(r *http.Request) (*Charge, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	payload := Charge{}
	if r.Header.Get("Content-Type") == "application/json" {
		err := json.Unmarshal(body, &payload)
		if err != nil {
			return nil, err
		}

		return &payload, nil
	}

	values, err := url.ParseQuery(string(body))
	if err != nil {
		return nil, err
	}

	payload.Currency = values.Get("currency")
	payload.Description = values.Get("description")

	id := values.Get("customer_id")
	if id != "" {
		payload.CustomerID = &id
	}

	amount := values.Get("amount")
	payload.Amount, err = strconv.ParseInt(amount, 10, 64)
	if err != nil {
		return nil, err
	}

	return &payload, nil
}
