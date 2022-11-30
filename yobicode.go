package api

import ()

type CreateYobicodeSettings struct {
	Currency string  `json:"coin_name"` // ticker (example: BTC)
	Amount   float64 `json:"amount"`    // amount to withdraw
}

type RedeemYobicodeSettings struct {
	Coupon string `json:"coupon"` // yobicode to redeem (example: YOBITUZ0HHSTB...OQX3H01BTC)
}

type CreateYobicode struct {
	Success uint8    `json:"success"`
	Return  CYReturn `json:"return"`
	Error   string   `json:"error"`
}

type CYReturn struct {
	Coupon  string             `json:"coupon"`  // Yobicode
	TransID uint8              `json:"transID"` // always 1 for compatibility with api of other exchanges
	Funds   map[string]float64 `json:"funds"`   // balances active after request
}

func NewCreateYobicode() CreateYobicode {
	createYobicode := CreateYobicode{}
	createYobicode.Return = CYReturn{}
	return createYobicode
}

type RedeemYobicode struct {
	Success uint8    `json:"success"`
	Return  RYReturn `json:"return"`
	Error   string   `json:"error"`
}

type RYReturn struct {
	CouponAmount   float64            `json:"couponAmount"`   // The amount that has been redeemed.
	CouponCurrency string             `json:"couponCurrency"` // The currency of the yobicode that has been redeemed.
	TransID        uint8              `json:"transID"`        // always 1 for compatibility with api of other exchanges
	Funds          map[string]float64 `json:"funds"`          // balances active after request
}

func NewRedeemYobicode() RedeemYobicode {
	redeemYobicode := RedeemYobicode{}
	redeemYobicode.Return = RYReturn{}
	return redeemYobicode
}
