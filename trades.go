package api

type Trade struct {
	Success uint8       `json:"success"`
	Return  TradeReturn `json:"return"`
	Error   string      `json:"error"`
}

type TradeReturn struct {
	Received float64            `json:"received"` // amount of currency bought / sold
	Remains  float64            `json:"remains"`  // amount of currency to buy / to sell
	OrderID  int                `json:"order_id"` // created order ID
	Funds    map[string]float64 `json:"funds"`    // funds active after request
}

func NewTrade() Trade {
	return Trade{
		Return: TradeReturn{
			Funds: map[string]float64{},
		},
	}
}

type Trades struct {
	Success  int                    `json:"success"`
	PairData map[string][]TradeData `json:"return"`
	Error    string                 `json:"error"`
}

type TradeData struct {
	Type      string  `json:"type"`      // ask - sell, bid - buy
	Price     float64 `json:"price"`     // buying / selling price
	Amount    float64 `json:"amount"`    // amount
	Tid       uint    `json:"tid"`       // transaction id
	Timestamp int64   `json:"timestamp"` // transaction timestamp
}

type TradeSettings struct {
	Pair   string  `json:"pair"`   // pair (example: ltc_btc)
	Type   string  `json:"type"`   // transaction type (example: buy or sell)
	Rate   float64 `json:"rate"`   // exchange rate for buying or selling (value: numeral)
	Amount float64 `json:"amount"` // amount needed for buying or selling (value: numeral)
}

type TradesSettings struct {
	Pair  string `json:"pair"`  // pair (example: ltc_btc)
	Limit uint64 `json:"limit"` // limit stipulates size of withdrawal (on default 150 to 2000 maximum)
}

// NewTrades returns new structure for Trade response
func NewTrades() Trades {
	return Trades{
		PairData: map[string][]TradeData{},
	}
}

// Separate separates asks and bids from the main result
func Separate(t Trades, pair string) ([]TradeData, []TradeData) {
	var asks, bids []TradeData
	for i := 0; i < len(t.PairData[pair]); i++ {
		switch t.PairData[pair][i].Type {
		case "ask":
			asks = append(asks, t.PairData[pair][i])
		case "bid":
			bids = append(bids, t.PairData[pair][i])
		}
	}
	return asks, bids
}

// GetPriceBefore returns first action price before specified timestamp
func GetPriceBefore(tds []TradeData, before int64) (price float64) {
	var currentTimeVal int64
	for _, val := range tds {
		if val.Timestamp < before && val.Timestamp > currentTimeVal {
			currentTimeVal = val.Timestamp
			price = val.Price
		}
	}
	return
}
