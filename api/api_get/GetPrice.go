package api_get

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"

	"github.com/dellosaneil/stocktracking-backend/api/api_response"
	"github.com/dellosaneil/stocktracking-backend/constants"
	"github.com/dellosaneil/stocktracking-backend/model"
	"github.com/dellosaneil/stocktracking-backend/util"
)

func GetPriceCall(stockTicker string, timeSeries string) ([]model.PriceModel, error) {
	var price []model.PriceModel
	response, errApi := http.Get(fmt.Sprintf("https://www.alphavantage.co/query?function=%s&symbol=%s&apikey=90ZLSXISLFOGZIJO", timeSeries, stockTicker))

	if errApi != nil {
		return price, errApi
	}
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return price, err
	}
	var ts, unmarshalError = getPrice(contents, timeSeries)

	if unmarshalError != nil {
		return price, unmarshalError
	}
	keys := make([]string, 0, len(ts))
	for k := range ts {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, t := range keys {
		open, errOpen := strconv.ParseFloat(ts[t]["1. open"], 32)
		high, errHigh := strconv.ParseFloat(ts[t]["2. high"], 32)
		low, errLow := strconv.ParseFloat(ts[t]["3. low"], 32)
		close, closeErr := strconv.ParseFloat(ts[t]["4. close"], 32)
		volume, errVolume := strconv.ParseInt(ts[t]["5. volume"], 10, 32)
		if errVolume != nil || errOpen != nil || errLow != nil || closeErr != nil || errHigh != nil {
			return price, errOpen
		}

		tempPrice := model.PriceModel{
			util.RoundPrecision(open, 4),
			util.RoundPrecision(high, 4),
			util.RoundPrecision(low, 4),
			util.RoundPrecision(close, 4),
			volume,
		}
		price = append(price, tempPrice)
	}
	return price, nil
}

func getPrice(contents []byte, timeSeries string) (map[string]api_response.PriceResponse, error) {
	switch timeSeries {
	case constants.INTRADAY:
		var ts api_response.IntradayPriceResponse
		err := json.Unmarshal(contents, &ts)
		return ts.Item, err
	case constants.DAILY:
		var ts api_response.DailyPriceResponse
		err := json.Unmarshal(contents, &ts)
		return ts.Item, err
	case constants.WEEKLY:
		var ts api_response.WeeklyPriceResponse
		err := json.Unmarshal(contents, &ts)
		return ts.Item, err
	case constants.MONTHLY:
		var ts api_response.MonthlyPriceResponse
		err := json.Unmarshal(contents, &ts)
		return ts.Item, err
	}
	return make(map[string]api_response.PriceResponse), nil
}
