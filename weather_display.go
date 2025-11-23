package main

import "fmt"

func displayWeather(weatherData *WeatherResponse) {
	fmt.Println("\n--- 날씨 정보 ---")
	fmt.Printf("도시: %s (위도: %.2f, 경도: %.2f)\n", weatherData.Name, weatherData.Coord.Lat, weatherData.Coord.Lon)

	if len(weatherData.Weather) > 0 {
		fmt.Printf("날씨: %s (%s)\n", weatherData.Weather[0].Main, weatherData.Weather[0].Description)
	}

	fmt.Printf("현재 온도: %.1f°C\n", weatherData.Main.Temp)
	fmt.Printf("체감 온도: %.1f°C\n", weatherData.Main.FeelsLike)
	fmt.Printf("최저/최고 온도: %.1f°C / %.1f°C\n", weatherData.Main.TempMin, weatherData.Main.TempMax)
	fmt.Printf("습도: %d%%\n", weatherData.Main.Humidity)
}
