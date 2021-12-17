package dolyame

import (
	"github.com/shopspring/decimal"
)

type Order struct {
	ID            string          `json:"id"`
	Amount        decimal.Decimal `json:"amount"`
	PrepaidAmount decimal.Decimal `json:"prepaid_amount"`
	Items         []*Item         `json:"items"`
}

func NewOrder() *Order {
	return new(Order)
}

func (o *Order) SetID(id string) *Order {
	o.ID = id
	return o
}

func (o *Order) SetAmount(amount decimal.Decimal) *Order {
	o.Amount = amount
	return o
}

func (o *Order) SetPrepaidAmount(amount decimal.Decimal) *Order {
	o.PrepaidAmount = amount
	return o
}

func (o *Order) AddItem(item *Item) *Order {
	o.Items = append(o.Items, item)
	return o
}

type Item struct {
	Name     string          `json:"name"`
	Quantity int32           `json:"quantity"`
	Price    decimal.Decimal `json:"price"`
	SKU      string          `json:"sku"`
}

func NewItem() *Item {
	return new(Item)
}

func (i *Item) SetQuantity(quantity int32) *Item {
	i.Quantity = quantity
	return i
}

func (i *Item) SetName(name string) *Item {
	i.Name = name
	return i
}

func (i *Item) SetPrice(price decimal.Decimal) *Item {
	i.Price = price
	return i
}

type ClientInfo struct {
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	BirthDate  string `json:"birthdate"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`
}

func NewClientInfo() *ClientInfo {
	return new(ClientInfo)
}

func (ci *ClientInfo) SetPhone(phone string) *ClientInfo {
	ci.Phone = phone
	return ci
}

func (ci *ClientInfo) SetEmail(email string) *ClientInfo {
	ci.Email = email
	return ci
}

func (ci *ClientInfo) SetBirthDate(date string) *ClientInfo {
	ci.BirthDate = date
	return ci
}

func (ci *ClientInfo) SetFirstName(name string) *ClientInfo {
	ci.FirstName = name
	return ci
}

func (ci *ClientInfo) SetLastName(name string) *ClientInfo {
	ci.LastName = name
	return ci
}

func (ci *ClientInfo) SetMiddleName(name string) *ClientInfo {
	ci.MiddleName = name
	return ci
}

type PaymentSchedule struct {
	Amount decimal.Decimal `json:"amount"`
	Date   string          `json:"date"`
	Status string          `json:"status"`
}

type RefundItems struct {
	RefundID              string          `json:"refund_id"`
	RefundedAmount        decimal.Decimal `json:"refunded_amount"`
	RefundedPrepaidAmount decimal.Decimal `json:"refunded_prepaid_amount"`
	ReturnedItems         []Item          `json:"returned_items"`
	Status                string          `json:"status"`
}

type RefundInfo struct {
	TotalRefundedAmount        decimal.Decimal `json:"total_refunded_amount"`
	TotalRefundedPrepaidAmount decimal.Decimal `json:"total_refunded_prepaid_amount"`
	Items                      []RefundItems   `json:"items"`
}
