package main

import (
	"fmt"

	"github.com/dellosaneil/stocktracking-backend/api/api_get"
	"github.com/dellosaneil/stocktracking-backend/constants"
	"github.com/dellosaneil/stocktracking-backend/indicators"
)

func main() {
	s, err := api_get.GetPriceCall("TSLA", constants.INTRADAY)
	if err != nil {
		fmt.Println(err)
	}
	indicators.VolumnWeightedAveragePrice(s)
}
