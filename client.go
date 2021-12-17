package dolyame

import (
	"net/http"
)

type Client struct {
	httpClient *http.Client
	login      string
	password   string
	host       string
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

func (c *Client) SetHttpClient(client *http.Client) *Client {
	c.httpClient = client
	return c
}
