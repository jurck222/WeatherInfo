### Weather info
Simple app that shows weather by hour for selected city written in go with the help of weather api.
To use the app get yourself an api key from https://www.weatherapi.com/ (its free) and store it inside an environmental variable called WEATHER_API_KEY. You might have to restart you PC for it to work.

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
weather current <your-location>
```
If the location is not specified it will show the weather for London UK.
