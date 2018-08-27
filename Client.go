package mws

type Account struct {
}

type Client struct {
	Acount Acount
}

type Request map[string]interface{}

func NewClient(account Account) *Client {
	return &Client{Account: account}
}
