package api

type MainResponse struct {
	Temp      float64
	FeelsLike float64
	TempMin   float64
	TempMax   float64
	Pressure  int
	Humidity  int
	SeaLevel  int
	GrndLevel int
}

type WeatherDescription struct {
	Main string
	Icon string
}
