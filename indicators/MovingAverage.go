package indicators

import (
	"github.com/dellosaneil/stocktracking-backend/model"
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
	return (float64(total) / float64(period))
}

func ExponentialMovingAverage(prices []float64, period int) []float64 {
	var ema []float64
	k := 2.0 / (float64(period) + 1.0)
	slicedArray := prices[0:period]
	previousEma := SimpleMovingAverage(slicedArray, period)[0]
	ema = append(ema, previousEma)
	for index := period; index < len(prices); index++ {
		previousEma = util.RoundPrecision((float64(k)*(prices[index]-previousEma))+previousEma, 4)
		ema = append(ema, previousEma)
	}
	return ema
}

func MovingAverageConvergenceDivergence(prices []float64, fastPeriod int, slowPeriod int, signalPeriod int) []model.MACD {
	fastSlowGap := slowPeriod - fastPeriod
	fastPeriodEma := ExponentialMovingAverage(prices, fastPeriod)
	slowPeriodEma := ExponentialMovingAverage(prices, slowPeriod)
	var macd []model.MACD
	var macdValues []float64
	var histogramValues []float64
	for index := 0; index < len(slowPeriodEma); index++ {
		s := util.RoundPrecision(fastPeriodEma[index+fastSlowGap]-slowPeriodEma[index], 4)
		macdValues = append(macdValues, s)
	}
	signalValues := ExponentialMovingAverage(macdValues, signalPeriod)
	for index := 0; index < len(signalValues); index++ {
		histogram := util.RoundPrecision(macdValues[index+(signalPeriod-1)]-signalValues[index], 4)
		histogramValues = append(histogramValues, histogram)
	}
	for index := 0; index < len(signalValues); index++ {
		macd = append(macd, model.MACD{
			macdValues[index+(signalPeriod-1)],
			signalValues[index],
			histogramValues[index]})
	}
	return macd
}
