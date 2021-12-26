package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dellosaneil/stocktracking-backend/api/api_get"
	"github.com/dellosaneil/stocktracking-backend/constants"
	"github.com/dellosaneil/stocktracking-backend/util"
	"github.com/dellosaneil/stocktracking-backend/websockets/ws_indicators"
	"github.com/dellosaneil/stocktracking-backend/websockets/ws_price"
	"github.com/gorilla/mux"
)

func HandleRoutes(router *mux.Router) {
	router.HandleFunc("/api/price", retrievePrice).Methods("GET")
	router.HandleFunc("/api/websocket/price", wsPrice)
	router.HandleFunc("/api/websocket/indicator/sma", wsSimpleMovingAverage)
	router.HandleFunc("/api/websocket/indicator/ema", wsExponentialMovingAverage)
	router.HandleFunc("/api/websocket/indicator/macd", wsMovingAverageConvergenceDivergence)
	router.HandleFunc("/api/websocket/indicator/rsi", wsRelativeStrengthIndex)
	router.HandleFunc("/api/websocket/indicator/stochastic", wsStochasticOscillator)
	router.HandleFunc("/api/websocket/indicator/vwap", wsVolumeWeightedAveragePrice)
}

func retrievePrice(w http.ResponseWriter, r *http.Request) {
	var timeSeries string
	w.Header().Set("Content-Type", "application/json")
	urlParams := r.URL.Query()
	ticker := urlParams["stockTicker"][0]
	timeSeriesQuery := urlParams["timeseries"]
	if len(timeSeriesQuery) == 0 {
		timeSeries = constants.DAILY
	} else {
		timeSeries = timeSeriesQuery[0]
	}
	prices, err := api_get.GetPriceCall(ticker, timeSeries)
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(prices)

}

func wsPrice(w http.ResponseWriter, r *http.Request) {
	var timeSeries string
	urlParams := r.URL.Query()
	ticker := urlParams["stockTicker"][0]
	timeSeriesQuery := urlParams["timeseries"]
	if len(timeSeriesQuery) == 0 {
		timeSeries = constants.DAILY
	} else {
		timeSeries = timeSeriesQuery[0]
	}
	ws, err := ws_price.WSPriceUpgrade(w, r)
	if err != nil {
		panic(err)
	}
	go ws_price.WSPriceWriter(ws, ticker, timeSeries)
	go util.CloseWebSocketConnection(ws)
}

func wsSimpleMovingAverage(w http.ResponseWriter, r *http.Request) {
	var period, priceType string
	urlParams := r.URL.Query()
	periodQuery := urlParams["period"]
	priceTypeQuery := urlParams["priceType"]
	if len(periodQuery) == 0 {
		period = "14"
	} else {
		period = periodQuery[0]
	}
	if len(priceTypeQuery) == 0 {
		priceType = constants.CLOSE
	} else {
		priceType = priceTypeQuery[0]
	}

	periodInt, periodErr := strconv.Atoi(period)
	if periodErr != nil {
		periodInt = 14
	}
	ws, err := ws_indicators.WSSimpleMovingAverageUpgrade(w, r)
	if err != nil {
		panic(err)
	}
	go ws_indicators.WSSimpleMovingAverageWriter(ws, periodInt, priceType)
	go util.CloseWebSocketConnection(ws)
}

func wsExponentialMovingAverage(w http.ResponseWriter, r *http.Request) {
	var period, priceType string
	urlParams := r.URL.Query()
	periodQuery := urlParams["period"]
	priceTypeQuery := urlParams["priceType"]
	if len(periodQuery) == 0 {
		period = "14"
	} else {
		period = periodQuery[0]
	}
	if len(priceTypeQuery) == 0 {
		priceType = constants.CLOSE
	} else {
		priceType = priceTypeQuery[0]
	}

	periodInt, periodErr := strconv.Atoi(period)
	if periodErr != nil {
		periodInt = 14
	}
	ws, err := ws_indicators.WSExponentialMovingAverageUpgrade(w, r)
	if err != nil {
		panic(err)
	}
	go ws_indicators.WSExponentialMovingAverageWriter(ws, periodInt, priceType)
	go util.CloseWebSocketConnection(ws)
}

func wsMovingAverageConvergenceDivergence(w http.ResponseWriter, r *http.Request) {
	var priceType, fastPeriod, slowPeriod, signalPeriod string
	urlParams := r.URL.Query()
	priceTypeQuery := urlParams["priceType"]
	fastPeriodQuery := urlParams["fastPeriod"]
	slowPeriodQuery := urlParams["slowPeriod"]
	signalPeriodQuery := urlParams["signalPeriod"]
	if len(priceTypeQuery) == 0 {
		priceType = constants.CLOSE
	} else {
		priceType = priceTypeQuery[0]
	}
	if len(fastPeriodQuery) == 0 {
		fastPeriod = "12"
	} else {
		fastPeriod = fastPeriodQuery[0]
	}
	if len(slowPeriod) == 0 {
		slowPeriod = "26"
	} else {
		slowPeriod = slowPeriodQuery[0]
	}
	if len(signalPeriodQuery) == 0 {
		signalPeriod = "9"
	} else {
		signalPeriod = signalPeriodQuery[0]
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
	ws, err := ws_indicators.WSMovingAverageConvergenceDivergenceUpgrade(w, r)
	if err != nil {
		panic(err)
	}
	go ws_indicators.WSMovingAverageConvergenceDivergenceWriter(ws, fastPeriodInt, slowPeriodInt, signalPeriodInt, priceType)
	go util.CloseWebSocketConnection(ws)
}

func wsRelativeStrengthIndex(w http.ResponseWriter, r *http.Request) {
	var period, priceType string
	urlParams := r.URL.Query()
	periodQuery := urlParams["period"]
	priceTypeQuery := urlParams["priceType"]
	if len(periodQuery) == 0 {
		period = "14"
	} else {
		period = periodQuery[0]
	}
	if len(priceTypeQuery) == 0 {
		priceType = constants.CLOSE
	} else {
		priceType = priceTypeQuery[0]
	}
	periodInt, periodErr := strconv.Atoi(period)
	if periodErr != nil {
		periodInt = 14
	}

	ws, err := ws_indicators.WSRelativeStrengthIndexUpgrade(w, r)
	if err != nil {
		panic(err)
	}
	go ws_indicators.WSRelativeStrengthIndexWriter(ws, periodInt, priceType)
	go util.CloseWebSocketConnection(ws)
}

func wsStochasticOscillator(w http.ResponseWriter, r *http.Request) {
	var fastKPeriod, slowKPeriod, slowDPeriod string
	urlParams := r.URL.Query()
	fastKPeriodQuery := urlParams["fastKPeriod"]
	slowKPeriodQuery := urlParams["slowKPeriod"]
	slowDPeriodQuery := urlParams["slowDPeriod"]
	if len(fastKPeriodQuery) == 0 {
		fastKPeriod = "14"
	} else {
		fastKPeriod = fastKPeriodQuery[0]
	}
	if len(slowKPeriodQuery) == 0 {
		slowKPeriod = "3"
	} else {
		slowKPeriod = slowKPeriodQuery[0]
	}
	if len(slowDPeriodQuery) == 0 {
		slowDPeriod = "3"
	} else {
		slowDPeriod = slowDPeriodQuery[0]
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
	ws, err := ws_indicators.WSStochasticOscillatorUpgrade(w, r)
	if err != nil {
		panic(err)
	}
	go ws_indicators.WSStochasticOscillatorWriter(ws, fastKPeriodInt, slowKPeriodInt, slowDPeriodInt)
	go util.CloseWebSocketConnection(ws)

}

func wsVolumeWeightedAveragePrice(w http.ResponseWriter, r *http.Request) {
	ws, err := ws_indicators.WSVolumeWeightedAveragePriceUpgrade(w, r)
	if err != nil {
		panic(err)
	}
	go ws_indicators.WSVolumeWeightedAveragePriceWriter(ws)
	go util.CloseWebSocketConnection(ws)
}
