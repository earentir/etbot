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

			return fmt.Sprintf("The temp in %s is %v° (feels like %v°) with %s", city, openweathermapdata.Main.Temp, openweathermapdata.Main.FeelsLike, openweathermapdata.Weather[0].Description)
		}

		return "Please add a correct City, Country combination, like  \"Athens, Greece\""
	} else {
		return "Please setup your OpenWeather API key @ https://openweathermap.org"
	}
}

func cmdWeather(bb *BasicBot, cmd, msg string) {
	if isCMD(cmd, msg) {
		botSay(bb, getWeather(settings.API.Weather.DefaultCity))
	} else {
		var city string = getCleanMessage(msg)
		if strings.TrimSpace(city) == "" {
			city = settings.API.Weather.DefaultCity
		}

		if isAttr(msg) {
			msgOut := fmt.Sprintf("%s %s", getWeather(city), getAttributedUser(msg, true))
			botSay(bb, msgOut)
		} else {
			msgOut := getWeather(city)
			botSay(bb, msgOut)
		}
	}
}
