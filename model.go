package main

type WazirxMarkets struct {
	Markets []MarketStatus `json:markets`
	Assets  []Assets       `json:assets`
}

type MarketStatus struct {
	BaseMarket     string      `json:baseMarket`
	QuoteMarket    string      `json:quoteMarket`
	MinBuyAmount   float32     `json:minBuyAmount`
	MinSellAmount  float32     `json:minSellAmount`
	BasePrecision  int         `json:basePrecision`
	QuotePrecision int         `json:quotePrecision`
	Status         string      `json:status`
	Low            string      `json:low`
	high           string      `json:high`
	Last           string      `json:last`
	Type           string      `json:type`
	Open           interface{} `json:open`
	Volume         string      `json:volume`
	Sell           string      `json:sell`
	Buy            string      `json:buy`
	At             int64       `json:at`
}

type Assets struct {
	Type              string  `json:type`
	Name              string  `json:name`
	Deposit           string  `json:deposit`
	Withdrawal        string  `json:withdrawal`
	ListingType       string  `json:listingType`
	Category          string  `json:category`
	WithdrawFee       float32 `json:withdrawFee`
	MinWithdrawAmount float32 `json:minWithdrawAmount`
	MaxWithdrawAmount float32 `json:maxWithdrawAmount`
	MinDepositAmount  float32 `json:minDepositAmount`
	Confirmations     int     `json:confirmations`
}
