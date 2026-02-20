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
		weather, weatherErr := api.GetWeatherData(lat, lon)
		description, descErr := api.GetWeatherDescription(lat, lon)

		if weatherErr != nil {
			fmt.Println("Error fetching weather data:", weatherErr)
			return
		}

		if descErr != nil {
			fmt.Println("Error fetching weather description:", descErr)
			return
		}

		fmt.Println(weather)
		fmt.Println(description)
	},
}

func init() {
	rootCmd.AddCommand(weatherCmd)

	weatherCmd.Flags().Float64Var(&lat, "lat", 0.0, "Latitude")
	weatherCmd.Flags().Float64Var(&lon, "lon", 0.0, "Longitude")
}
