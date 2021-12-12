package util

import "math"

func RoundPrecision(floatNumber float64, precision int) float64 {
	return math.Round(floatNumber*math.Pow10(precision)) / math.Pow10(precision)
}
