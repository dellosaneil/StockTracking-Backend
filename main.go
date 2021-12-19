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

	test := indicators.StochasticOscillator(s, 5, 3, 3)
	for _, t := range test {
		fmt.Println(t)
	}

	// indicators.StochasticOscillator(s, 5, 3, 3)

}
