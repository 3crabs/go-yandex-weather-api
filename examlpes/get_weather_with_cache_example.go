package main

import (
	"fmt"
	weather "github.com/3crabs/go-yandex-weather-api"
)

func main() {
	yandexWeatherApiKey := "YOUR_YANDEX_WEATHER_API_KEY"
	w, _ := weather.GetWeatherWithCache(yandexWeatherApiKey, 53.3, 83.5, 10*60)
	fmt.Println(w)
}
