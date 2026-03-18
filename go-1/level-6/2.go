package main

type WeatherCondition int

const (
	Clear WeatherCondition = iota // Ключевое слово iota присваивает каждой константе числовое значение по порядку (0, 1, 2, 3 и т.д.)
	Rain
	HeavyRain
	Snow
)

type WeatherData struct {
	Condition WeatherCondition
	WindSpeed int
}

func GetWeatherMultiplier(weather WeatherData) float64 {
	k := 1.0
	if weather.Condition == Rain {
		k += 0.125
	}
	if weather.Condition == HeavyRain {
		k += 0.2
	}
	if weather.Condition == Snow {
		k += 0.15
	}
	if weather.WindSpeed > 15 {
		k += 0.1
	}
	return k
}
