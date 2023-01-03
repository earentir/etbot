package main

import (
	"fmt"
	"strconv"
	"strings"
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

func cmdDaysOff(bb *BasicBot, cmd, userName, msg string) {
	var (
		daysoffStr string = ""
		outMessage string = ""
		country    string = ""
		daysahead  int    = settings.API.Calendar.DaysAhead
	)
	fields := strings.Fields(msg)

	if isCMD(cmd, msg) {
		country = settings.API.Calendar.Country
		daysahead = settings.API.Calendar.DaysAhead
	} else {
		if len(fields[1]) != 2 {
			botSay(bb, "Please use ISO 3166 country format, ex. !daysoff GR")
		} else {
			country = fields[1]
			switch len(fields) {
			case 2:
				daysahead = settings.API.Calendar.DaysAhead
			case 3:
				days, _ := strconv.ParseInt(fields[2], 10, 64)

				if days > 14 {
					days = 14
				}

				daysahead = int(days)
			}
		}
	}

	daysoffStr = getDaysOff(country, daysahead)

	if len(daysoffStr) > 1 {
		outMessage = fmt.Sprintf("In the next %v days these are these days off in %s: %s", daysahead, country, daysoffStr)
		botSay(bb, outMessage)
	} else {
		botSay(bb, fmt.Sprintf("No Days Off found in the next %v days for %s", daysahead, country))
	}
}
