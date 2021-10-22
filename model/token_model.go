package model

type Token struct {
	ID         string      `json:"id"`
	CoinMarket interface{} `json:"coin_market"`
}

type Tokenlist struct {
	Data Tokendata `json:"data"`
}

type Tokendata struct {
	Coin Tokendetail `json:"btc"`
}

type Tokendetail struct {
	Currency interface{} `json:"quote"`
}

type CoinMarketQuote struct {
	Currency         string  `json:"currency"`
	Price            float64 `json:"price"`
	Coin             string  `json:"coin"`
	PercentChange24H float64 `json:"percentChange24H"`
	Slug             string  `json:"slug"`
}
