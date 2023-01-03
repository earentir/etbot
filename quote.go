package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func cmdQuote(bb *BasicBot, cmd, userName, msg string) {
	var (
		cleanmsg  string   = getCleanMessage(msg)
		attrUser  string   = getAttributedUser(msg, false)
		fields    []string = strings.Fields(msg)
		quotelist QuoteList
	)

	loadData("Quotes", &quotelist)

	if isCMD(cmd, msg) {
		if len(quotelist.QuoteItems) > 0 {
			rand.Seed(time.Now().UnixNano())
			botSay(bb, quotelist.QuoteItems[rand.Intn(len(quotelist.QuoteItems))].QuotedMessage+" >by "+quotelist.QuoteItems[rand.Intn(len(quotelist.QuoteItems))].AtributedUser)
		}
	} else {
		switch fields[1] {
		case "add":
			if len(fields) >= 3 {
				if len(fields) > 3 {
					botSay(bb, addQuote(userName, attrUser, cleanmsg))
				}
			}

		case "search":
			for i := 0; i < len(quotelist.QuoteItems); i++ {
				if strings.EqualFold(quotelist.QuoteItems[i].AtributedUser, attrUser) {
					time := time.Unix(quotelist.QuoteItems[i].QuoteDate, 0)
					botSay(bb, fmt.Sprintf("%s Said \"%s\" on %v", quotelist.QuoteItems[i].AtributedUser, quotelist.QuoteItems[i].QuotedMessage, time.UTC()))
				} else {
					if strings.Contains(quotelist.QuoteItems[i].QuotedMessage, cleanmsg) && cleanmsg != "" {
						time := time.Unix(quotelist.QuoteItems[i].QuoteDate, 0)
						botSay(bb, fmt.Sprintf("%s Said \"%s\" on %v", quotelist.QuoteItems[i].AtributedUser, quotelist.QuoteItems[i].QuotedMessage, time.UTC()))
					}
				}
			}
		case "help":
			msgOut := "!quote add @user message | !quote search @user | !quote search string"
			botSay(bb, msgOut)
		}
	}
}
