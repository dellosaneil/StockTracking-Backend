package main

import (
	"fmt"

	"github.com/dellosaneil/stocktracking-backend/api/api_get"
	"github.com/dellosaneil/stocktracking-backend/constants"
	"github.com/dellosaneil/stocktracking-backend/indicators"
)

func main() {
	s, err := api_get.GetPriceCall("TSLA", constants.DAILY)
	if err != nil {
		fmt.Println(err)
	}
	var prices []float64
	for _, price := range s {
		prices = append(prices, price.Close)
	}

	indicators.RelativeStrengthIndex(prices, 14)
}
