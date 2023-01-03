package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type gumroadProducts struct {
	Success  bool `json:"success"`
	Products []struct {
		Name                 string        `json:"name"`
		PreviewURL           string        `json:"preview_url"`
		Description          string        `json:"description"`
		CustomizablePrice    bool          `json:"customizable_price"`
		RequireShipping      bool          `json:"require_shipping"`
		CustomReceipt        string        `json:"custom_receipt"`
		CustomPermalink      string        `json:"custom_permalink"`
		SubscriptionDuration string        `json:"subscription_duration"`
		ID                   string        `json:"id"`
		URL                  interface{}   `json:"url"`
		Price                int           `json:"price"`
		Currency             string        `json:"currency"`
		ShortURL             string        `json:"short_url"`
		ThumbnailURL         string        `json:"thumbnail_url"`
		Tags                 []interface{} `json:"tags"`
		FormattedPrice       string        `json:"formatted_price"`
		Published            bool          `json:"published"`
		ShownOnProfile       bool          `json:"shown_on_profile"`
		FileInfo             struct {
		} `json:"file_info"`
		MaxPurchaseCount   int           `json:"max_purchase_count"`
		Deleted            bool          `json:"deleted"`
		CustomFields       []interface{} `json:"custom_fields"`
		CustomSummary      string        `json:"custom_summary"`
		IsTieredMembership bool          `json:"is_tiered_membership"`
		Recurrences        []string      `json:"recurrences"`
		Variants           []struct {
			Title   string `json:"title"`
			Options []struct {
				Name             string `json:"name"`
				PriceDifference  int    `json:"price_difference"`
				IsPayWhatYouWant bool   `json:"is_pay_what_you_want"`
				RecurrencePrices struct {
					Monthly struct {
						PriceCents          int `json:"price_cents"`
						SuggestedPriceCents int `json:"suggested_price_cents"`
					} `json:"monthly"`
					Quarterly struct {
						PriceCents          int `json:"price_cents"`
						SuggestedPriceCents int `json:"suggested_price_cents"`
					} `json:"quarterly"`
					Biannually struct {
						PriceCents          int `json:"price_cents"`
						SuggestedPriceCents int `json:"suggested_price_cents"`
					} `json:"biannually"`
					Yearly struct {
						PriceCents          int `json:"price_cents"`
						SuggestedPriceCents int `json:"suggested_price_cents"`
					} `json:"yearly"`
				} `json:"recurrence_prices"`
				URL interface{} `json:"url"`
			} `json:"options"`
		} `json:"variants"`
		SalesCount    int     `json:"sales_count"`
		SalesUsdCents float64 `json:"sales_usd_cents"`
	} `json:"products"`
}

func GumroadAPI() string {

	var gumroadproducts gumroadProducts

	if creds.Gumroad != "" {
		gumroadjson := HTTPGetBody(fmt.Sprintf("https://api.gumroad.com/v2/products?access_token=%s", creds.Gumroad))
		json.Unmarshal([]byte(gumroadjson), &gumroadproducts)

		var products string

		for _, product := range gumroadproducts.Products {
			products += product.Name + " @ " + product.FormattedPrice + " buy at " + product.ShortURL + " " + strconv.Itoa(product.SalesCount) + " of " + strconv.Itoa(product.MaxPurchaseCount) + " sold |  "
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

func cmdJokeAPI(bb *BasicBot, cmd, msg string) {
	var (
		jokes []string
		jkstr string
	)

	if cmd == "yoke" {
		cmd = "joke"
	}

	jokes = JokesAPI(HTTPGetBody("http://api.esgr.xyz/fun.json/jokes/" + cmd))

	for _, jk := range jokes {
		jkstr = jkstr + jk + " "
	}

	if isAttr(msg) {
		botSay(bb, jkstr+" "+getAttributedUser(msg, true))
	} else {
		botSay(bb, jkstr)
	}
}
