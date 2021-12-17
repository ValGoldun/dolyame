package dolyame

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"net/http"
)

type RefundRequest struct {
	Amount        decimal.Decimal `json:"amount"`
	PrepaidAmount decimal.Decimal `json:"refunded_prepaid_amount"`
	Items         []*Item         `json:"returned_items"`
}

func NewRefundRequest() *RefundRequest {
	return new(RefundRequest)
}

func (rr *RefundRequest) SetAmount(amount decimal.Decimal) *RefundRequest {
	rr.Amount = amount
	return rr
}

func (rr *RefundRequest) SetPrepaidAmount(amount decimal.Decimal) *RefundRequest {
	rr.PrepaidAmount = amount
	return rr
}

func (rr *RefundRequest) AddItem(item *Item) *RefundRequest {
	rr.Items = append(rr.Items, item)
	return rr
}

type RefundResponse struct {
	Amount   decimal.Decimal `json:"amount"`
	RefundID string          `json:"refund_id"`
}

func (c *Client) Refund(request *RefundRequest, orderID string, correlationID string) (*RefundResponse, error) {
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("https://%v/v1/orders/%v/refund", c.host, orderID),
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

	var resp RefundResponse

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
