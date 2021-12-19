package model

type MACD struct {
	Value      float64
	SignalLine float64
	Histogram  float64
}

type StochasticOscillator struct {
	SlowK float64
	SlowD float64
}
