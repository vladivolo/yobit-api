package api

import ()

type ActiveOrdersSettings struct {
	Pair string `json:"pair"` // pair (example: ltc_btc)
}

type OrderInfoSettings struct {
	OrderID uint64 `json:"order_id"` // order ID
}

type CancelOrderSettings struct {
	OrderID uint64 `json:"order_id"` // order ID
}

type ActiveOrders struct {
	Success uint8                             `json:"success"`
	Return  map[uint64]map[string]interface{} `json:"return"`
	Error   string                            `json:"error"`
}

type TradeHistorySettings struct {
	From   uint64 `json:"from"`    // No. of transaction from which withdrawal starts (value: numeral, on default: 0)
	Count  uint64 `json:"count"`   // quantity of withrawal transactions (value: numeral, on default: 1000)
	FromID uint64 `json:"from_id"` // ID of transaction from which withdrawal starts (value: numeral, on default: 0)
	EndID  uint64 `json:"end_id"`  // ID of transaction at which withdrawal finishes (value: numeral, on default: ∞)
	Order  string `json:"order"`   // sorting at withdrawal (value: ASC or DESC, on default: DESC)
	Since  uint64 `json:"since"`   // the time to start the display (value: unix time, on default: 0)
	End    uint64 `json:"end"`     // the time to end the display (value: unix time, on default: ∞)
	Pair   string `json:"pair"`    // pair (example: ltc_btc)
}

func NewActiveOrders() ActiveOrders {
	activeOrders := ActiveOrders{}
	activeOrders.Return = make(map[uint64]map[string]interface{})
	return activeOrders
}

type OrderInfo struct {
	Success uint8                          `json:"success"`
	Return  map[int]map[string]interface{} `json:"return"`
	Error   string                         `json:"error"`
}

func NewOrderInfo() OrderInfo {
	orderInfo := OrderInfo{}
	orderInfo.Return = make(map[int]map[string]interface{})
	return orderInfo
}

type CancelOrder struct {
	Success uint8  `json:"success"`
	Return  COData `json:"return"`
	Error   string `json:"error"`
}

type COData struct {
	OrderID int                    `json:"order_id"` // order ID
	Funds   map[string]interface{} `json:"funds"`    // balances active after request
}

func NewCancelOrder() CancelOrder {
	cancelOrder := CancelOrder{}
	cancelOrder.Return.Funds = make(map[string]interface{})
	return cancelOrder
}

type TradeHistory struct {
	Success uint8                       `json:"success"`
	Return  map[int]map[string]THReturn `json:"return"`
	Error   string                      `json:"error"`
}

type THReturn struct {
	Pair        string  `json:"pair"`          // pair
	Type        string  `json:"type"`          // transaction type
	Amount      float64 `json:"amount"`        // amount
	Rate        float64 `json:"rate"`          // price of buying or selling
	OrderID     int     `json:"order_id"`      // order ID
	IsYourOrder bool    `json:"is_your_order"` // is the order yours
	Timestamp   int     `json:"timestamp"`     // transaction time
}

func NewTradeHistory() TradeHistory {
	tradeHistory := TradeHistory{}
	tradeHistory.Return = make(map[int]map[string]THReturn)
	return tradeHistory
}
