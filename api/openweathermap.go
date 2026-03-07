package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"weatherly/config"
)

const weatherApiUrl = "https://api.openweathermap.org/data/2.5/weather?lat=%v&lon=%v&appid=%v"
const geolocationApiUrl = "http://api.openweathermap.org/geo/1.0/direct?q=%v&limit=1&appid=%v"

func FetchDataFromApi(latitude float64, longitude float64) (*WeatherResponse, error) {
	if latitude <= 0 || longitude <= 0 {
		return nil, errors.New("latitude or longitude is empty or negative")
	}

	formatedUrl := fmt.Sprintf(weatherApiUrl, latitude, longitude, config.GetApiKey())
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

func GetLatAndLonByCityName(cityName string) (*GeolocationResponse, error) {
	if cityName == "" {
		return nil, errors.New("city name can't be null")
	}

	formattedUrl := fmt.Sprintf(geolocationApiUrl, cityName, config.GetApiKey())
	response, err := http.Get(formattedUrl)

	if err != nil {
		return nil, err
	}

	if response.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(response.Body)

		if err != nil {
			return nil, err
		}

		var result GeolocationResponse
		_ = json.Unmarshal(bodyBytes, &result)

		return &result, nil
	}

	return nil, fmt.Errorf("server responded with status: %v", response.StatusCode)
}
