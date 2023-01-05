package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func JokesAPI(msg string) []string {
	jokejson := jokes{}
	var retuarr []string

	err := json.Unmarshal([]byte(msg), &jokejson)
	if err == nil {
		fmt.Print(err)
	}

	if jokejson.JOKE.BOFHline != "" {
		retuarr = append(retuarr, jokejson.JOKE.BOFHline)
	} else {
		retuarr = append(retuarr, jokejson.JOKE.Q)
		if jokejson.JOKE.A != "" {
			retuarr = append(retuarr, jokejson.JOKE.A)
		}
	}
	return retuarr
}

func cmdJoke(bb *BasicBot, userName, cmd, msg string) {
	var (
		cleanmsg string   = getCleanMessage(msg)
		attrUser string   = getAttributedUser(msg, false)
		fields   []string = strings.Fields(msg)
		jokelist JokeList
	)

	loadData("Jokes", &jokelist)

	if isCMD(cmd, msg) {
		if len(jokelist.JokeItems) > 0 {
			rand.Seed(time.Now().UnixNano())
			botSay(bb, jokelist.JokeItems[rand.Intn(len(jokelist.JokeItems))].JokeMessage+" >by "+jokelist.JokeItems[rand.Intn(len(jokelist.JokeItems))].AtributedUser)
		}
	} else {
		switch fields[1] {
		case "add":
			if len(fields) >= 3 {
				if len(fields) > 3 {
					botSay(bb, addJoke(userName, attrUser, cleanmsg))
				}
			}

		case "search":
			for i := 0; i < len(jokelist.JokeItems); i++ {
				if strings.EqualFold(jokelist.JokeItems[i].AtributedUser, attrUser) {
					time := time.Unix(jokelist.JokeItems[i].JokeDate, 0)
					botSay(bb, fmt.Sprintf("%s Said \"%s\" on %v", jokelist.JokeItems[i].AtributedUser, jokelist.JokeItems[i].JokeMessage, time.UTC()))
				} else {
					if strings.Contains(jokelist.JokeItems[i].JokeMessage, cleanmsg) && cleanmsg != "" {
						time := time.Unix(jokelist.JokeItems[i].JokeDate, 0)
						botSay(bb, fmt.Sprintf("%s Said \"%s\" on %v", jokelist.JokeItems[i].AtributedUser, jokelist.JokeItems[i].JokeMessage, time.UTC()))
					}
				}
			}

		case "help":
			msgOut := "!joke add @user joke | !joke search @user | !joke search string"
			botSay(bb, msgOut)
		}
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
