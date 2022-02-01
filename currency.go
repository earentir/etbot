package main

import (
	"encoding/json"
	"fmt"
)

func CurrencyConversion(ISOFROM, ISOTO string, amount float64) string {

	buildPair := fmt.Sprintf("%s_%s", ISOFROM, ISOTO)

	var pairValue map[string]interface{}

	if creds.OpenWeatherAPIKey != "" {

		exchangerate := HTTPGetBody(fmt.Sprintf("https://free.currconv.com/api/v7/convert?q=%s&compact=ultra&apiKey=%s", buildPair, creds.CurrencyAPIKey))

		if len(exchangerate) > 0 {
			err := json.Unmarshal([]byte(exchangerate), &pairValue)
			if err != nil {
				fmt.Print(err)
			}

			return fmt.Sprintf("The Exchange Rate of %s to %s is %v", ISOFROM, ISOTO, pairValue[buildPair])
		}

		return "Please add a correct City, Country combination, like  \"Athens, Greece\""
	} else {
		return "Please setup your openweathermap API key @ https://openweathermap.org"
	}

}
