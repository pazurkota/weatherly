package api

type WeatherData struct {
	Dt         int
	Sunrise    int
	Sunset     int
	Temp       float64
	FeelsLike  float64
	Pressure   int
	Humidity   int
	DewPoint   float64
	Uvi        float64
	Clouds     int
	Visibility int
	WindSpeed  float64
	WindDeg    int
}

type WeatherDescription struct {
	Main string
	Icon string
}
