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
