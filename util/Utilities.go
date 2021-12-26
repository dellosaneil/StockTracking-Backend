package util

import (
	"math"

	"github.com/dellosaneil/stocktracking-backend/constants"
	"github.com/dellosaneil/stocktracking-backend/model"
	"github.com/gorilla/websocket"
)

func RoundPrecision(floatNumber float64, precision int) float64 {
	return math.Round(floatNumber*math.Pow10(precision)) / math.Pow10(precision)
}

func Sum(array []float64) float64 {
	result := float64(0)
	for _, v := range array {
		result += v
	}
	return result
}

func PriceType(prices []model.PriceModel, pricetype string) []float64 {
	var price []float64
	switch pricetype {
	case constants.OPEN:
		for _, p := range prices {
			price = append(price, p.Open)
		}
	case constants.CLOSE:
		for _, p := range prices {
			price = append(price, p.Close)
		}
	case constants.LOW:
		for _, p := range prices {
			price = append(price, p.Low)
		}
	case constants.HIGH:
		for _, p := range prices {
			price = append(price, p.High)
		}
	}
	return price
}

func CloseWebSocketConnection(conn *websocket.Conn) {
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			conn.Close()
			break
		}
	}
}
