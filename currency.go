package main

import (
	"encoding/json"
	"fmt"
)

func CurrencyConversion(ISOFROM, ISOTO string, amount float64) string {

	buildPair := fmt.Sprintf("%s_%s", ISOFROM, ISOTO)

	var pairValue map[string]float64

	if creds.OpenWeatherAPIKey != "" {

		exchangerate := HTTPGetBody(fmt.Sprintf("https://free.currconv.com/api/v7/convert?q=%s&compact=ultra&apiKey=%s", buildPair, creds.CurrencyAPIKey))
		if len(exchangerate) > 0 {
			err := json.Unmarshal([]byte(exchangerate), &pairValue)
			if err != nil {
				fmt.Print(err)
			}

			return fmt.Sprintf("The Exchange Rate of %v %s to %s is %s", fmt.Sprintf("%.2f", amount), ISOFROM, ISOTO, fmt.Sprintf("%.2f", pairValue[buildPair]*amount))
		}

		return "Please add the ISO codes of FIAT currencies"
	} else {
		return "Please setup your Currency Conversion API key @ https://currconv.com"
	}

}
