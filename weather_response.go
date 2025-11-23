package main

import (
	"fmt"
	"net/http"
)

type WeatherResponse struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`

	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`

	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`

	Name string `json:"name"`
}

func FetchWeather(city, units, apiKey string) (*WeatherResponse, error) {
	url := fmt.Sprintf(
		"http://api.openweathermap.org/data/2.5/weather?q=%s&units=%s&appid=%s",
		city, units, apiKey,
	)

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	var weatherData WeatherResponse

	return &weatherData, nil
}
