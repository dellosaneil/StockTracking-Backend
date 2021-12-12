package api_response

type MetaDataResponse struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	OutputSize    string `json:"4. Output Size"`
	TimeZone      string `json:"5. Time Zone"`
}

type values map[string]string

type DailyPriceResponse struct {
	Item map[string]values `json:"Time Series (Daily)"`
}

type IntradayPriceResponse struct {
	Item map[string]values `json:"Time Series (Daily)"`
}
