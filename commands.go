package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func botSay(bb *BasicBot, msg string) {
	if settings.Servers.BotServers.AllowedToSay {
		bb.Say(msg)
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
		botSay(bb, jkstr+" "+getAttributedUser(msg, true))
	} else {
		botSay(bb, jkstr)
	}
}

func cmdLurk(bb *BasicBot, userName, cmd, msg string) {
	var (
		lurklist []string
	)

	if isCMD(cmd, msg) {
		msgOut := fmt.Sprintf("Thank you for lurking %s, you smart hooman, go have as much fun as possible on your endeavours", userName)
		addLurker(userName, cmd, msg)
		botSay(bb, msgOut)
	} else {
		if strings.Fields(msg)[1] == "list" {
			for i := 0; i < len(settings.Lurklists); i++ {
				lurklist = append(lurklist, settings.Lurklists[i].Lurker)
			}
			botSay(bb, fmt.Sprintf("Current Lurkers %s", lurklist))
		} else {
			addLurker(userName, cmd, msg)
			msgOut := fmt.Sprintf("Thank you for lurking %s, you smart hooman, go have fun with %s", userName, getCleanMessage(cmd, msg))
			botSay(bb, msgOut)
		}
	}
}

func cmdUnlurk(bb *BasicBot, userName string) {
	for i := 0; i < len(settings.Lurklists); i++ {
		if strings.EqualFold(userName, settings.Lurklists[i].Lurker) {
			if settings.Lurklists[i].LurkMessage == "" {
				botSay(bb, fmt.Sprintf("Welcome back @%s", userName))
			} else {
				botSay(bb, fmt.Sprintf("Welcome back @%s, how was your %s", userName, settings.Lurklists[i].LurkMessage))
			}
			removeLurker(userName)
		}
	}
}

func cmdExchange(bb *BasicBot, msg string) {
	fields := strings.Fields(strings.ToUpper(msg))

	var petFound bool = false

	for _, j := range fields {
		if strings.EqualFold(j, settings.Pets[0].Name) {
			petFound = true
		}
	}

	if !petFound {
		switch len(fields) {
		case 1:
			botSay(bb, CurrencyConversion(settings.Curency.DefaultCurrency, settings.Curency.CurrencyTo, 1))

		case 2:
			amount, err := strconv.ParseFloat(fields[1], 64)
			if err != nil {
				msgOut := CurrencyConversion(settings.Curency.DefaultCurrency, fields[0], 1)
				botSay(bb, msgOut)
			} else {
				msgOut := CurrencyConversion(settings.Curency.DefaultCurrency, settings.Curency.CurrencyTo, amount)
				botSay(bb, msgOut)
			}

		case 3:
			amount, err := strconv.ParseFloat(fields[1], 64)
			if err != nil {
				msgOut := CurrencyConversion(fields[1], fields[2], 1)
				botSay(bb, msgOut)
			} else {
				msgOut := CurrencyConversion(settings.Curency.DefaultCurrency, fields[2], amount)
				botSay(bb, msgOut)
			}

		case 4:
			amount, _ := strconv.ParseFloat(fields[1], 64)
			msgOut := CurrencyConversion(fields[2], fields[3], amount)
			botSay(bb, msgOut)
		}
	} else {
		botSay(bb, fmt.Sprintf("%s is Priceless.", settings.Pets[0].Name))
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

			var solinks string = strings.Join(getUserSocials(strings.ToLower(atrUser)), " ")
			if solinks == "" {
				solinks = "https://twitch.tv/" + strings.ToLower(atrUser)
			}

			if twitchChannelData[0].GameName == "" || twitchChannelData[0].Title == "" {
				if getUserData(atrUser).Love != "" {
					lovestr = fmt.Sprintf(" and they love %s", getUserData(atrUser).Love)
				}

				msgOut = fmt.Sprintf("Please check out & follow %s they are amazing. You can find them here: %s%s", getAttributedUser(msg, true), solinks, lovestr)
			} else {

				if getUserData(atrUser).Love != "" {
					lovestr = fmt.Sprintf(" and they love %s", getUserData(atrUser).Love)
				}

				msgOut = fmt.Sprintf("Please check out & follow %s they are amazing. They streamed in %s about \"%s\",  You can find them here: %s%s",
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
		title = getCleanMessage(cmd, msg)[1:]
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
	botSay(bb, fmt.Sprintf("%s", getUserSocials(settings.General.Twitch.Channel)))
}

func cmdProject(bb *BasicBot, cmd, userName, msg string) {
	if isCMD(cmd, msg) {
		botSay(bb, settings.General.Project.Description)
	} else {
		if userName == settings.General.Twitch.Channel {
			settings.General.Project.Description = msg[len(cmd)+1:]
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
				for i := 0; i < len(settings.Users); i++ {
					if settings.Users[i].Name == strings.ToLower(fields[0][1:]) {
						var socexists = false
						for j := 0; j < len(settings.Users[i].Socials); j++ {
							if settings.Users[i].Socials[j].SocNet == fields[1] {
								settings.Users[i].Socials[j].Link = fields[2]
								botSay(bb, fmt.Sprintf("%s's %s profile is now %s", fields[0][1:], fields[1], fields[2]))
								socexists = true
							}
						}

						if !socexists {
							socs.SocNet = fields[1]
							socs.Link = fields[2]
							settings.Users[i].Socials = append(settings.Users[i].Socials, socs)
							botSay(bb, fmt.Sprintf("%s's %s profile is now %s", fields[0][1:], fields[1], fields[2]))
						}
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
		for i := 0; i < len(settings.Users); i++ {
			if strings.EqualFold(settings.Users[i].Name, userName) || strings.EqualFold(settings.Users[i].Nick, userName) {
				settings.Users[i].Love = msg[len(cmd)+2:]
				botSay(bb, fmt.Sprintf("You love %s, thank you for letting me know.", settings.Users[i].Love))
			}
		}
	}
}

func cmdZoe(bb *BasicBot, cmd, userName, msg string) {
	if isCMD(cmd, msg) {
		botSay(bb, fmt.Sprintf("%s | Treat: %v(%v)  Petting Minutes: %v", "!zoe pet or !zoe feed or !zoe name", settings.Pets[0].Feed, settings.Pets[0].FeedLimit, settings.Pets[0].Pet))
	} else {
		cmdFields := strings.Fields(msg)
		if len(cmdFields) > 1 {
			if len(settings.Pets) < 1 {
				var pet Pet
				pet.Feed = 0
				pet.Pet = 0

				settings.Pets = append(settings.Pets, pet)
				botSay(bb, "Pet Registered")
			} else {
				switch cmdFields[1] {
				case "pet":
					if len(cmdFields) == 3 {
						if userName == settings.General.Twitch.Channel {
							rem, err := strconv.Atoi(cmdFields[2])
							if err != nil {
								settings.Pets[0].Pet = 0
								botSay(bb, "Pet Reset")
							} else {
								settings.Pets[0].Pet = settings.Pets[0].Pet - rem
								botSay(bb, fmt.Sprintf("%s needs to be peted for %v minutes", settings.Pets[0].Name, settings.Pets[0].Pet))
							}
						}
					} else {
						settings.Pets[0].Pet++
						botSay(bb, fmt.Sprintf("%s needs to be peted for %v minutes", settings.Pets[0].Name, settings.Pets[0].Pet))
					}
				case "feed":
					if len(cmdFields) == 3 {
						if userName == settings.General.Twitch.Channel {
							rem, err := strconv.Atoi(cmdFields[2])
							if err != nil {
								settings.Pets[0].Feed = 0
								botSay(bb, "Feed Reset")
							} else {
								settings.Pets[0].Feed = settings.Pets[0].Feed - rem
								botSay(bb, fmt.Sprintf("%s needs to be given %v treats", settings.Pets[0].Name, settings.Pets[0].Feed))
							}
						}
					} else {
						if settings.Pets[0].FeedLimit < settings.Pets[0].Feed {
							botSay(bb, "No more feed can be added, feed the pet first")
						} else {
							settings.Pets[0].Feed++
							botSay(bb, fmt.Sprintf("%s needs to be given %v treats", settings.Pets[0].Name, settings.Pets[0].Feed))
						}
					}
				case "name":
					if len(cmdFields) == 3 {
						if userName == settings.General.Twitch.Channel {
							settings.Pets[0].Name = cmdFields[2]
						}
					} else {
						botSay(bb, fmt.Sprintf("Pet name is %s", settings.Pets[0].Name))
					}
				case "limit":
					if len(cmdFields) == 3 {
						flimit, err := strconv.Atoi(cmdFields[2])
						if err != nil {
						} else {
							settings.Pets[0].FeedLimit = flimit
						}
					}
				}
			}
		} else {
			botSay(bb, fmt.Sprintf("%s | Treat: %v(%v)  Petting Minutes: %v", "!zoe pet or !zoe feed or !zoe name", settings.Pets[0].Feed, settings.Pets[0].FeedLimit, settings.Pets[0].Pet))
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
			from := settings.Curency.CryptoDefault
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
	var cleanmsg string
	attrUser := getAttributedUser(msg, false)
	fields := strings.Fields(msg)

	if attrUser != "" {
		cleanmsg = msg[len(cmd)+1+1+len(fields[1])+1+len(attrUser)+1:]
	} else {
		cleanmsg = msg[len(cmd)+1+1+len(fields[1])+1:]
	}

	if isCMD(cmd, msg) {
		msgOut := "!quote add @user message | !quote search @user | !quote search string"
		botSay(bb, msgOut)
	} else {

		switch fields[1] {
		case "add":
			if len(fields) > 2 {
				addQuote(userName, attrUser, cleanmsg)
			}

		case "search":
			for i := 0; i < len(settings.Quotes); i++ {
				if strings.EqualFold(settings.Quotes[i].AtributedUser, attrUser) {
					time := time.Unix(settings.Quotes[i].QuoteDate, 0)
					botSay(bb, fmt.Sprintf("%s Said \"%s\" on %v", settings.Quotes[i].AtributedUser, settings.Quotes[i].QuotedMessage, time.UTC()))

				} else {
					if strings.Contains(settings.Quotes[i].QuotedMessage, cleanmsg) && cleanmsg != "" {
						time := time.Unix(settings.Quotes[i].QuoteDate, 0)
						botSay(bb, fmt.Sprintf("%s Said \"%s\" on %v", settings.Quotes[i].AtributedUser, settings.Quotes[i].QuotedMessage, time.UTC()))
					}
				}
			}
		}

	}
}

func isUsrCmdAlias(index int, cmd string) bool {
	var found bool = false
	for i := 0; i < len(usercommands[index].Alias); i++ {
		if usercommands[index].Alias[i] == cmd {
			found = true
		}
	}

	return found
}

func isUsrCmd(cmd string) bool {
	var found bool = false
	for i := 0; i < len(usercommands); i++ {
		if usercommands[i].UserCmdName == cmd || isUsrCmdAlias(i, cmd) {
			found = true
		}
	}

	if !found {
		if getUserSocial(cmd) != "" {
			found = true
		}
	}

	return found
}

func usrCmdList() []string {
	var list []string

	for i := 0; i < len(usercommands); i++ {
		list = append(list, usercommands[i].UserCmdName)
	}
	return list
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
		fmt.Printf("usrCMD: %s\n", usercommands[foundCMDIndex].UserCmdName)
		rand.Seed(time.Now().UnixNano())

		switch usercommands[foundCMDIndex].UserCmdType {
		case "punchline":
			outMessage = usercommands[foundCMDIndex].Messages[rand.Intn(len(usercommands[foundCMDIndex].Messages))]
		case "tree":

		case "counter":

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

func getUserSocial(socnet string) string {
	user := getUserData(settings.General.Twitch.Channel)
	var outMessage string = ""
	for i := 0; i < len(user.Socials); i++ {
		if strings.EqualFold(socnet, user.Socials[i].SocNet) {
			outMessage = user.Socials[i].Link
		}
	}

	return outMessage
}
