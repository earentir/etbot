package main

import (
	"fmt"
	"time"
)

func getDaysOff(country string, days int) string {
	var daysoff DaysOfF
	var outMessage string = ""
	var date string

	if creds.Calendarific != "" {
		daysoffjson := HTTPGetBody(fmt.Sprintf("https://calendarific.com/api/v2/holidays?api_key=%s&country=%s&year=%v", creds.Calendarific, country, time.Time.Year(time.Now())))
		LoadJSONTOStruct([]byte(daysoffjson), &daysoff)

		for _, k := range daysoff.Response.Holidays {

			if len(k.Date.Iso) == 10 {
				date = k.Date.Iso + "T00:00:01.000Z"
			} else {
				date = k.Date.Iso
			}

			daydiff, _ := time.Parse(time.RFC3339, date)

			if int((time.Until(daydiff).Hours()/24)) > 0 && int((time.Until(daydiff).Hours()/24)) <= days {
				outMessage = outMessage + fmt.Sprintf("%s on %v", k.Name, k.Date.Iso) + " | "
			}
		}

		return outMessage
	} else {
		return "Please setup your Calendarific API key @ https://calendarific.com/"
	}
}
