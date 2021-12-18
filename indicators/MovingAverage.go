package indicators

import (
	"fmt"

	"github.com/dellosaneil/stocktracking-backend/util"
)

func SimpleMovingAverage(prices []float64, period int) []float64 {
	var movingAverage []float64
	var slicedArray = prices[0:period]
	var average = movingAverageCalculation(slicedArray, period)
	movingAverage = append(movingAverage, util.RoundPrecision(average, 4))
	for i := period; i < len(prices); i++ {
		var previousTotal = average * float64(period)
		previousTotal = previousTotal - slicedArray[0] + prices[i]
		average = previousTotal / float64(period)
		slicedArray = slicedArray[1:period]
		slicedArray = append(slicedArray, prices[i])
		movingAverage = append(movingAverage, util.RoundPrecision(average, 4))
	}
	return movingAverage
}

func movingAverageCalculation(prices []float64, period int) float64 {
	var total float64
	for _, price := range prices {
		total += price
	}
	return util.RoundPrecision((float64(total) / float64(period)), 4)
}

func ExponentialMovingAverage(prices []float64, period int) []float64 {
	var ema []float64
	k := 2.0 / (float64(period) + 1.0)
	slicedArray := prices[0:period]
	previousEma := SimpleMovingAverage(slicedArray, period)[0]
	ema = append(ema, previousEma)
	for index := period; index < len(prices); index++ {
		previousEma = util.RoundPrecision((float64(k)*(prices[index]-previousEma))+previousEma, 4)
		fmt.Println(previousEma)
		ema = append(ema, previousEma)
	}
	return ema
}
