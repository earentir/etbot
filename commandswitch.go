package main

import (
	"fmt"
)

func ParseCommand(bb *BasicBot, msgType, msg, userName string) {
	switch msgType {
	case "PRIVMSG":

		CPrint("c", fmt.Sprintf("[%s] %s: %s\n", timeStamp(), userName, msg))
		// parse commands from user message
		cmdMatches := CmdRegex.FindStringSubmatch(msg)
		if nil != cmdMatches {
			cmd := cmdMatches[1]
			if CMDCanRun(userName, cmd) {

				if isUsrCmd(cmd) {
					usrCmdSay(bb, userName, cmd, msg)
				} else {
					switch cmd {
					case "hi":
						cmdHi(bb, userName, cmd, msg)
					case "bofh":
						cmdJokeAPI(bb, cmd, msg)
					case "yoke":
						cmdJokeAPI(bb, cmd, msg)
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

						//Not allowed to be renamed - System Commands
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
						cmdList(bb, userName)
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
