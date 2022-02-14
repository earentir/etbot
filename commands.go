package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func botSay(bb *BasicBot, msg string) {
	bb.Say(msg)
}

func cmdHi(bb *BasicBot, userName, cmd, msg string) {
	if isAttr(msg) {
		msgOut := fmt.Sprintf("earentHey %s, @%s says Hi!", getAttributedUser(msg, true), userName)
		botSay(bb, msgOut)
	} else {
		msgOut := fmt.Sprintf("earentHey @%s", userName)
		botSay(bb, msgOut)
	}
}

func cmdNo(bb *BasicBot) {
	botSay(bb, "No")
}

func cmdSoon(bb *BasicBot) {
	botSay(bb, "TBD")
}

func cmdVulgar(bb *BasicBot) {
	botSay(bb, "Fuck Off")
}

func cmdJokeAPI(bb *BasicBot, cmd, msg string) {
	var jokes []string
	var jkstr string

	if cmd == "yoke" {
		cmd = "joke"
	}

	jokes = JokesAPI(HTTPGetBody("http://api.esgr.xyz/fun.json/jokes/" + cmd))

	for _, jk := range jokes {
		jkstr = jkstr + jk + " "
	}

	if isAttr(msg) {
		botSay(bb, jkstr)
	} else {
		botSay(bb, jkstr)
	}
}

func cmdBan(bb *BasicBot, userName, cmd, msg string) {
	if isAttr(msg) {
		msgOut := fmt.Sprintf("%s is %sned.", getAttributedUser(msg, true), cmd)
		botSay(bb, msgOut)
	} else {
		msgOut := fmt.Sprintf("%s is %sned.", userName, cmd)
		botSay(bb, msgOut)
	}
}

func cmdLurk(bb *BasicBot, userName, cmd, msg string) {
	if isCMD(cmd, msg) {
		msgOut := fmt.Sprintf("Thank you for lurking %s, you smart hooman, go have as much fun as possible on your endeavours", userName)
		botSay(bb, msgOut)
	} else {
		msgOut := fmt.Sprintf("Thank you for lurking %s, you smart hooman, go have fun with %s", userName, getCleanMessage(cmd, msg))
		botSay(bb, msgOut)
	}
}

func cmdHype(bb *BasicBot, msg string) {
	if isAttr(msg) {
		msgOut := fmt.Sprintf(jokesJSON("sub_insults"), getAttributedUser(msg, true))
		botSay(bb, msgOut)
	} else {
		botSay(bb, "Please @ a user")
	}
}

func cmdExchange(bb *BasicBot, msg string) {
	parts := strings.Fields(strings.ToUpper(msg))
	switch len(parts) {
	case 1:
		botSay(bb, CurrencyConversion(settings.Curency.DefaultCurrency, settings.Curency.CurrencyTo, 1))

	case 2:
		amount, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			msgOut := CurrencyConversion(settings.Curency.DefaultCurrency, parts[0], 1)
			botSay(bb, msgOut)
		} else {
			msgOut := CurrencyConversion(settings.Curency.DefaultCurrency, settings.Curency.CurrencyTo, amount)
			botSay(bb, msgOut)
		}

	case 3:
		amount, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			msgOut := CurrencyConversion(parts[1], parts[2], 1)
			botSay(bb, msgOut)
		} else {
			msgOut := CurrencyConversion(settings.Curency.DefaultCurrency, parts[2], amount)
			botSay(bb, msgOut)
		}

	case 4:
		amount, _ := strconv.ParseFloat(parts[1], 64)
		msgOut := CurrencyConversion(parts[2], parts[3], amount)
		botSay(bb, msgOut)
	}
}

func cmdWeather(bb *BasicBot, cmd, msg string) {
	if isCMD(cmd, msg) {
		botSay(bb, getWeather(settings.Weather.DefaultCity))
	} else {
		if isAttr(msg) {
			msgOut := getWeather(getCleanMessage(cmd, msg)) + " " + getAttributedUser(msg, true)
			botSay(bb, msgOut)
		} else {
			msgOut := getWeather(getCleanMessage(cmd, msg))
			botSay(bb, msgOut)
		}
	}
}

func cmdSO(bb *BasicBot, userName, cmd, msg string) {
	fmt.Println(getAttributedUser(msg, true))
	if userExists(getAttributedUser(msg, false)) {
		if isCMD(cmd, msg) {
			msgOut := fmt.Sprintf("@%s Please use !so @username", userName)
			botSay(bb, msgOut)
		} else {
			msgOut := fmt.Sprintf("Please check out & follow %s @ https://twitch.tv/%s they are amazing.%s", getAttributedUser(msg, true), strings.ToLower(getAttributedUser(msg, false)), getItchIOProfile(getAttributedUser(msg, false)))
			botSay(bb, msgOut)
		}
	}
}

func cmdFr(bb *BasicBot, userName, cmd, msg string) {
	var title string
	if !isCMD(cmd, msg) {
		title = getCleanMessage(cmd, msg)[1:]
		cmdres := exec.Command("gh", "issue", "create", fmt.Sprintf("-t %s from %s", title, userName), "-b \"\" ", "-lchat-bot")

		var errbuf bytes.Buffer
		cmdres.Stderr = &errbuf

		cmdres.Run()
		if len(errbuf.String()) == 0 {
			msgOut := fmt.Sprintf("Feature Request Ticket Opened by %s with a title of \"%s\"", userName, title)
			botSay(bb, msgOut)
		}
		fmt.Println(errbuf.String())
	} else {
		msgOut := "Please type your feature request after the cmd, ex: !fr add more stuff FFS"
		botSay(bb, msgOut)
	}
}

func cmdList(bb *BasicBot, userName string) {
	msgOut := fmt.Sprintf("Available Commands:  %s", getCMDS(userName))
	botSay(bb, msgOut)
}

func cmdTime(bb *BasicBot, cmd, msg string) {
	msgOut := timeStamp()
	botSay(bb, msgOut)
}

func cmdVersion(bb *BasicBot) {
	botSay(bb, etbver)
}

func cmdMic(bb *BasicBot) {
	msgOut := "earentFfs Check your mic moron @earentir"
	botSay(bb, msgOut)
}

func cmdJoke(bb *BasicBot, cmd string) {
	msgOut := jokesJSON(cmd)
	botSay(bb, msgOut)
}
