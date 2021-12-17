package dolyame

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"net/http"
)

type CreateRequest struct {
	Order           *Order      `json:"order"`
	ClientInfo      *ClientInfo `json:"client_info,omitempty"`
	NotificationURL string      `json:"notification_url"`
	FailURL         string      `json:"fail_url"`
	SuccessURL      string      `json:"success_url"`
}

func NewCreateRequest() *CreateRequest {
	return new(CreateRequest)
}

func (cr *CreateRequest) SetOrder(order *Order) *CreateRequest {
	cr.Order = order
	return cr
}

func (cr *CreateRequest) SetClientInfo(info *ClientInfo) *CreateRequest {
	cr.ClientInfo = info
	return cr
}

func (cr *CreateRequest) SetNotificationURL(url string) *CreateRequest {
	cr.NotificationURL = url
	return cr
}

func (cr *CreateRequest) SetFailURL(url string) *CreateRequest {
	cr.FailURL = url
	return cr
}

func (cr *CreateRequest) SetSuccessURL(url string) *CreateRequest {
	cr.SuccessURL = url
	return cr
}

type CreateResponse struct {
	Status          string            `json:"status"`
	Amount          decimal.Decimal   `json:"amount"`
	ResidualAmount  decimal.Decimal   `json:"residual_amount"`
	Link            string            `json:"link"`
	PaymentSchedule []PaymentSchedule `json:"payment_schedule"`
	RefundInfo      RefundInfo        `json:"refund_info"`
}

func (c *Client) Create(request *CreateRequest, correlationID string) (*CreateResponse, error) {
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("https://%v/v1/orders/create", c.host),
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

	var resp CreateResponse

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
