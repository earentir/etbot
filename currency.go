package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func CurrencyConversion(ISOFROM, ISOTO string, amount float64) string {

	buildPair := fmt.Sprintf("%s_%s", ISOFROM, ISOTO)

	var pairValue map[string]float64

	if creds.CurrencyAPIKey != "" {

		exchangerate := HTTPGetBody(fmt.Sprintf("https://free.currconv.com/api/v7/convert?q=%s&compact=ultra&apiKey=%s", buildPair, creds.CurrencyAPIKey))
		if strings.HasPrefix(exchangerate, "http/error/") {
			return fmt.Sprintf("Currency Exchange Error: %s", exchangerate)
		}

		if len(exchangerate) > 0 && exchangerate != "{}" {
			err := json.Unmarshal([]byte(exchangerate), &pairValue)
			if err != nil {
				fmt.Println("Currency Exchange Error: ", err)
			}

			return fmt.Sprintf("The Exchange Rate of %v %s to %s is %s", fmt.Sprintf("%.2f", amount), ISOFROM, ISOTO, fmt.Sprintf("%.2f", pairValue[buildPair]*amount))
		}

		return "Please add the ISO codes of FIAT currencies"
	} else {
		return "Please setup your Currency Conversion API key @ https://currconv.com"
	}

}

func cmdExchange(bb *BasicBot, msg string) {
	var fields []string = strings.Fields(strings.ToUpper(msg))
	var petFound bool = false

	for _, j := range fields {
		if strings.EqualFold(j, petlist.Pets[0].Name) {
			petFound = true
		}
	}

	if !petFound {
		switch len(fields) {
		case 1:
			botSay(bb, CurrencyConversion(settings.API.Curency.DefaultCurrency, settings.API.Curency.CurrencyTo, 1))

		case 2:
			amount, err := strconv.ParseFloat(fields[1], 64)
			if err != nil {
				msgOut := CurrencyConversion(settings.API.Curency.DefaultCurrency, fields[0], 1)
				botSay(bb, msgOut)
			} else {
				msgOut := CurrencyConversion(settings.API.Curency.DefaultCurrency, settings.API.Curency.CurrencyTo, amount)
				botSay(bb, msgOut)
			}

		case 3:
			amount, err := strconv.ParseFloat(fields[1], 64)
			if err != nil {
				msgOut := CurrencyConversion(fields[1], fields[2], 1)
				botSay(bb, msgOut)
			} else {
				msgOut := CurrencyConversion(settings.API.Curency.DefaultCurrency, fields[2], amount)
				botSay(bb, msgOut)
			}

		case 4:
			amount, _ := strconv.ParseFloat(fields[1], 64)
			msgOut := CurrencyConversion(fields[2], fields[3], amount)
			botSay(bb, msgOut)
		}
	} else {
		botSay(bb, fmt.Sprintf("%s is Priceless.", petlist.Pets[0].Name))
	}
}
