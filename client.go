package dolyame

import (
	"crypto/tls"
	"net/http"
)

type Client struct {
	certificate     tls.Certificate
	httpClient      *http.Client
	login           string
	password        string
	host            string
	notificationURL string
	successURL      string
	failURL         string
}

func NewClient() *Client {
	return new(Client)
}

func (c *Client) SetLogin(login string) *Client {
	c.login = login
	return c
}

func (c *Client) SetPassword(password string) *Client {
	c.password = password
	return c
}

func (c *Client) SetHost(host string) *Client {
	c.host = host
	return c
}

func (c *Client) SetNotificationURL(url string) *Client {
	c.notificationURL = url
	return c
}

func (c *Client) SetSuccessURL(url string) *Client {
	c.successURL = url
	return c
}

func (c *Client) SetFailURL(url string) *Client {
	c.failURL = url
	return c
}
