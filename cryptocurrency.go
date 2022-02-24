package main

import (
	"encoding/json"
	"fmt"
)

func getCCJSON(from, to string) BinanceData {
	var cryptoData BinanceData

	httpData := HTTPGetBody(fmt.Sprintf("https://www.binance.com/api/v3/ticker/24hr?symbol=%s%s", from, to))
	json.Unmarshal([]byte(httpData), &cryptoData)

	return cryptoData
}
