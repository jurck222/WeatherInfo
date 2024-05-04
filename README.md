### Weather info
Simple app that shows weather by hour for selected city written in GO with the help of weather api and Cobra library. To use the app get yourself an api key from https://www.weatherapi.com/ (its free) and store it inside an environmental variable called WEATHER_API_KEY. You might have to restart you PC for it to work.

### Usage
To build and install the cli tool, clone the repo and run the following commands
```
go build
```

```
go install
```
When the tool is installed you can run it from the terminal with:
```
weather [command]
```
To check the current weather for a location use:
```
weather current <your-location>
```
If the location is not specified it will show the weather for London UK.

To check the forecast you can use the `forecast` command and specify your location and number of days with the -c and -d flags:
```
weather forecast -c Ljubljana -d 2
```

By default the location is set to London UK for 3 days ahead.
