package api

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// API is the Trade API that included in the main client
type TradeAPI struct {
	apiKey    string
	apiSecret string

	VirtualNonce bool // is for saving to file or not (false = to file by edfault)
	Nonce        int  // current nonce parameter
}

// NewAPI creates and returns the Trade API to the main client
func NewTradeAPI(api_key string, api_secret string) *TradeAPI {
	return &TradeAPI{
		apiKey:    api_key,
		apiSecret: api_secret,
	}
}

// GetInfo shows info about account's balance.
func (api *TradeAPI) GetInfo() (GetInfo, error) {
	values := api.createLinkGetInfo()

	body, err := api.sendRequest(values)
	if err != nil {
		return GetInfo{}, err
	}

	balance := NewBalance()
	err = json.Unmarshal(body, &balance)
	if err != nil {
		return GetInfo{}, err
	}

	return balance, nil
}

// Trade allows creating new orders.
func (api *TradeAPI) Trade(t *TradeSettings) (Trade, error) {
	values := api.createLinkTrade(t)

	body, err := api.sendRequest(values)
	if err != nil {
		return Trade{}, err
	}

	trade := NewTrade()
	err = json.Unmarshal(body, &trade)
	if err != nil {
		return Trade{}, err
	}

	return trade, err
}

// ActiveOrders returns list of user's active orders.
func (api *TradeAPI) ActiveOrders(t *ActiveOrdersSettings) (ActiveOrders, error) {
	values := api.createLinkActiveOrders(t)

	body, err := api.sendRequest(values)
	if err != nil {
		return ActiveOrders{}, err
	}

	activeOrders := NewActiveOrders()
	err = json.Unmarshal(body, &activeOrders)
	if err != nil {
		return ActiveOrders{}, err
	}

	return activeOrders, err
}

// OrderInfo returns detailed information about the chosen order.
func (api *TradeAPI) OrderInfo(t *OrderInfoSettings) (OrderInfo, error) {
	values := api.createLinkOrderInfo(t)

	body, err := api.sendRequest(values)
	if err != nil {
		return OrderInfo{}, err
	}

	orderInfo := NewOrderInfo()
	err = json.Unmarshal(body, &orderInfo)
	if err != nil {
		return OrderInfo{}, err
	}

	return orderInfo, err
}

// CancelOrder cancells the chosen order.
func (api *TradeAPI) CancelOrder(t *CancelOrderSettings) (CancelOrder, error) {
	values := api.createLinkCancelOrder(t)

	body, err := api.sendRequest(values)
	if err != nil {
		return CancelOrder{}, err
	}

	cancelOrder := NewCancelOrder()
	err = json.Unmarshal(body, &cancelOrder)
	if err != nil {
		return CancelOrder{}, err
	}

	return cancelOrder, err
}

// TradeHistory returns transaction history.
func (api *TradeAPI) TradeHistory(t *TradeHistorySettings) (TradeHistory, error) {
	values := api.createLinkTradeHistory(t)

	body, err := api.sendRequest(values)
	if err != nil {
		return TradeHistory{}, err
	}

	fmt.Println(string(body))

	tradeHistory := NewTradeHistory()
	err = json.Unmarshal(body, &tradeHistory)
	if err != nil {
		return TradeHistory{}, err
	}

	return tradeHistory, err
}

// GetDepositAddress returns deposit address.
func (api *TradeAPI) GetDepositAddress(t *GetDepositAddressSettings) (GetDepositAddress, error) {
	values := api.createLinkGetDepositAddress(t)

	body, err := api.sendRequest(values)
	if err != nil {
		return GetDepositAddress{}, err
	}

	getDepositAddress := NewGetDepositAddress()
	err = json.Unmarshal(body, &getDepositAddress)
	if err != nil {
		return GetDepositAddress{}, err
	}

	return getDepositAddress, err
}

// WithdrawCoinsToAddress creates withdrawal request.
func (api *TradeAPI) WithdrawCoinsToAddress(t *WithdrawCoinsToAddressSettings) (WithdrawCoinsToAddress, error) {
	values := api.createLinkWithdrawCoinsToAddress(t)

	body, err := api.sendRequest(values)
	if err != nil {
		return WithdrawCoinsToAddress{}, err
	}

	getDepositAddress := NewWithdrawCoinsToAddress()
	err = json.Unmarshal(body, &getDepositAddress)
	if err != nil {
		return WithdrawCoinsToAddress{}, err
	}

	return getDepositAddress, err
}

// CreateYobicode allows you to create Yobicodes (coupons).
func (api *TradeAPI) CreateYobicode(t *CreateYobicodeSettings) (CreateYobicode, error) {
	values := api.createLinkCreateYobicode(t)

	body, err := api.sendRequest(values)
	if err != nil {
		return CreateYobicode{}, err
	}

	createYobicode := NewCreateYobicode()
	err = json.Unmarshal(body, &createYobicode)
	if err != nil {
		return CreateYobicode{}, err
	}

	return createYobicode, err
}

// RedeemYobicode is used to redeem Yobicodes (coupons).
func (api *TradeAPI) RedeemYobicode(t *RedeemYobicodeSettings) (RedeemYobicode, error) {
	values := api.createLinkRedeemYobicode(t)

	body, err := api.sendRequest(values)
	if err != nil {
		return RedeemYobicode{}, err
	}

	redeemYobicode := NewRedeemYobicode()
	err = json.Unmarshal(body, &redeemYobicode)
	if err != nil {
		return RedeemYobicode{}, err
	}

	return redeemYobicode, err
}

func (api *TradeAPI) createLinkGetInfo() *url.Values {
	nonce, err := api.GetNonce(api.apiKey)
	if err != nil {
		return nil
	}

	values := url.Values{
		"method": []string{"getInfo"},
		"nonce":  []string{strconv.Itoa(nonce)},
	}

	return &values

}

func (api *TradeAPI) createLinkTrade(th *TradeSettings) *url.Values {
	nonce, err := api.GetNonce(api.apiKey)
	if err != nil {
		return nil
	}

	values := url.Values{
		"method": []string{"Trade"},
		"nonce":  []string{strconv.Itoa(nonce)},
	}

	if th.Pair != "" {
		values.Add("pair", th.Pair)
	} else {
		return nil
	}
	if th.Type != "" {
		values.Add("type", th.Type)
	}
	if th.Rate != 0 {
		values.Add("rate", strconv.FormatFloat(th.Rate, 'f', -1, 64))
	}
	if th.Amount != 0 {
		values.Add("amount", strconv.FormatFloat(th.Amount, 'f', -1, 64))
	}

	return &values

}

func (api *TradeAPI) createLinkActiveOrders(th *ActiveOrdersSettings) *url.Values {
	nonce, err := api.GetNonce(api.apiKey)
	if err != nil {
		return nil
	}

	values := url.Values{
		"method": []string{"ActiveOrders"},
		"nonce":  []string{strconv.Itoa(nonce)},
	}

	if th.Pair != "" {
		values.Add("pair", th.Pair)
	} else {
		return nil
	}

	return &values

}

func (api *TradeAPI) createLinkOrderInfo(th *OrderInfoSettings) *url.Values {
	nonce, err := api.GetNonce(api.apiKey)
	if err != nil {
		return nil
	}

	values := url.Values{
		"method": []string{"OrderInfo"},
		"nonce":  []string{strconv.Itoa(nonce)},
	}

	if th.OrderID != 0 {
		values.Add("order_id", strconv.FormatUint(th.OrderID, 10))
	} else {
		return nil
	}

	return &values

}

func (api *TradeAPI) createLinkCancelOrder(th *CancelOrderSettings) *url.Values {
	nonce, err := api.GetNonce(api.apiKey)
	if err != nil {
		return nil
	}

	values := url.Values{
		"method": []string{"CancelOrder"},
		"nonce":  []string{strconv.Itoa(nonce)},
	}

	if th.OrderID != 0 {
		values.Add("order_id", strconv.FormatUint(th.OrderID, 10))
	} else {
		return nil
	}

	return &values

}

func (api *TradeAPI) createLinkTradeHistory(th *TradeHistorySettings) *url.Values {
	nonce, err := api.GetNonce(api.apiKey)
	if err != nil {
		return nil
	}

	values := url.Values{
		"method": []string{"TradeHistory"},
		"nonce":  []string{strconv.Itoa(nonce)},
	}

	if th.From != 0 {
		values.Add("From", strconv.FormatUint(th.From, 10))
	}
	if th.Count != 0 {
		values.Add("Count", strconv.FormatUint(th.Count, 10))
	}
	if th.FromID != 0 {
		values.Add("FromID", strconv.FormatUint(th.FromID, 10))
	}
	if th.EndID != 0 {
		values.Add("EndID", strconv.FormatUint(th.EndID, 10))
	}
	if th.Order != "" {
		values.Add("Order", th.Order)
	}
	if th.Since != 0 {
		values.Add("Since", strconv.FormatUint(th.Since, 10))
	}
	if th.End != 0 {
		values.Add("End", strconv.FormatUint(th.End, 10))
	}
	if th.Pair != "" {
		values.Add("pair", th.Pair)
	} else {
		return nil
	}

	return &values

}

func (api *TradeAPI) createLinkGetDepositAddress(th *GetDepositAddressSettings) *url.Values {
	nonce, err := api.GetNonce(api.apiKey)
	if err != nil {
		return nil
	}

	values := url.Values{
		"method": []string{"GetDepositAddress"},
		"nonce":  []string{strconv.Itoa(nonce)},
	}

	if th.CoinName != "" {
		values.Add("coinName", th.CoinName)
	} else {
		panic("createLinkGetDepositAddress Pair hasn't been set")
	}
	if th.NeedNew != 0 {
		values.Add("need_new", strconv.FormatUint(th.NeedNew, 10))
	}

	return &values

}

func (api *TradeAPI) createLinkWithdrawCoinsToAddress(th *WithdrawCoinsToAddressSettings) *url.Values {
	nonce, err := api.GetNonce(api.apiKey)
	if err != nil {
		return nil
	}

	values := url.Values{
		"method": []string{"WithdrawCoinsToAddress"},
		"nonce":  []string{strconv.Itoa(nonce)},
	}

	if th.CoinName != "" {
		values.Add("coinName", th.CoinName)
	} else {
		panic("createLinkWithdrawCoinsToAddress Pair hasn't been set")
	}
	if th.Amount != 0 {
		values.Add("amount", strconv.FormatFloat(th.Amount, 'f', -1, 64))
	} else {
		panic("createLinkWithdrawCoinsToAddress Amount hasn't been set")
	}
	if th.Address != "" {
		values.Add("address", th.Address)
	} else {
		panic("createLinkWithdrawCoinsToAddress Address hasn't been set")
	}

	return &values

}

func (api *TradeAPI) createLinkCreateYobicode(th *CreateYobicodeSettings) *url.Values {
	nonce, err := api.GetNonce(api.apiKey)
	if err != nil {
		return nil
	}

	values := url.Values{
		"method": []string{"CreateYobicode"},
		"nonce":  []string{strconv.Itoa(nonce)},
	}

	if th.Currency != "" {
		values.Add("coinName", th.Currency)
	} else {
		panic("createLinkCreateYobicode Currency hasn't been set")
	}
	if th.Amount != 0 {
		values.Add("amount", strconv.FormatFloat(th.Amount, 'f', -1, 64))
	} else {
		panic("createLinkCreateYobicode Amount hasn't been set")
	}

	return &values

}

func (api *TradeAPI) createLinkRedeemYobicode(th *RedeemYobicodeSettings) *url.Values {
	nonce, err := api.GetNonce(api.apiKey)
	if err != nil {
		return nil
	}

	values := url.Values{
		"method": []string{"RedeemYobicode"},
		"nonce":  []string{strconv.Itoa(nonce)},
	}

	if th.Coupon != "" {
		values.Add("coupon", th.Coupon)
	} else {
		panic("createLinkRedeemYobicode Currency hasn't been set")
	}

	return &values

}

// sendRequest prepares and sends request to server by calling objective functions and returns the body of response
func (api *TradeAPI) sendRequest(values *url.Values) ([]byte, error) {
	req, err := api.prepareRequest(values)
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
func (api *TradeAPI) prepareRequest(values *url.Values) (*http.Request, error) {
	requestString := values.Encode()

	sign := hmac.New(sha512.New, []byte(api.apiSecret))
	sign.Write([]byte(requestString))

	req, err := http.NewRequest("POST", TradeApiLink, strings.NewReader(requestString))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Key", api.apiKey)
	req.Header.Add("Sign", hex.EncodeToString(sign.Sum(nil)))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return req, err
}

// sendPost sends POST request to the TradeAPI server
func (api *TradeAPI) sendPost(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// GetNonce is a maintenance function for getting and storing nonce counter
func (api *TradeAPI) GetNonce(Key string) (int, error) {
	nonceFileName := "nonce." + Key[0:8] + ".txt"
	if api.VirtualNonce != true {
		nonceBytes, err := ioutil.ReadFile(nonceFileName)
		if err == nil {
			api.Nonce, _ = strconv.Atoi(string(nonceBytes))
		}
	}
	api.Nonce++
	err := api.WriteNonce(api.Nonce, nonceFileName)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return api.Nonce, err
}

// WriteNonce writes nonce to the file
func (api *TradeAPI) WriteNonce(nonce int, nonceFileName string) error {
	return ioutil.WriteFile(nonceFileName, []byte(strconv.Itoa(nonce)), 0644)
}
