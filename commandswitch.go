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
				switch cmd {
				case "etbdown":
					CPrint("c", fmt.Sprintf("[%s] Shutdown command received. Shutting down now...\n", timeStamp()))
					bb.Disconnect()
				case "version":
					cmdVersion(bb)
					//region joke commands
				case "hi":
					cmdHi(bb, userName, cmd, msg)
				case "olive":
					fallthrough
				case "oil":
					fallthrough
				case "nvidia":
					fallthrough
				case "yogurt":
					cmdJoke(bb, cmd, msg)
				case "bofh":
					cmdJokeAPI(bb, cmd, msg)
				case "joke":
					fallthrough
				case "yoke":
					cmdJokeAPI(bb, cmd, msg)
				case "ban":
					cmdBan(bb, userName, cmd, msg)
				case "unban":
					cmdBan(bb, userName, cmd, msg)
				case "mic":
					cmdMic(bb)
					//endregion
				case "lurk":
					cmdLurk(bb, userName, cmd, msg)
				case "hype":
					cmdHype(bb, msg)
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
				case "commands":
					cmdList(bb, userName)
				case "fr":
					cmdFr(bb, userName, cmd, msg)
				case "love":
					cmdILOVE(bb, cmd, userName, msg)
				case "zoe":
					cmdZoe(bb, cmd, userName, msg)
					//Not Final
				case "socials":
					fallthrough
				case "github":
					fallthrough
				case "youtube":
					fallthrough
				case "itchio":
					fallthrough
				case "instagram":
					fallthrough
				case "artstation":
					fallthrough
				case "aboutme":
					fallthrough
				case "udemy":
					fallthrough
				case "discord":
					cmdSocial(bb, cmd)
				case "sudo":
					cmdVulgar(bb)
				case "pro": //check if they stream and say pro streamer otherwise pro viewer
					cmdSoon(bb)
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
				case "level":
					cmdLevel(bb, cmd, userName, msg)
				case "crypto":
					fallthrough
				case "cr":
					cmdCryptoExchange(bb, cmd, userName, msg)
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
	default:
		// fmt.Printf("no priv: %s", line)
	}
}
