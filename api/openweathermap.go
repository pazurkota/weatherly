package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"weatherly/config"
)

const apiUrl = "https://api.openweathermap.org/data/2.5/weather?lat=%v&lon=%v&appid=%v"

func GetWeatherData(response WeatherResponse) *MainResponse {
	var body MainResponse

	body.Temp = response.Main.Temp
	body.FeelsLike = response.Main.FeelsLike
	body.TempMin = response.Main.TempMin
	body.TempMax = response.Main.TempMax
	body.Pressure = response.Main.Pressure
	body.Humidity = response.Main.Humidity
	body.SeaLevel = response.Main.SeaLevel
	body.GrndLevel = response.Main.GrndLevel

	return &body
}

func GetWeatherDescription(response WeatherResponse) *WeatherDescription {
	var body WeatherDescription

	body.Main = response.Weather[0].Main
	body.Icon = response.Weather[0].Icon

	return &body
}

func FetchDataFromApi(latitude float64, longitude float64) (*WeatherResponse, error) {
	if latitude <= 0 || longitude <= 0 {
		return nil, errors.New("latitude or longitude is empty or negative")
	}

	formatedUrl := fmt.Sprintf(apiUrl, latitude, longitude, config.GetApiKey())
	response, responseError := http.Get(formatedUrl)

	if responseError != nil {
		return nil, responseError
	}

	if response.StatusCode == http.StatusOK {
		bodyBytes, bodyBytesError := io.ReadAll(response.Body)

		if bodyBytesError != nil {
			return nil, bodyBytesError
		}

		var result WeatherResponse
		_ = json.Unmarshal(bodyBytes, &result)

		return &result, nil
	}

	return nil, fmt.Errorf("server responded with status: %v", response.StatusCode)
}
