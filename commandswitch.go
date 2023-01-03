package main

import (
	"fmt"
	"strconv"
	"time"
)

func ParseCommand(bb *BasicBot, msgType, msg, userName string) {
	switch msgType {
	case "PRIVMSG":

		if settings.Servers.BotServers.Log {
			var newmessage ChatMessage
			newmessage.Date = strconv.Itoa(int(time.Now().Unix()))
			newmessage.User = userName
			newmessage.Message = msg

			chatlog.ChatMessages = append(chatlog.ChatMessages, newmessage)
		}

		CPrint("c", fmt.Sprintf("[%s] %s: %s\n", timeStamp(), userName, msg))
		// parse commands from user message
		cmdMatch := CommandRegex.FindStringSubmatch(msg)
		if nil != cmdMatch {
			cmd := cmdMatch[0][1:]
			if CMDCanRun(userName, cmd) {
				if isUsrCmd(cmd) {
					usrCmdSay(bb, userName, cmd, msg)
				} else {
					switch cmd {
					case "hi":
						Call("cmdHi", bb, userName, cmd, msg)
					case "bofh":
						Call("cmdJokeAPI", bb, cmd, msg)
					case "yoke":
						Call("cmdJokeAPI", bb, cmd, msg)
					case "joke":
						cmdJoke(bb, userName, cmd, msg)
					case "lurk":
						cmdLurk(bb, userName, cmd, msg)
					case "ex":
						fallthrough
					case "exchange":
						cmdExchange(bb, msg)
					case "w":
						fallthrough
					case "weather":
						cmdWeather(bb, cmd, msg)
					case "so":
						cmdSO(bb, userName, cmd, msg)
					case "fr":
						cmdFr(bb, userName, cmd, msg)
					case "love":
						cmdILOVE(bb, cmd, userName, msg)
					case "zoe":
						cmdZoe(bb, cmd, userName, msg)
					case "socials":
						cmdSocial(bb, cmd)
					case "time":
						cmdTime(bb, cmd, msg)
					case "updsoc":
						cmdUPDSoc(bb, cmd, userName, msg)
					case "current":
						fallthrough
					case "project":
						cmdProject(bb, cmd, userName, msg)
					case "tmdb":
						cmdTMDB(bb, cmd, userName, msg)
					case "crypto":
						fallthrough
					case "cr":
						cmdCryptoExchange(bb, cmd, userName, msg)
					case "quote":
						cmdQuote(bb, cmd, userName, msg)
					case "year":
						cmdYear(bb, cmd, userName, msg)
					case "holidays":
						fallthrough
					case "daysoff":
						cmdDaysOff(bb, cmd, userName, msg)
					case "gpt":
						cmdGPTCompletion(bb, cmd, userName, msg, "completion")
					case "fact":
						cmdGPTCompletion(bb, cmd, userName, msg, "fact")
					case "gumroad":
						cmdGumroad(bb, cmd, userName, msg)
					case "goals":
						cmdGoals(bb, cmd, userName, msg)
					case "subs":
						cmdSubs(bb, cmd, userName, msg)

						//Not allowed to be renamed - System Commands
					case "lockdown":
						settings.General.Lockdown.Enabled = !settings.General.Lockdown.Enabled
						if settings.General.Lockdown.Enabled {
							botSay(bb, "Lockdown Initiated "+settings.General.Lockdown.Reason)
						} else {
							botSay(bb, "Lockdown Dissabled ")
						}
					case "level":
						cmdLevel(bb, cmd, userName, msg)
					case "setting":
						fallthrough
					case "set":
						cmdSetting(bb, cmd, userName, msg)
					case "etbdown":
						CPrint("c", fmt.Sprintf("[%s] Shutdown command received. Shutting down now...\n", timeStamp()))
						bb.Disconnect()
					case "version":
						cmdVersion(bb)
					case "commands":
						cmdList(bb, userName, msg)
					case "user":
						fallthrough
					case "usr":
						cmdUser(bb, cmd, userName, msg)
					case "savesettings":
						cmdSaveSettings(bb, cmd, userName, msg)
					default:
						// do nothing
						// }
					}
				}
			}
		}
	default:
		// fmt.Printf("no priv: %s", line)
	}
}
