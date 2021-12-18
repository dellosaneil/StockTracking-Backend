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
	var o []float64
	for _, price := range s {
		o = append(o, price.Close)
	}

	indicators.ExponentialMovingAverage(o, 14)

}
