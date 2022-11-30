package api

import ()

type GetDepositAddressSettings struct {
	CoinName string `json:"coin_name"` // ticker (example: BTC)
	NeedNew  uint64 `json:"need_new"`  // value: 0 or 1, on default: 0
}

type WithdrawCoinsToAddressSettings struct {
	CoinName string  `json:"coin_name"` // ticker (example: BTC)
	Amount   float64 `json:"amount"`    // amount to withdraw
	Address  string  `json:"address"`   // destination address
}

type GetDepositAddress struct {
	Success uint8     `json:"success"`
	Return  GDAReturn `json:"return"`
	Error   string    `json:"error"`
}

type GDAReturn struct {
	Address         string  `json:"address"`
	ProcessedAmount float64 `json:"processed_amount"`
	ServerTime      uint64  `json:"server_time"`
}

func NewGetDepositAddress() GetDepositAddress {
	getDepositAddress := GetDepositAddress{}
	getDepositAddress.Return = GDAReturn{}
	return getDepositAddress
}

type WithdrawCoinsToAddress struct {
	Success uint8      `json:"success"`
	Return  WCTAReturn `json:"return"`
	Error   string     `json:"error"`
}

type WCTAReturn struct {
	ServerTime uint64 `json:"server_time"`
}

func NewWithdrawCoinsToAddress() WithdrawCoinsToAddress {
	withdrawCoinsToAddress := WithdrawCoinsToAddress{}
	withdrawCoinsToAddress.Return = WCTAReturn{}
	return withdrawCoinsToAddress
}
