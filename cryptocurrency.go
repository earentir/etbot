package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func getCCJSON(from, to string) BinanceData {
	var cryptoData BinanceData

	httpData := HTTPGetBody(fmt.Sprintf("https://www.binance.com/api/v3/ticker/24hr?symbol=%s%s", from, to))
	json.Unmarshal([]byte(httpData), &cryptoData)

	return cryptoData
}

func cmdCryptoExchange(bb *BasicBot, cmd, userName, msg string) {
	if isCMD(cmd, msg) {
		botSay(bb, "!crypto [amount] SYMBOL SYMBOL")
	} else {
		fields := strings.Fields(msg)

		var marketData BinanceData
		switch len(fields) {
		case 2:
			from := settings.API.Curency.CryptoDefault
			to := fields[1]
			marketData = getCCJSON(from, to)
			lastprice, _ := strconv.ParseFloat(marketData.LastPrice, 64)

			botSay(bb, fmt.Sprintf("The Exchange Rate of 1 %s to %s is %s", from, to, fmt.Sprintf("%.2f", lastprice)))
		case 3:
			from := fields[1]
			to := fields[2]
			marketData = getCCJSON(from, to)
			lastprice, _ := strconv.ParseFloat(marketData.LastPrice, 64)

			botSay(bb, fmt.Sprintf("The Exchange Rate of 1 %s to %s is %s", from, to, fmt.Sprintf("%.2f", lastprice)))
		case 4:
			from := fields[2]
			to := fields[3]
			marketData = getCCJSON(from, to)
			amount, _ := strconv.Atoi(fields[1])
			lastprice, _ := strconv.ParseFloat(marketData.LastPrice, 64)

			botSay(bb, fmt.Sprintf("The Exchange Rate of %v %s to %s is %s", amount, from, to, fmt.Sprintf("%.2f", float64(amount)*lastprice)))
		}
	}
}
