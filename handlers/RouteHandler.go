package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dellosaneil/stocktracking-backend/api/api_get"
	"github.com/dellosaneil/stocktracking-backend/constants"
	"github.com/dellosaneil/stocktracking-backend/indicators"
	"github.com/dellosaneil/stocktracking-backend/util"
	"github.com/gorilla/mux"
)

func HandleRoutes(router *mux.Router) {
	router.HandleFunc("/api/price", retrievePrice).Methods("GET")
	router.HandleFunc("/api/indicator/sma", simpleMovingAverage).Methods("GET")
	router.HandleFunc("/api/indicator/ema", exponentialMovingAverage).Methods("GET")
	router.HandleFunc("/api/indicator/macd", movingAverageConvergenceDivergence).Methods("GET")
	router.HandleFunc("/api/indicator/rsi", relativeStrengthIndex).Methods("GET")
	router.HandleFunc("/api/indicator/stochastic", stochasticOscillator).Methods("GET")
	router.HandleFunc("/api/indicator/vwap", volumeWeightedAveragePrice).Methods("GET")
}

func retrievePrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	urlParams := r.URL.Query()
	ticker := urlParams["stockTicker"][0]
	timeSeries := urlParams["timeseries"][0]
	prices, err := api_get.GetPriceCall(ticker, timeSeries)
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(prices)

}

func simpleMovingAverage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	priceType := "close"
	urlParams := r.URL.Query()
	ticker := urlParams["stockTicker"][0]
	timeSeries := urlParams["timeseries"][0]
	period := urlParams["period"][0]
	priceType = urlParams["priceType"][0]
	periodInt, periodErr := strconv.Atoi(period)
	if periodErr != nil {
		periodInt = 14
	}
	prices, err := api_get.GetPriceCall(ticker, timeSeries)
	if err != nil {
		fmt.Println(err)
	}
	price := util.PriceType(prices, priceType)
	sma := indicators.SimpleMovingAverage(price, periodInt)
	json.NewEncoder(w).Encode(sma)
}

func exponentialMovingAverage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	priceType := "close"
	urlParams := r.URL.Query()
	ticker := urlParams["stockTicker"][0]
	timeSeries := urlParams["timeseries"][0]
	period := urlParams["period"][0]
	priceType = urlParams["priceType"][0]
	periodInt, periodErr := strconv.Atoi(period)
	if periodErr != nil {
		periodInt = 14
	}
	prices, err := api_get.GetPriceCall(ticker, timeSeries)
	if err != nil {
		fmt.Println(err)
	}
	price := util.PriceType(prices, priceType)
	sma := indicators.ExponentialMovingAverage(price, periodInt)
	json.NewEncoder(w).Encode(sma)
}

func movingAverageConvergenceDivergence(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	urlParams := r.URL.Query()
	ticker := urlParams["stockTicker"][0]
	timeSeries := urlParams["timeseries"][0]
	priceType := urlParams["priceType"][0]
	fastPeriod := urlParams["fastPeriod"][0]
	slowPeriod := urlParams["slowPeriod"][0]
	signalPeriod := urlParams["signalPeriod"][0]
	prices, err := api_get.GetPriceCall(ticker, timeSeries)
	if err != nil {
		fmt.Println(err)
	}
	fastPeriodInt, errFast := strconv.Atoi(fastPeriod)
	slowPeriodInt, errSlow := strconv.Atoi(slowPeriod)
	signalPeriodInt, errSignal := strconv.Atoi(signalPeriod)
	if errFast != nil {
		fastPeriodInt = 12
	}
	if errSlow != nil {
		slowPeriodInt = 26
	}
	if errSignal != nil {
		signalPeriodInt = 9
	}

	price := util.PriceType(prices, priceType)
	sma := indicators.MovingAverageConvergenceDivergence(price, fastPeriodInt, slowPeriodInt, signalPeriodInt)
	json.NewEncoder(w).Encode(sma)
}

func relativeStrengthIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	priceType := "close"
	urlParams := r.URL.Query()
	ticker := urlParams["stockTicker"][0]
	timeSeries := urlParams["timeseries"][0]
	period := urlParams["period"][0]
	priceType = urlParams["priceType"][0]
	periodInt, periodErr := strconv.Atoi(period)
	if periodErr != nil {
		periodInt = 14
	}
	prices, err := api_get.GetPriceCall(ticker, timeSeries)
	if err != nil {
		fmt.Println(err)
	}
	price := util.PriceType(prices, priceType)
	sma := indicators.RelativeStrengthIndex(price, periodInt)
	json.NewEncoder(w).Encode(sma)
}

func stochasticOscillator(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	urlParams := r.URL.Query()
	ticker := urlParams["stockTicker"][0]
	timeSeries := urlParams["timeseries"][0]
	fastKPeriod := urlParams["fastKPeriod"][0]
	slowKPeriod := urlParams["slowKPeriod"][0]
	slowDPeriod := urlParams["slowDPeriod"][0]
	prices, err := api_get.GetPriceCall(ticker, timeSeries)
	if err != nil {
		fmt.Println(err)
	}
	fastKPeriodInt, errFast := strconv.Atoi(fastKPeriod)
	slowKPeriodInt, errSlow := strconv.Atoi(slowKPeriod)
	slowDPeriodInt, errSignal := strconv.Atoi(slowDPeriod)
	if errFast != nil {
		fastKPeriodInt = 14
	}
	if errSlow != nil {
		slowKPeriodInt = 3
	}
	if errSignal != nil {
		slowDPeriodInt = 3
	}
	sma := indicators.StochasticOscillator(prices, fastKPeriodInt, slowKPeriodInt, slowDPeriodInt)
	json.NewEncoder(w).Encode(sma)
}

func volumeWeightedAveragePrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	urlParams := r.URL.Query()
	ticker := urlParams["stockTicker"][0]
	prices, err := api_get.GetPriceCall(ticker, constants.INTRADAY)
	if err != nil {
		fmt.Println(err)
	}
	sma := indicators.VolumeWeightedAveragePrice(prices)
	json.NewEncoder(w).Encode(sma)
}
