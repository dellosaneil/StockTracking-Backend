package indicators

import (
	"github.com/dellosaneil/stocktracking-backend/model"
	"github.com/dellosaneil/stocktracking-backend/util"
)

// currently default to 1 min, can only be intraday
func VolumnWeightedAveragePrice(prices []model.PriceModel) []float64 {
	previousDate := prices[0].Time[0:10]
	var volumnWeightedAveragePrice []float64
	volumeSummation := float64(0)
	priceVolumeSummation := float64(0)
	for _, price := range prices {
		averagePrice := averagePrice(price)
		priceVolume := averagePrice * float64(price.Volume)
		volumeSummation += float64(price.Volume)
		priceVolumeSummation += priceVolume
		if previousDate != price.Time[0:10] {
			volumeSummation = float64(price.Volume)
			priceVolumeSummation = priceVolume
			previousDate = price.Time[0:10]
		}
		volumnWeightedAveragePrice = append(volumnWeightedAveragePrice, util.RoundPrecision(priceVolumeSummation/volumeSummation, 4))
	}
	return volumnWeightedAveragePrice
}

func averagePrice(price model.PriceModel) float64 {
	return (price.Low + price.High + price.Close) / float64(3)
}
