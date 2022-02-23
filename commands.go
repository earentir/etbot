package main

import (
	"bytes"
	"fmt"
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
	if isAttr(msg) {
		msgOut := fmt.Sprintf("earentHey %s, @%s says Hi!", getAttributedUser(msg, true), userName)
		botSay(bb, msgOut)
	} else {
		msgOut := fmt.Sprintf("earentHey @%s", userName)
		botSay(bb, msgOut)
	}
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
		botSay(bb, jkstr+" "+getAttributedUser(msg, true))
	} else {
		botSay(bb, jkstr)
	}
}

func cmdJoke(bb *BasicBot, cmd, msg string) {
	var msgOut string
	atrUser := getAttributedUser(msg, true)

	if isAttr(msg) {
		msgOut = jokesJSON(cmd) + " " + atrUser
	} else {
		msgOut = jokesJSON(cmd)
	}

	botSay(bb, msgOut)
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
	var lurker LurkerList

	lurker.Lurker = userName
	lurker.LurkedOn = int(time.Now().Unix())
	settings.Lurklists = append(settings.Lurklists, lurker)

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
	var twitchChannelData TwitchChannelData
	var msgOut string

	atrUser := getAttributedUser(msg, false)

	//lets try to get a user without @
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
				msgOut = fmt.Sprintf("Please check out & follow %s they are amazing. You can find them here: %s", getAttributedUser(msg, true), solinks)
			} else {
				msgOut = fmt.Sprintf("Please check out & follow %s they are amazing. They streamed in %s about \"%s\",  You can find them here: %s",
					getAttributedUser(msg, true),
					twitchChannelData[0].GameName,
					twitchChannelData[0].Title,
					solinks)
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

func cmdSocial(bb *BasicBot, cmd string) {
	for _, j := range settings.Users {
		if j.Type == "root" {
			fmt.Printf("%#v", j.Socials)
		}
	}
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
	if levelNameTolvl("mod") <= UserLevel(userName).Level {
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
						fmt.Println(socexists)
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
		for i := 0; i < len(settings.Users)-1; i++ {
			if settings.Users[i].Name == userName || settings.Users[i].Nick == userName {
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
						if settings.Pets[0].FeedLimit > settings.Pets[0].Feed+1 {
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
						flimit, err := strconv.Atoi(cmdFields[1])
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
						if strings.Fields(msg)[1] == "movie" && strings.Fields(msg)[1] != "tv" {
							if searchresults[i].MediaType == "movie" {
								botSay(bb, fmt.Sprintf("📇 %s |%v| 📅 %s  🎥%s", searchresults[i].Title, searchresults[i].ID, searchresults[i].ReleaseDate, searchresults[i].Overview))
								fmt.Printf("📇 %s |%v| 📅 %s  🎥%s", searchresults[i].Title, searchresults[i].ID, searchresults[i].ReleaseDate, searchresults[i].Overview)
							}
						} else {
							if searchresults[i].MediaType == "tv" && strings.Fields(msg)[1] != "movie" {
								if searchresults[i].MediaType == "tv" {
									botSay(bb, fmt.Sprintf("📇 %s |%v| 📅 %s  📺%s", searchresults[i].Name, searchresults[i].ID, searchresults[i].FirstAirDate, searchresults[i].Overview))
									fmt.Printf("📇 %s |%v| 📅 %s  📺%s", searchresults[i].Name, searchresults[i].ID, searchresults[i].FirstAirDate, searchresults[i].Overview)
								}
							}
						}
					} else {
						botSay(bb, "Cant return adult movies")
					}
				}
			} else {
				botSay(bb, fmt.Sprintf("%s Cant find your movie", getAttributedUser(msg, true)))
			}
		} else {
			id, err := strconv.Atoi(strings.Fields(msg)[1])
			if err != nil {
				fmt.Println(err)
			} else {
				searchbyid := tmdbSearch(strings.Fields(msg)[1])
				if len(searchbyid.Results) > 0 {
					if len(strings.Fields(msg)) > 2 {
						if strings.Fields(msg)[2] == "movie" {
							movieData := tmdbMovie(id)
							botSay(bb, fmt.Sprintf("📇 %s |%v| 📅 %s  🎥%s", movieData.Title, movieData.ID, movieData.ReleaseDate, movieData.Overview))
							fmt.Printf("📇 %s |%v| 📅 %s  🎥%s", movieData.Title, movieData.ID, movieData.ReleaseDate, movieData.Overview)
						} else {
							tvData := tmdbTV(id)
							botSay(bb, fmt.Sprintf("📇 %s |%v| 📅 %s  📺%s", tvData.Name, tvData.ID, tvData.FirstAirDate, tvData.Overview))
							fmt.Printf("📇 %s |%v| 📅 %s  📺%s", tvData.Name, tvData.ID, tvData.FirstAirDate, tvData.Overview)
						}
					} else {
						movieData := tmdbMovie(id)
						botSay(bb, fmt.Sprintf("📇 %s |%v| 📅 %s  🎥%s", movieData.Title, movieData.ID, movieData.ReleaseDate, movieData.Overview))
						fmt.Printf("📇 %s |%v| 📅 %s  🎥%s", movieData.Title, movieData.ID, movieData.ReleaseDate, movieData.Overview)
					}
				}
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
