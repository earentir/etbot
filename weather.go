package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func getWeather(city string) string {

	if creds.OpenWeatherAPIKey != "" {
		weatherjson := HTTPGetBody(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, creds.OpenWeatherAPIKey))

		openweathermapdata := openweathermap{}

		if strings.HasPrefix(weatherjson, "http/error/") {
			return fmt.Sprintf("Weather Error: %s", weatherjson)
		}

		if len(weatherjson) > 0 {
			err := json.Unmarshal([]byte(weatherjson), &openweathermapdata)
			if err != nil {
				fmt.Print(err)
			}

			// fmt.Println(openweathermapdata.Weather[0].Description)
			return fmt.Sprintf("The temp in %s is %v° (feels like %v°) with %s", city, openweathermapdata.Main.Temp, openweathermapdata.Main.FeelsLike, openweathermapdata.Weather[0].Description)
		}

		return "Please add a correct City, Country combination, like  \"Athens, Greece\""
	} else {
		return "Please setup your OpenWeather API key @ https://openweathermap.org"
	}
}
