package yahoo

type yahooResponse struct {
	Chart chart `json:"chart"`
}

type chart struct {
	Result []result `json:"result"`
}

type result struct {
	Meta meta `json:"meta"`
}

type meta struct {
	Currency           string  `json:"currency"`
	Symbol             string  `json:"symbol"`
	RegularMarketPrice float64 `json:"regularMarketPrice"`
}
