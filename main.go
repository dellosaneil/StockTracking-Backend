package main

import (
	"fmt"

	"github.com/dellosaneil/stocktracking-backend/api_data/api_get"
	"github.com/dellosaneil/stocktracking-backend/constants"
	"github.com/dellosaneil/stocktracking-backend/indicators"
)

func main() {
	s, err := api_get.GetPriceCall(constants.DAILY, "TSLA")
	if err != nil {
		fmt.Println(err)
	}
	var o []float64
	for _, price := range s {
		o = append(o, price.Open)
	}

	fmt.Println(indicators.SimpleMovingAverage(o, 3))

}
