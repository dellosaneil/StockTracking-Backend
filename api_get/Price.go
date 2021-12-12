package api_get

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"

	"github.com/dellosaneil/stocktracking-backend/api_response"
	"github.com/dellosaneil/stocktracking-backend/model"
	"github.com/dellosaneil/stocktracking-backend/util"
)

func GetPriceCall() ([]model.PriceModel, error) {
	var price []model.PriceModel
	response, errApi := http.Get("https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=IBM&apikey=demo")

	if errApi != nil {
		return price, errApi
	}
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return price, err
	}
	var ts api_response.DailyPriceResponse
	err = json.Unmarshal(contents, &ts)
	if err != nil {
		return price, err
	}
	keys := make([]string, 0, len(ts.Item))
	for k := range ts.Item {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, t := range keys {
		open, errOpen := strconv.ParseFloat(ts.Item[t]["1. open"], 32)
		high, errHigh := strconv.ParseFloat(ts.Item[t]["2. high"], 32)
		low, errLow := strconv.ParseFloat(ts.Item[t]["3. low"], 32)
		close, closeErr := strconv.ParseFloat(ts.Item[t]["4. close"], 32)
		volume, errVolume := strconv.ParseInt(ts.Item[t]["5. volume"], 10, 32)
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
