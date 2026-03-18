package main

const (
	pricePerKm     = 10.0
	pricePerMinute = 2.0
)

type TripParameters struct {
	Distance float64
	Duration float64
}

func CalculateBasePrice(tp TripParameters) float64 {
	return tp.Distance*pricePerKm + tp.Duration*pricePerMinute
}
