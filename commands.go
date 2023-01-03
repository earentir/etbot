package main

import (
	"fmt"
	"math/rand"
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
		fmt.Println("Console Print Only")
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

func cmdLevel(bb *BasicBot, cmd, userName, msg string) {
	if isAttr(msg) {
		atrUser := getAttributedUser(msg, false)
		botSay(bb, fmt.Sprintf("@%s level is %v as a %s", atrUser, UserLevel(atrUser).Level, UserLevel(atrUser).Name))
	} else {
		botSay(bb, fmt.Sprintf("@%s level is %v as a %s", userName, UserLevel(userName).Level, UserLevel(userName).Name))
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
		saveDefaultData()
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
