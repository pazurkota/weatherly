package api

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
