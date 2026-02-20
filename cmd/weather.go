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
		weather, err := api.GetWeatherData(lat, lon)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(weather.Temp)
		}
	},
}

func init() {
	rootCmd.AddCommand(weatherCmd)

	weatherCmd.Flags().Float64Var(&lat, "lat", 0.0, "Latitude")
	weatherCmd.Flags().Float64Var(&lon, "lon", 0.0, "Longitude")
}
