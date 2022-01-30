package main

import (
	"encoding/json"
	"fmt"
)

type openweathermap struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

func getWeather(city string) string {
	weatherjson := HTTPGetBody(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, readTextFile("owapikey.txt")[0]))

	openweathermapdata := openweathermap{}

	// fmt.Print("we: " + string(weatherjson))

	if len(weatherjson) > 0 {
		err := json.Unmarshal([]byte(weatherjson), &openweathermapdata)
		if err != nil {
			fmt.Print(err)
		}

		// fmt.Println(openweathermapdata.Weather[0].Description)
		return fmt.Sprintf("%v° (%v°) with %s", openweathermapdata.Main.Temp, openweathermapdata.Main.FeelsLike, openweathermapdata.Weather[0].Description)
	}

	return "Please add a correct City, Country combination, like  \"Athens, Greece\""
}
