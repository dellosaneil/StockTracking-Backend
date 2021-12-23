package util

import (
	"math"

	"github.com/dellosaneil/stocktracking-backend/model"
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
	case "open":
		for _, p := range prices {
			price = append(price, p.Open)
		}
	case "close":
		for _, p := range prices {
			price = append(price, p.Close)
		}
	case "low":
		for _, p := range prices {
			price = append(price, p.Low)
		}
	case "high":
		for _, p := range prices {
			price = append(price, p.High)
		}
	}
	return price
}
