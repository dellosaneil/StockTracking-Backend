package util

import "math"

func RoundPrecision(floatNumber float64, precision int) float64 {
	return math.Round(floatNumber*math.Pow10(precision)) / math.Pow10(precision)
}

func Sum(array []float64) float64 {
	result := float64(0)
	for _, v := range array {
		result += v
	}
	return result
}
