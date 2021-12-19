package indicators

import (
	"github.com/dellosaneil/stocktracking-backend/model"
	"github.com/dellosaneil/stocktracking-backend/util"
)

func StochasticOscillator(prices []model.PriceModel, fastKPeriod int, slowKPeriod int, slowDPeriod int) []model.StochasticOscillator {
	var stochastic []model.StochasticOscillator
	var prepareForKValue []float64
	var kValues []float64
	highs, lows, closing := getPrices(prices)
	for index := fastKPeriod - 1; index < len(prices); index++ {
		highSlice := highs[index-fastKPeriod+1 : index+1]
		lowSlice := lows[index-fastKPeriod+1 : index+1]
		hh := highestHighFastKPeriod(highSlice)
		ll := lowestLowFastKPeriod(lowSlice)
		preK := prepareKValue(hh, ll, closing[index])
		prepareForKValue = append(prepareForKValue, preK)
	}
	for index := slowDPeriod; index < len(prepareForKValue); index++ {
		k := calculateDValue(prepareForKValue[index-slowDPeriod+1 : index+1])
		kValues = append(kValues, k)
	}
	for index := slowDPeriod; index < len(kValues); index++ {
		d := calculateDValue(kValues[index-slowDPeriod+1 : index+1])
		stochastic = append(stochastic, model.StochasticOscillator{kValues[index], d})
	}

	return stochastic
}

func calculateDValue(kValues []float64) float64 {
	var sum float64
	for _, k := range kValues {
		sum += k
	}
	average := sum / float64(len(kValues))
	return util.RoundPrecision(average, 4)
}

func prepareKValue(highestHigh float64, lowestLow float64, closingPrice float64) float64 {
	top := closingPrice - lowestLow
	bottom := highestHigh - lowestLow
	return util.RoundPrecision((top/bottom)*float64(100), 4)

}

func getPrices(prices []model.PriceModel) ([]float64, []float64, []float64) {
	var highs []float64
	var lows []float64
	var close []float64
	for _, price := range prices {
		highs = append(highs, price.High)
		lows = append(lows, price.Low)
		close = append(close, price.Close)
	}
	return highs, lows, close
}

func highestHighFastKPeriod(prices []float64) float64 {
	highest := prices[0]
	for _, price := range prices {
		if highest < price {
			highest = price
		}
	}
	return util.RoundPrecision(highest, 4)
}

func lowestLowFastKPeriod(prices []float64) float64 {
	lowest := prices[0]
	for _, price := range prices {
		if lowest > price {
			lowest = price
		}
	}
	return util.RoundPrecision(lowest, 4)
}
