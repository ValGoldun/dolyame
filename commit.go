package dolyame

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"net/http"
)

type CommitRequest struct {
	Amount        decimal.Decimal `json:"amount"`
	PrepaidAmount decimal.Decimal `json:"prepaid_amount"`
	Items         []*Item         `json:"items"`
}

func NewCommitRequest() *CommitRequest {
	return new(CommitRequest)
}

func (cr *CommitRequest) SetAmount(amount decimal.Decimal) *CommitRequest {
	cr.Amount = amount
	return cr
}

func (cr *CommitRequest) SetPrepaidAmount(amount decimal.Decimal) *CommitRequest {
	cr.PrepaidAmount = amount
	return cr
}

func (cr *CommitRequest) AddItem(item *Item) *CommitRequest {
	cr.Items = append(cr.Items, item)
	return cr
}

type CommitResponse struct {
	Status          string            `json:"status"`
	Amount          decimal.Decimal   `json:"amount"`
	ResidualAmount  decimal.Decimal   `json:"residual_amount"`
	Link            string            `json:"link"`
	PaymentSchedule []PaymentSchedule `json:"payment_schedule"`
	RefundInfo      RefundInfo        `json:"refund_info"`
}

func (c *Client) Commit(request *CommitRequest, orderID string, correlationID string) (*CommitResponse, error) {
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("https://%v/v1/orders/%v/commit", c.host, orderID),
		bytes.NewReader(jsonBody),
	)

	c.setAuth(req)

	req.Header.Add("X-Correlation-ID", correlationID)
	req.Header.Add("Content-Type", "application/json")

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

	var resp CommitResponse

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
