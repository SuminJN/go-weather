package main

import (
	"fmt"
	"os"
)

func main() {
	// 설정 로드
	config, err := LoadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 도시 이름 결정
	city, err := determineCity(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 날씨 데이터 가져오기
	weatherData, err := FetchWeather(city, config.Units, config.APIKey)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 날씨 정보 출력
	displayWeather(weatherData)
}
