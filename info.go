package dolyame

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"net/http"
)

type InfoResponse struct {
	Status          string            `json:"status"`
	Amount          decimal.Decimal   `json:"amount"`
	ResidualAmount  decimal.Decimal   `json:"residual_amount"`
	Link            string            `json:"link"`
	PaymentSchedule []PaymentSchedule `json:"payment_schedule"`
	RefundInfo      RefundInfo        `json:"refund_info"`
}

func (c *Client) Info(orderID string, correlationID string) (*InfoResponse, error) {
	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("https://%v/v1/orders/%v/info", c.host, orderID),
		nil,
	)

	c.setAuth(req)

	req.Header.Add("X-Correlation-ID", correlationID)

	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("dolyame bad response %v: %v", response.Status, string(body))
	}

	var resp InfoResponse

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
