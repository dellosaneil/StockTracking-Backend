package main

import (
	"fmt"

	"github.com/dellosaneil/stocktracking-backend/api_data/api_get"
	"github.com/dellosaneil/stocktracking-backend/constants"
)

func main() {
	s, err := api_get.GetPriceCall(constants.DAILY, "TSLA")
	if err != nil {
		fmt.Println(err)
	}
	for _, q := range s {
		fmt.Println(q.Open)
	}

}
