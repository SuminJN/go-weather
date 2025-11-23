package main

import (
	"encoding/json"
	"fmt"
	"io"
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

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP 요청 오류: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API 응답 오류: 상태 코드 %d, 응답 본문: %s", resp.StatusCode, string(bodyBytes))
	}

	var weatherData WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		return nil, fmt.Errorf("JSON 디코딩 오류: %w", err)
	}

	return &weatherData, nil
}
