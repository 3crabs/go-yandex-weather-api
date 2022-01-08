package main

import (
	"context"
	"fmt"
	weather "github.com/3crabs/go-yandex-weather-api/wheather"
	"log"
)

func main() {
	yandexWeatherApiKey := "YOUR_YANDEX_WEATHER_API_KEY"
	w, err := weather.GetWeather(context.TODO(), yandexWeatherApiKey, 53.3, 83.5)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(w)
}
