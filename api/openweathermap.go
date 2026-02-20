package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const apiUrl = "https://api.openweathermap.org/data/2.5/weather?lat=%v&lon=%v&appid=%v"

func GetWeatherData(latitude float64, longitude float64) (*WeatherData, error) {
	if latitude <= 0 || longitude <= 0 {
		return nil, errors.New("latitude or longitude is empty or negative")
	}

	formatedUrl := fmt.Sprintf(apiUrl, latitude, longitude, getApiKey())
	fmt.Println(formatedUrl)
	response, responseError := http.Get(formatedUrl)
	weatherData := WeatherData{}

	if responseError != nil {
		return nil, responseError
	}

	if response.StatusCode == http.StatusOK {
		bodyBytes, bodyBytesError := io.ReadAll(response.Body)

		if bodyBytesError != nil {
			return nil, bodyBytesError
		}

		var result WeatherData
		_ = json.Unmarshal(bodyBytes, &result)

		weatherData.Temp = result.Temp
	} else {
		return nil, fmt.Errorf("server responded with status: %v", response.StatusCode)
	}

	return &weatherData, nil
}
