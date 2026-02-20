package cmd

import (
	"fmt"
	"weatherly/api"

	"github.com/spf13/cobra"
)

var lat float64
var lon float64

var weatherCmd = &cobra.Command{
	Use:   "weather",
	Short: "Get current weather data",
	Run: func(cmd *cobra.Command, args []string) {
		response, err := api.FetchDataFromApi(lat, lon)

		if err != nil {
			fmt.Printf("error while fetching data from api: %v\n", err)
			return
		}

		weatherData := api.GetWeatherData(*response)
		weatherDesc := api.GetWeatherDescription(*response)

		fmt.Println(weatherData)
		fmt.Println(weatherDesc)
	},
}

func init() {
	rootCmd.AddCommand(weatherCmd)

	weatherCmd.Flags().Float64Var(&lat, "lat", 0.0, "Latitude")
	weatherCmd.Flags().Float64Var(&lon, "lon", 0.0, "Longitude")
}
