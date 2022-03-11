package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func botSay(bb *BasicBot, msg string) {
	if settings.Servers.BotServers.SendMessages {
		if len(msg) > 500 {
			quo := len(msg) / 500
			rem := len(msg) % 500

			for i := 0; i < quo; i++ {
				bb.Say(msg[i*500 : (i+1)*500])
			}

			if rem > 0 {
				bb.Say(msg[quo*500:])
			}

		} else {
			bb.Say(msg)
		}
	} else {
		fmt.Println(msg)
	}
}

func cmdHi(bb *BasicBot, userName, cmd, msg string) {
	var msgOut string = ""
	if isAttr(msg) {
		msgOut = fmt.Sprintf("earentHey %s, @%s says Hi!", getAttributedUser(msg, true), userName)
	} else {
		msgOut = fmt.Sprintf("earentHey @%s", userName)
	}
	botSay(bb, msgOut)
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

func cmdLurk(bb *BasicBot, userName, cmd, msg string) {
	var (
		lurkers  []string
		lurklist LurkList
	)

	LoadJSONFileTOStruct("settings/lurkers.json", &lurklist)

	if isCMD(cmd, msg) {
		msgOut := fmt.Sprintf(getCMDOptions("lurk").Msg, userName)
		addLurker(userName, cmd, msg)
		botSay(bb, msgOut)
	} else {
		switch strings.Fields(msg)[1] {
		case "list":
			for i := 0; i < len(lurklist.Lurkers); i++ {
				lurkers = append(lurkers, lurklist.Lurkers[i].Name)
			}
			botSay(bb, fmt.Sprintf("Current Lurkers %s", lurkers))
		case "help":
			botSay(bb, "!lurk | !lurk [optional reasons | !lurk list]")
		default:
			addLurker(userName, cmd, msg)
			msgOut := fmt.Sprintf(getCMDOptions("lurk").Atmsg, userName, getCleanMessage(msg))
			botSay(bb, msgOut)
		}
	}
}

func cmdUnlurk(bb *BasicBot, userName string) {
	var (
		lurklist LurkList
	)

	LoadJSONFileTOStruct("settings/lurkers.json", &lurklist)

	for i := 0; i < len(lurklist.Lurkers); i++ {
		if strings.EqualFold(userName, lurklist.Lurkers[i].Name) {
			if lurklist.Lurkers[i].Message == "" {
				botSay(bb, fmt.Sprintf("Welcome back @%s", userName))
			} else {
				botSay(bb, fmt.Sprintf("Welcome back @%s, how was your %s", userName, lurklist.Lurkers[i].Message))
			}
			removeLurker(userName)
		}
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

func cmdWeather(bb *BasicBot, cmd, msg string) {
	if isCMD(cmd, msg) {
		botSay(bb, getWeather(settings.API.Weather.DefaultCity))
	} else {
		if isAttr(msg) {
			msgOut := getWeather(getCleanMessage(msg)) + " " + getAttributedUser(msg, true)
			botSay(bb, msgOut)
		} else {
			msgOut := getWeather(getCleanMessage(msg))
			botSay(bb, msgOut)
		}
	}
}

func cmdSO(bb *BasicBot, userName, cmd, msg string) {
	var (
		twitchChannelData TwitchChannelData
		lovestr           string = ""
		msgOut            string
	)

	atrUser := getAttributedUser(msg, false)

	//lets try to get a user without @ special SO case, we SO event without an @
	if atrUser == "" {
		fields := strings.Fields(msg)
		if len(fields) > 1 {
			atrUser = fields[1]
		}
	}

	if userExists(atrUser) {
		if isCMD(cmd, msg) {
			msgOut := fmt.Sprintf("@%s Please use !so @username", userName)
			botSay(bb, msgOut)
		} else {
			tu := getTwitchUser(strings.ToLower(atrUser))[0]
			if tu.ID != "" {
				twitchChannelData = getChannelInfo(tu.ID)
			}

			var solinks string = getUserSocials(strings.ToLower(atrUser))
			if solinks == "" {
				solinks = "https://twitch.tv/" + strings.ToLower(atrUser)
			}

			if twitchChannelData[0].GameName == "" || twitchChannelData[0].Title == "" {
				if getUserData(atrUser).Love != "" {
					lovestr = fmt.Sprintf(" and they love %s", getUserData(atrUser).Love)
				}

				msgOut = fmt.Sprintf(getCMDOptions("so").Atmsg, getAttributedUser(msg, true), solinks, lovestr)
			} else {

				if getUserData(atrUser).Love != "" {
					lovestr = fmt.Sprintf(" and they love %s", getUserData(atrUser).Love)
				}

				msgOut = fmt.Sprintf(getCMDOptions("so").Msg,
					getAttributedUser(msg, true),
					twitchChannelData[0].GameName,
					twitchChannelData[0].Title,
					solinks, lovestr)
			}

			botSay(bb, msgOut)
		}
	}
}

func cmdFr(bb *BasicBot, userName, cmd, msg string) {
	var title string
	if !isCMD(cmd, msg) {
		title = getCleanMessage(msg)[1:]
		cmdres := exec.Command("gh", "issue", "create", fmt.Sprintf("-t %s from %s", title, userName), "-b \"\" ", "-lchat-bot")

		var out, errbuf bytes.Buffer
		cmdres.Stderr = &errbuf
		cmdres.Stdout = &out

		cmdres.Run()
		if len(errbuf.String()) == 0 {
			msgOut := fmt.Sprintf("Feature Request Ticket Opened by %s with a title of \"%s\"", userName, title)
			botSay(bb, msgOut)
		}
		fmt.Println(errbuf.String())
	} else {
		cmdFrList(bb, userName, cmd, msg)
	}
}

func cmdFrList(bb *BasicBot, userName, cmd, msg string) {
	cmdres := exec.Command("gh", "issue", "list")

	var errbuf bytes.Buffer
	var frs []string
	cmdres.Stderr = &errbuf

	out, err := cmdres.Output()
	if err == nil {
		msgOut := fmt.Sprintf("Feature Request List, Opened by %s", userName)
		botSay(bb, msgOut)
		outputLines := strings.Split(string(out), "\n")
		for _, j := range outputLines {
			if strings.Contains(j, "chat-bot") {
				if strings.Contains(j, userName) {
					str := regexp.MustCompile("\t").Split(j, -1)
					date := strings.Fields(str[4])
					line := str[0] + " " + strings.ReplaceAll(str[2], "from "+userName, "") + " on " + date[0] + " " + date[1]
					frs = append(frs, line)
				}
			}
		}
		if len(frs) > 0 {
			for _, j := range frs {
				botSay(bb, j)
			}
		} else {
			msgOut := "Please type your feature request after the cmd, ex: !fr add more stuff FFS"
			botSay(bb, msgOut)
		}
	}
}

func cmdList(bb *BasicBot, userName string) {
	msgOut := fmt.Sprintf("Available Commands:  %s", getCMDS(userName))
	botSay(bb, msgOut)
}

func cmdTime(bb *BasicBot, cmd, msg string) {
	var msgOut string = ""

	if isCMD(cmd, msg) {
		msgOut = timeStamp()
	} else {
		msgOut = timeZone(msg[len(cmd)+2:])
	}

	botSay(bb, msgOut)
}

func cmdVersion(bb *BasicBot) {
	botSay(bb, etbver)
}

func cmdSocial(bb *BasicBot, cmd string) {
	botSay(bb, getUserSocials(settings.General.Twitch.Channel))
}

func cmdProject(bb *BasicBot, cmd, userName, msg string) {
	if isCMD(cmd, msg) {
		botSay(bb, settings.General.Project.Description)
	} else {
		if strings.EqualFold(userName, settings.General.Twitch.Channel) {
			settings.General.Project.Description = getCleanMessage(msg)
		}
	}
}

func cmdUPDSoc(bb *BasicBot, cmd, userName, msg string) {
	var socs Social
	if UserLevel(userName).Level <= levelNameTolvl("mod") {
		if isCMD(cmd, msg) {
			botSay(bb, "Set a users social. ex. !updsoc @earentir github https://github.com/earentir")
		} else {
			fields := strings.Fields(msg[len(cmd)+2:])
			if len(fields) == 3 {
				for i := 0; i < len(userlist.Users); i++ {

					if strings.EqualFold(userlist.Users[i].Name, strings.ToLower(fields[0][1:])) {
						fmt.Println(userlist.Users[i].Name, strings.ToLower(fields[0][1:]))
						var socexists = false
						for j := 0; j < len(userlist.Users[i].Socials); j++ {
							if userlist.Users[i].Socials[j].SocNet == fields[1] {
								userlist.Users[i].Socials[j].Link = fields[2]
								botSay(bb, fmt.Sprintf("%s's %s profile is now %s", fields[0][1:], fields[1], fields[2]))
								socexists = true
							}
						}

						if !socexists {
							socs.SocNet = fields[1]
							socs.Link = fields[2]
							userlist.Users[i].Socials = append(userlist.Users[i].Socials, socs)
							botSay(bb, fmt.Sprintf("%s's %s profile is now %s", fields[0][1:], fields[1], fields[2]))
						}

						userfile, _ := json.MarshalIndent(userlist, "", "\t")
						_ = ioutil.WriteFile("settings/etb-users.json", userfile, 0644)
					}
				}
			} else {
				botSay(bb, "Set a users social. ex. !updsoc @earentir github https://github.com/earentir")
			}
		}
	} else {
		botSay(bb, "Only mods can update socials manually, please ask a mod for help.")
	}
}

func cmdILOVE(bb *BasicBot, cmd, userName, msg string) {
	if isCMD(cmd, msg) {
		botSay(bb, "Add what you love and it will be included in your SO, ex. !love coffee")
	} else {
		for i := 0; i < len(userlist.Users); i++ {
			if strings.EqualFold(userlist.Users[i].Name, userName) || strings.EqualFold(userlist.Users[i].Nick, userName) {
				userlist.Users[i].Love = msg[len(cmd)+2:]
				botSay(bb, fmt.Sprintf("You love %s, thank you for letting me know.", userlist.Users[i].Love))
			}
		}
	}
}

func cmdZoe(bb *BasicBot, cmd, userName, msg string) {
	if isCMD(cmd, msg) {
		var petData = ""

		if len(petlist.Pets) > 0 {
			petData = fmt.Sprintf(" | Treat: %v/%v  Petting Minutes: %v", petlist.Pets[0].Feed, petlist.Pets[0].FeedLimit, petlist.Pets[0].Pet)
		}

		botSay(bb, fmt.Sprintf("!zoe pet or !zoe feed or !zoe name %s", petData))
	} else {
		cmdFields := strings.Fields(msg)
		if len(cmdFields) > 1 {
			if len(petlist.Pets) < 1 {
				var pet Pet
				pet.Feed = 0
				pet.Pet = 0

				petlist.Pets = append(petlist.Pets, pet)
				botSay(bb, "Pet Registered")
			} else {
				switch cmdFields[1] {
				case "pet":
					if len(cmdFields) == 3 {
						if userName == settings.General.Twitch.Channel {
							rem, err := strconv.Atoi(cmdFields[2])
							if err != nil {
								petlist.Pets[0].Pet = 0
								botSay(bb, "Pet Reset")
							} else {
								petlist.Pets[0].Pet = petlist.Pets[0].Pet - rem
								botSay(bb, fmt.Sprintf("%s needs to be peted for %v minutes", petlist.Pets[0].Name, petlist.Pets[0].Pet))
							}
						}
					} else {
						petlist.Pets[0].Pet++
						botSay(bb, fmt.Sprintf("%s needs to be peted for %v minutes", petlist.Pets[0].Name, petlist.Pets[0].Pet))
					}
				case "treat":
					fallthrough
				case "feed":
					if len(cmdFields) == 3 {
						if userName == settings.General.Twitch.Channel {
							rem, err := strconv.Atoi(cmdFields[2])
							if err != nil {
								petlist.Pets[0].Feed = 0
								botSay(bb, "Feed Reset")
							} else {
								petlist.Pets[0].Feed = petlist.Pets[0].Feed - rem
								botSay(bb, fmt.Sprintf("%s needs to be given %v treats", petlist.Pets[0].Name, petlist.Pets[0].Feed))
							}
						}
					} else {
						if petlist.Pets[0].FeedLimit < petlist.Pets[0].Feed {
							botSay(bb, "No more feed can be added, feed the pet first")
						} else {
							petlist.Pets[0].Feed++
							botSay(bb, fmt.Sprintf("%s needs to be given %v treats", petlist.Pets[0].Name, petlist.Pets[0].Feed))
						}
					}
				case "name":
					if len(cmdFields) == 3 {
						if userName == settings.General.Twitch.Channel {
							petlist.Pets[0].Name = cmdFields[2]
						}
					} else {
						botSay(bb, fmt.Sprintf("Pet name is %s", petlist.Pets[0].Name))
					}
				case "limit":
					if len(cmdFields) == 3 {
						flimit, err := strconv.Atoi(cmdFields[2])
						if err != nil {
						} else {
							petlist.Pets[0].FeedLimit = flimit
						}
					}
				}
			}
			//save the pets
			petfile, _ := json.MarshalIndent(petlist, "", "\t")
			_ = ioutil.WriteFile("settings/pets.json", petfile, 0644)
		} else {
			botSay(bb, fmt.Sprintf("%s | Treat: %v (%v)  Petting Minutes: %v", "!zoe pet or !zoe feed or !zoe name", petlist.Pets[0].Feed, petlist.Pets[0].FeedLimit, petlist.Pets[0].Pet))
		}
	}
}

func cmdTMDB(bb *BasicBot, cmd, userName, msg string) {
	var searchresults []TMDBSearchResults

	if isCMD(cmd, msg) {
		botSay(bb, "Get Movie & TV Information. ex. !tmdb movie Blade Runner or !tmdb tv Supernatural or use an ID: !tmdb 78 movie or !tmdb 1622 tv")
	} else {
		if strings.Fields(msg)[1] == "movie" || strings.Fields(msg)[1] == "tv" {
			searchresults = tmdbSearch(msg[len(cmd)+2+len(strings.Fields(msg)[1]):]).Results
			if len(searchresults) > 0 {
				for i := 0; i < len(searchresults); i++ {
					if !searchresults[i].Adult {
						if searchresults[i].MediaType == "movie" {
							botSay(bb, fmt.Sprintf("📇 %s | %v | 📅 %s  🎥%s", searchresults[i].Title, searchresults[i].ID, searchresults[i].ReleaseDate, searchresults[i].Overview))
						}
						if searchresults[i].MediaType == "tv" {
							botSay(bb, fmt.Sprintf("📇 %s | %v | 📅 %s  📺%s", searchresults[i].Name, searchresults[i].ID, searchresults[i].FirstAirDate, searchresults[i].Overview))
						}
					} else {
						botSay(bb, "Cant return adult movies")
					}
				}
			} else {
				botSay(bb, fmt.Sprintf("%s Cant find your movie", getAttributedUser(msg, true)))
			}
		} else { //search by ID
			id, err := strconv.Atoi(strings.Fields(msg)[1])
			if err != nil {
				fmt.Println(err)
			} else {
				movieData := tmdbMovie(id)
				botSay(bb, fmt.Sprintf("📇 %s | %v | 📅 %s  🎥%s", movieData.Title, movieData.ID, movieData.ReleaseDate, movieData.Overview))

				tvData := tmdbTV(id)
				botSay(bb, fmt.Sprintf("📇 %s | %v | 📅 %s  📺%s", tvData.Name, tvData.ID, tvData.FirstAirDate, tvData.Overview))
			}
		}
	}
}

func cmdLevel(bb *BasicBot, cmd, userName, msg string) {
	if isAttr(msg) {
		atrUser := getAttributedUser(msg, false)
		botSay(bb, fmt.Sprintf("@%s level is %v as a %s", atrUser, UserLevel(atrUser).Level, UserLevel(atrUser).Name))
	} else {
		botSay(bb, fmt.Sprintf("@%s level is %v as a %s", userName, UserLevel(userName).Level, UserLevel(userName).Name))
	}
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

func cmdUser(bb *BasicBot, cmd, userName, msg string) {
	if isCMD(cmd, msg) {
		botSay(bb, "!user add @username type")
	}

	if UserLevel(userName).Level <= levelNameTolvl("mod") {
		fields := strings.Fields(msg)
		attrUsr := getAttributedUser(msg, false)

		if len(fields) >= 3 && attrUsr != "" {
			switch fields[1] {
			case "add":
				fallthrough
			case "a":
				if len(fields) > 3 {
					if UserLevel(userName).Level < levelNameTolvl(fields[3]) {
						botSay(bb, addUser(attrUsr, fields[3]))
					} else {
						botSay(bb, addUser(attrUsr, lvlToLevelName(UserLevel(userName).Level-2)))
					}
				} else {
					botSay(bb, addUser(attrUsr, lvlToLevelName(10)))
				}
			case "delete":
				fallthrough
			case "del":
				fallthrough
			case "d":
				if strings.EqualFold(userName, settings.General.Twitch.Channel) {
					botSay(bb, delUser(attrUsr))
				}
			}
		}
	}
}

func cmdSaveSettings(bb *BasicBot, cmd, userName, msg string) {
	if userName == settings.General.Twitch.Channel {
		saveSettings()
		botSay(bb, "Settings Saved")
	}
}

func cmdSetting(bb *BasicBot, cmd, userName, msg string) {
	fields := strings.Fields(msg)

	switch fields[1] {
	case "servers":
		fallthrough
	case "server":
		fallthrough
	case "srv":
		switch fields[2] {
		case "web":

		case "bot":

		}
	}

}

func cmdQuote(bb *BasicBot, cmd, userName, msg string) {
	var (
		cleanmsg  string   = getCleanMessage(msg)
		attrUser  string   = getAttributedUser(msg, false)
		fields    []string = strings.Fields(msg)
		quotelist QuoteList
	)

	LoadJSONFileTOStruct("settings/quotes.json", &quotelist)

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

func cmdJoke(bb *BasicBot, userName, cmd, msg string) {
	var (
		cleanmsg string   = getCleanMessage(msg)
		attrUser string   = getAttributedUser(msg, false)
		fields   []string = strings.Fields(msg)
		jokelist JokeList
	)

	LoadJSONFileTOStruct("settings/jokes.json", &jokelist)

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

func usrCmdSay(bb *BasicBot, userName, cmd, msg string) {
	botSay(bb, usrCmd(userName, cmd, msg))
}

func usrCmd(userName, cmd, msg string) string {
	var outMessage string = ""
	var foundCMDIndex int = -1

	for i := 0; i < len(usercommands); i++ {
		if usercommands[i].UserCmdName == cmd || isUsrCmdAlias(i, cmd) {
			foundCMDIndex = i
		}
	}

	if foundCMDIndex == -1 {
		user := getUserData(settings.General.Twitch.Channel).Socials
		for i := 0; i < len(user); i++ {
			if strings.EqualFold(user[i].SocNet, cmd) {
				outMessage = user[i].Link
			}
		}
	}

	//this should be a func
	if foundCMDIndex > -1 {
		rand.Seed(time.Now().UnixNano())

		switch usercommands[foundCMDIndex].UserCmdType {
		case "punchline":
			outMessage = usercommands[foundCMDIndex].Messages[rand.Intn(len(usercommands[foundCMDIndex].Messages))]
		case "tree":

		case "counter":
			outMessage = fmt.Sprintf(usercommands[foundCMDIndex].Messages[0], cmd, usercommands[foundCMDIndex].UserCmdOptions.Counter)
		case "varpunchline":
			outMessage = usercommands[foundCMDIndex].Messages[rand.Intn(len(usercommands[foundCMDIndex].Messages))]

			attrUser := getAttributedUser(msg, false)
			if attrUser != "" {
				outMessage = strings.ReplaceAll(outMessage, "^a", attrUser)
			} else {
				outMessage = strings.ReplaceAll(outMessage, "^a", userName)
			}
			outMessage = strings.ReplaceAll(outMessage, "^u", userName)
		}
	}

	return outMessage
}

func cmdYear(bb *BasicBot, cmd, userName, msg string) {
	currentDay := time.Time.YearDay(time.Now())
	totalDays := time.Time.YearDay(time.Date(time.Time.Year(time.Now()), 12, 31, 00, 00, 00, 0, time.UTC))
	currentPercent := (float64(currentDay) / float64(totalDays)) * 100
	_, week := time.Time.ISOWeek(time.Now())
	botSay(bb, fmt.Sprintf("It is day %v, week %v of %v (%v days), we have used %2.2f%% [%s] of the year till now.", currentDay, week, time.Time.Year(time.Now()), totalDays, currentPercent, progressbar(currentPercent)))
}

func cmdDaysOff(bb *BasicBot, cmd, userName, msg string) {
	var (
		daysoffStr string = ""
		outMessage string = ""
		country    string = ""
		daysahead  int    = settings.API.Calendar.DaysAhead
	)
	fields := strings.Fields(msg)

	if isCMD(cmd, msg) {
		country = settings.API.Calendar.Country
		daysahead = settings.API.Calendar.DaysAhead
	} else {
		if len(fields[1]) != 2 {
			botSay(bb, "Please use ISO 3166 country format, ex. !daysoff GR")
		} else {
			country = fields[1]
			switch len(fields) {
			case 2:
				daysahead = settings.API.Calendar.DaysAhead
			case 3:
				days, _ := strconv.ParseInt(fields[2], 10, 64)

				if days > 14 {
					days = 14
				}

				daysahead = int(days)
			}

		}
	}

	daysoffStr = getDaysOff(country, daysahead)

	if len(daysoffStr) > 1 {
		outMessage = fmt.Sprintf("In the next %v days these are these days off in %s: %s", daysahead, country, daysoffStr)
		botSay(bb, outMessage)
	} else {
		botSay(bb, fmt.Sprintf("No Days Off found in the next %v days for %s", daysahead, country))
	}
}
