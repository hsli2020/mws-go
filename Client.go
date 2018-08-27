package mws

type Account struct {
    SellerId      string
    AccessKey     string
    SecretKey     string
    MarketplaceId string
    AuthToken     string
}

type Client struct {
	Account
}

type Request map[string]interface{}

func NewClient(account Account) *Client {
	return &Client{Account: account}
}
