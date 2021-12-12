package indicators

import (
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
