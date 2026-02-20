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

func GetWeatherData(latitude float64, longitude float64) (*MainResponse, error) {
	if latitude <= 0 || longitude <= 0 {
		return nil, errors.New("latitude or longitude is empty or negative")
	}

	formatedUrl := fmt.Sprintf(apiUrl, latitude, longitude, config.GetApiKey())
	response, responseError := http.Get(formatedUrl)
	weatherData := MainResponse{}

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

		weatherData.Temp = result.Main.Temp
		weatherData.FeelsLike = result.Main.FeelsLike
		weatherData.TempMin = result.Main.TempMin
		weatherData.TempMax = result.Main.TempMax
		weatherData.Pressure = result.Main.Pressure
		weatherData.Humidity = result.Main.Humidity
		weatherData.SeaLevel = result.Main.SeaLevel
		weatherData.GrndLevel = result.Main.GrndLevel

	} else {
		return nil, fmt.Errorf("server responded with status: %v", response.StatusCode)
	}

	return &weatherData, nil
}

func GetWeatherDescription(latitude float64, longitude float64) (*WeatherDescription, error) {
	if latitude <= 0 || longitude <= 0 {
		return nil, errors.New("latitude or longitude is empty or negative")
	}

	formatedUrl := fmt.Sprintf(apiUrl, latitude, longitude, config.GetApiKey())
	response, responseError := http.Get(formatedUrl)
	weatherData := WeatherDescription{}

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

		weatherData.Main = result.Weather[0].Description
		weatherData.Icon = result.Weather[0].Icon

	} else {
		return nil, fmt.Errorf("server responded with status: %v", response.StatusCode)
	}

	return &weatherData, nil
}
