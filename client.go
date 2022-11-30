package api

import ()

// Client is a main struct for executing all the operations of the API
type Client struct {
	apiKey    string
	apiSecret string

	Public *PublicAPI
	Trade  *TradeAPI
}

// NewClient is a constructor for the Client
func NewClient(api_key string, api_secret string) *Client {
	client := &Client{
		apiKey:    api_key,
		apiSecret: api_secret,
	}

	client.Public = NewPublicAPI(api_key, api_secret)
	client.Trade = NewTradeAPI(api_key, api_secret)

	return client
}
