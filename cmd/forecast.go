package cmd

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"encoding/json"
	"io"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// forecastCmd represents the forecast command
var forecastCmd = &cobra.Command{
	Use:   "forecast",
	Short: "Shows forecast for selected city",
	Long: `shows forecast for selected city using the weather api for a few days ahead.
	Usage: forecast -c <city> -d <days>`,
	Run: getForecast,
}

func getForecast(cmd *cobra.Command, args []string) {
	api_key := os.Getenv("WEATHER_API_KEY")
	q, _ := cmd.Flags().GetString("city")
	d, _ := cmd.Flags().GetInt("days")
	if len(args) >= 1 {
		q = args[0]
	}
	city := url.QueryEscape(q)
	url := fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=%d&aqi=no&alerts=no", api_key, city, d)
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Data fetching failed with " + res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}

	location, current, days := weather.Location, weather.Current, weather.Forecast.ForecastDay

	fmt.Printf("%s, %s: %.0fC, %s\n", location.Name, location.Country, current.TempC, current.Condition.Text)
	for _, day := range days {
		date := time.Unix(int64(day.DateEpoch), 0)

		formattedDate := date.Format("Monday 2 Jan 2006")

		color.Green("\nForecast for %s:\n", formattedDate)
		hour := day.Hour
		for _, hour := range hour {
			date := time.Unix(int64(hour.TimeEpoch), 0)
			if date.Before(time.Now()) {
				continue
			}
			message := fmt.Sprintf("\t%s - %.0fC, %.0f, %s\n", date.Format("15:04"), hour.TempC, hour.ChanceOfRain, hour.Condition.Text)

			if hour.ChanceOfRain < 40 {
				fmt.Print(message)
			} else {
				color.Red(message)
			}
		}
	}
}

func init() {
	rootCmd.AddCommand(forecastCmd)

	forecastCmd.Flags().StringP("city", "c", "London", "City")
	forecastCmd.Flags().IntP("days", "d", 3, "Days")
}
