package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var weatherCmd = &cobra.Command{
	Use:   "current",
	Short: "Weather info",
	Long: `Weather info by hour for selected city with the use of weather api.
	Usage: current <city>
	`,
	Run: getWeather,
}

func init() {
	rootCmd.AddCommand(weatherCmd)
}

func getWeather(cmd *cobra.Command, args []string) {
	api_key := os.Getenv("WEATHER_API_KEY")
	q := "London"
	if len(args) >= 1 {
		q = args[0]
	}
	city := url.QueryEscape(q)
	url := fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=1&aqi=no&alerts=no", api_key, city)
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

	location, current, hour := weather.Location, weather.Current, weather.Forecast.ForecastDay[0].Hour

	fmt.Printf("%s, %s: %.0fC, %s\n", location.Name, location.Country, current.TempC, current.Condition.Text)

	for _, hour := range hour {
		date := time.Unix(int64(hour.TimeEpoch), 0)
		if date.Before(time.Now()) {
			continue
		}
		message := fmt.Sprintf("%s - %.0fC, %.0f, %s\n", date.Format("15:04"), hour.TempC, hour.ChanceOfRain, hour.Condition.Text)

		if hour.ChanceOfRain < 40 {
			fmt.Print(message)
		} else {
			color.Red(message)
		}
	}
}
