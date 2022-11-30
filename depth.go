package api

import ()

type DepthSettings struct {
	Pair  string `json:"pair"`  // pair (example: ltc_btc)
	Limit uint64 `json:"limit"` // limit stipulates size of withdrawal (on default 150 to 2000 maximum)
}

type Depth struct {
	Success  uint8 `json:"success"`
	PairData map[string]PData
	Error    string `json:"error"`
}

type PData struct {
	Asks [][2]float64 `json:"asks"` // selling orders
	Bids [][2]float64 `json:"bids"` // buying orders
}

func NewDepth() Depth {
	depth := Depth{}
	depth.PairData = make(map[string]PData)
	return depth
}
