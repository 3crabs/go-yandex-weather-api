package main

import (
	"fmt"
	weather "github.com/3crabs/go-yandex-weather-api"
)

func main() {
	yandexWeatherApiKey := "YOUR_YANDEX_WEATHER_API_KEY"
	w, _ := weather.GetWeather(yandexWeatherApiKey, 53.366, 83.5)
	fmt.Println(w)
}
