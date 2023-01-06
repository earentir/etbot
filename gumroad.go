package main

import (
	"encoding/json"
	"fmt"
)

func GumroadAPI() string {

	var gumroadproducts gumroadProducts

	if creds.Gumroad != "" {
		gumroadjson := HTTPGetBody(fmt.Sprintf("https://api.gumroad.com/v2/products?access_token=%s", creds.Gumroad))
		json.Unmarshal([]byte(gumroadjson), &gumroadproducts)

		var products string

		for _, product := range gumroadproducts.Products {
			products += fmt.Sprintf("%s @ %s buy at %s %d of %d sold |", product.Name, product.FormattedPrice, product.ShortURL, product.SalesCount, product.MaxPurchaseCount)
		}
		return products
	}
	return "Please setup your Gumroad API key @ https://gumroad.com"
}

func cmdGumroad(bb *BasicBot, cmd, userName, msg string) {
	var gumroadproducts string = GumroadAPI()

	if isAttr(msg) {
		botSay(bb, gumroadproducts)
	} else {
		botSay(bb, gumroadproducts)
	}
}
