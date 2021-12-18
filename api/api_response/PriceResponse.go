package api_response

type PriceResponse map[string]string

type DailyPriceResponse struct {
	Item map[string]PriceResponse `json:"Time Series (Daily)"`
}

type IntradayPriceResponse struct {
	Item map[string]PriceResponse `json:"Time Series (1min)"`
}

type WeeklyPriceResponse struct {
	Item map[string]PriceResponse `json:"Weekly Time Series"`
}

type MonthlyPriceResponse struct {
	Item map[string]PriceResponse `json:"Monthly Time Series"`
}
