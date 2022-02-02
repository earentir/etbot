package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	rgb "github.com/foresthoffman/rgblog"
)

func ParseCommand(bb *BasicBot, msgType, msg, userName string) {
	switch msgType {
	case "PRIVMSG":

		rgb.GPrintf("[%s] %s: %s\n", timeStamp(), userName, msg)

		// parse commands from user message
		cmdMatches := CmdRegex.FindStringSubmatch(msg)
		if nil != cmdMatches {
			cmd := cmdMatches[1]

			switch cmd {
			case "etbdown":
				if inArray(AdminUsers, userName) {
					rgb.CPrintf(
						"[%s] Shutdown command received. Shutting down now...\n",
						timeStamp(),
					)

					bb.Disconnect()
					// return nil
				}

			case "hi":
				atIndex := strings.Index(msg, "@")
				if atIndex > -1 {
					atUser := getAttributedUser(msg, true)
					bb.Say(fmt.Sprintf("earentHey %s, @%s says Hi!", atUser, userName))
				} else {
					bb.Say(fmt.Sprintf("earentHey @%s", userName))
				}

			case "oil":
				bb.Say(oliveoil() + " " + getAttributedUser(msg, true))

			case "bofh":
				bb.Say(JokesBOFH(HTTPGetBody("http://api.esgr.xyz/fun.json/jokes/bofh")) + " " + getAttributedUser(msg, true))

			case "joke":
				fallthrough
			case "yoke":
				bb.Say(JokesJoke(HTTPGetBody("http://api.esgr.xyz/fun.json/jokes/joke")) + " " + getAttributedUser(msg, true))

			case "lurk":
				if msg == "!"+cmd {
					bb.Say(fmt.Sprintf("Thank you for lurking %s, you smart hooman, go have as much fun as possible on your endeavours", userName))
				} else {
					bb.Say(fmt.Sprintf("Thank you for lurking %s, you smart hooman, go have fun with %s", userName, getCleanMessage(cmd, msg)))
				}

			case "socials":
				fallthrough
			case "discord":
				bb.Say("no")

			case "hype":
				atUser := getAttributedUser(msg, true)
				if atUser != "" {
					bb.Say(fmt.Sprintf("earentFfs %s, dont you think there is better places to spend your money ? Stop wasting it !!!", atUser))
				} else {
					bb.Say("Please @ a user")
				}

			case "exchange":
				parts := strings.Fields(msg)

				switch len(parts) {
				case 1:
					bb.Say(CurrencyConversion(settings.General.Curency.DefaultCurrency, settings.General.Curency.CurrencyTo, 1))
				case 2:
					bb.Say(CurrencyConversion(settings.General.Curency.DefaultCurrency, parts[1], 1))
				case 3:
					_, err := strconv.Atoi(parts[2])
					if err != nil {
						bb.Say(CurrencyConversion(parts[1], parts[2], 1))
					} else {
						mnt, _ := strconv.ParseFloat(parts[2], 64)
						bb.Say(CurrencyConversion(settings.General.Curency.DefaultCurrency, parts[1], mnt))
					}
				case 4:
					mnt, _ := strconv.ParseFloat(parts[3], 64)
					bb.Say(CurrencyConversion(parts[1], parts[2], mnt))
				}

			case "w":
				fallthrough
			case "weather":
				if msg == "!"+cmd {
					bb.Say(getWeather(settings.General.Weather.DefaultCity))
				} else {
					bb.Say(getWeather(getCleanMessage(cmd, msg)) + " " + getAttributedUser(msg, true))
				}

			case "pro": //check if they stream and say pro streamer otherwise pro viewer
				bb.Say("soon")

			case "time":
				bb.Say(timeStamp())

			case "commands":
				bb.Say("Available Commands: hi, so, bofh, joke, oil, weather, hype, lurk, time, fr, exchange")

			case "so":
				if inArray(AdminUsers, userName) || inArray(MODUsers, userName) || inArray(VIPUsers, userName) {
					var souser string

					if msg != "!so" {
						if msg[4:5] == "@" {
							souser = msg[5:]
						} else {
							souser = msg[4:]
						}
					} else {
						bb.Say(fmt.Sprintf("@%s Please use !so @username", userName))
					}

					bb.Say(fmt.Sprintf("Please check out & follow %s @ https://twitch.tv/%s they are amazing.%s", souser, strings.ToLower(souser), getItchIOProfile(souser)))
				}

			case "fr":
				if inArray(AdminUsers, userName) || inArray(MODUsers, userName) || inArray(VIPUsers, userName) {

					if len(getCleanMessage(cmd, msg)) > 1 {

						title := getCleanMessage(cmd, msg)[1:]
						cmdres := exec.Command("gh", "issue", "create", fmt.Sprintf("-t %s from %s", title, userName), "-b \"\" ", "-lchat-bot")

						var errbuf bytes.Buffer
						cmdres.Stderr = &errbuf

						cmdres.Run()
						if len(errbuf.String()) == 0 {
							bb.Say("Feature Request Issue Opened")
						}
						fmt.Println(errbuf.String())
					}
				}

			default:
				// do nothing
				// }
			}
		}
	default:
		// fmt.Printf("no priv: %s", line)
	}
}
