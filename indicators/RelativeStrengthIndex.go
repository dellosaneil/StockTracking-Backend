package indicators

import (
	"math"

	"github.com/dellosaneil/stocktracking-backend/util"
)

func RelativeStrengthIndex(prices []float64, period int) []float64 {
	var rsi []float64
	priceChange := calculatePriceChange(prices)
	averageGains, averageLosses := getAverages(priceChange, period)
	rs := (averageGains[0] / averageLosses[0])
	r := calculateRSI(rs)
	rsi = append(rsi, r)

	for index := 1; index < len(averageGains); index++ {
		rs = averageGains[index] / averageLosses[index]
		rsi = append(rsi, calculateRSI(rs))
	}
	return rsi
}

func calculateRSI(relativeStrength float64) float64 {
	return util.RoundPrecision(float64(100)-(float64(100)/(float64(1)+relativeStrength)), 4)
}

func calculatePriceChange(prices []float64) []float64 {
	var priceChange []float64
	previousPrice := prices[0]
	for index := 1; index < len(prices); index++ {
		change := prices[index] - previousPrice
		priceChange = append(priceChange, change)
		previousPrice = prices[index]
	}
	return priceChange
}

func getAverages(priceChange []float64, period int) ([]float64, []float64) {
	var gains []float64
	var losses []float64
	var averageGains []float64
	var averageLosses []float64
	for index := 0; index < period; index++ {
		if priceChange[index] > 0 {
			gains = append(gains, priceChange[index])
		} else {
			losses = append(losses, math.Abs(priceChange[index]))
		}
	}
	firstGainAverage := util.Sum(gains) / float64(period)
	firstLossAverage := util.Sum(losses) / float64(period)
	averageGains = append(averageGains, firstGainAverage)
	averageLosses = append(averageLosses, firstLossAverage)
	for index := period; index < len(priceChange); index++ {
		previousAverageGain := averageGains[index-period]
		previousAverageLoss := averageLosses[index-period]
		currentPrice := priceChange[index]
		aGain := previousAverageGain * float64(period-1)
		aLoss := previousAverageLoss * float64(period-1)
		if currentPrice > 0 {
			aGain += currentPrice
		} else {
			aLoss += math.Abs(currentPrice)
		}
		newAverageGain := aGain / float64(period)
		newAverageLoss := math.Abs(aLoss) / float64(period)
		averageGains = append(averageGains, newAverageGain)
		averageLosses = append(averageLosses, newAverageLoss)
	}
	return averageGains, averageLosses
}
