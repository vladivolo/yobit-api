package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// API is the Public API that included in the main client.
type PublicAPI struct {
	apiKey    string
	apiSecret string
}

// NewAPI creates and returns the Public API to the main client.
func NewPublicAPI(api_key string, api_secret string) *PublicAPI {
	return &PublicAPI{
		apiKey:    api_key,
		apiSecret: api_secret,
	}
}

// Trades returns information about the last transactions of selected pairs.
func (api *PublicAPI) Trades(t *TradesSettings) (Trades, error) {
	values, link := api.createLinkTrades(t)

	body, err := api.sendRequest(values, link)
	if err != nil {
		return Trades{}, err
	}

	trades := NewTrades()
	err = json.Unmarshal(body, &trades.PairData)
	if err != nil {
		return Trades{}, err
	}

	return trades, err
}

// Info returns information about server time and active pairs.
func (api *PublicAPI) Info() (Info, error) {
	values, link := api.createLinkInfo()

	body, err := api.sendRequest(values, link)
	if err != nil {
		return Info{}, err
	}

	info := Info{
		Pairs: map[string]map[string]interface{}{},
	}

	err = json.Unmarshal(body, &info)
	if err != nil {
		return Info{}, err
	}

	return info, err
}

// return first Ask & Bid
func (api *PublicAPI) OpenInterest(symbol string) (float64, float64, error) {
	ticker, err := api.Ticker(
		&TickerSettings{
			Pairs: []string{symbol},
		})
	if err != nil {
		return 0, 0, err
	}

	tdata := ticker.PairData[symbol]

	return tdata.Sell, tdata.Buy, nil
}

// Ticker provides statistic data for the last 24 hours.
func (api *PublicAPI) Ticker(t *TickerSettings) (Ticker, error) {
	values, link := api.createLinkTicker(t)

	body, err := api.sendRequest(values, link)
	if err != nil {
		return Ticker{}, err
	}

	ticker := NewTicker()
	err = json.Unmarshal(body, &ticker.PairData)
	if err != nil {
		return Ticker{}, err
	}

	return ticker, err
}

// Depth returns information about lists of active orders for selected pairs.
func (api *PublicAPI) Depth(t *DepthSettings) (Depth, error) {
	values, link := api.createLinkDepth(t)

	body, err := api.sendRequest(values, link)
	if err != nil {
		return Depth{}, err
	}

	depth := NewDepth()
	err = json.Unmarshal(body, &depth.PairData)
	if err != nil {
		return Depth{}, err
	}

	return depth, err
}

// sendRequest prepares and sends request to server by calling objective functions and returns the body of response
func (api *PublicAPI) sendRequest(values *url.Values, link string) ([]byte, error) {
	req, err := api.prepareRequest(values, link)
	if err != nil {
		return []byte{}, err
	}

	resp, err := api.sendPost(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, err
}

// prepareRequest creates link and prepares request to send
func (api *PublicAPI) prepareRequest(values *url.Values, link string) (*http.Request, error) {
	requestString := values.Encode()

	req, err := http.NewRequest("POST", link, strings.NewReader(requestString))
	if err != nil {
		return nil, err
	}

	return req, err
}

// sendPost sends POST request to the API server
func (api *PublicAPI) sendPost(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (api *PublicAPI) createLinkInfo() (*url.Values, string) {
	values := url.Values{}
	link := PublicApiLink + "info"

	return &values, link

}

func (api *PublicAPI) createLinkTicker(th *TickerSettings) (*url.Values, string) {
	values := url.Values{}
	pairs := strings.Join(th.Pairs, "-")
	link := PublicApiLink + "ticker" + "/" + pairs

	return &values, link

}

func (api *PublicAPI) createLinkDepth(th *DepthSettings) (*url.Values, string) {
	values := url.Values{}

	if th.Limit != 0 {
		values.Add("limit", strconv.FormatUint(th.Limit, 10))
	}

	link := PublicApiLink + "depth" + "/" + th.Pair

	return &values, link

}

func (api *PublicAPI) createLinkTrades(th *TradesSettings) (*url.Values, string) {
	values := url.Values{}

	if th.Limit != 0 {
		values.Add("limit", strconv.FormatUint(th.Limit, 10))
	}

	link := PublicApiLink + "trades" + "/" + th.Pair

	return &values, link
}
