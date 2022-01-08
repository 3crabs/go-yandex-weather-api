package main

import (
	"context"
	"fmt"
	weather "github.com/3crabs/go-yandex-weather-api/wheather"
	"log"
	"time"
)

func main() {
	yandexWeatherApiKey := "YOUR_YANDEX_WEATHER_API_KEY"
	w, err := weather.GetWeatherWithCache(context.TODO(), yandexWeatherApiKey, 53.3, 83.5, time.Hour)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Сегодня %s\n", w.Fact.GetCondition())
	fmt.Printf("Температура %d°C\n", w.Fact.Temp)
	fmt.Printf("Ощущается как %d°C\n", w.Fact.FeelsLike)
	fmt.Printf("Порывы ветра до %.1f м/с\n", w.Fact.WindGust)
}
