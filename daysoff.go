package main

import (
	"fmt"
	"time"
)

func getDaysOff(country, state string) string {
	var daysoff DaysOfF
	// var outMessage string = ""

	if creds.Calendarific != "" {
		daysoffjson := HTTPGetBody(fmt.Sprintf("https://calendarific.com/api/v2/holidays?api_key=%s&country=%s&year=%v", creds.Calendarific, country, time.Time.Year(time.Now())))
		LoadJSONTOStruct([]byte(daysoffjson), &daysoff)

		for _, k := range daysoff.Response.Holidays {
			fmt.Println(k.Name, "\n", k.Date.Iso, "\n", k.Locations, "\n")
		}

		return "Please add a correct City, Country combination, like  \"Athens, Greece\""
	} else {
		return "Please setup your OpenWeather API key @ https://openweathermap.org"
	}
}
