package main

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
