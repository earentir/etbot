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
					cmdJoke(bb, cmd)
				case "yogurt":
					cmdJoke(bb, cmd)
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

					//Not Final
				case "socials":
					fallthrough
				case "discord":
					cmdNo(bb)
				case "sudo":
					cmdVulgar(bb)
				case "pro": //check if they stream and say pro streamer otherwise pro viewer
					cmdSoon(bb)
				case "time":
					cmdTime(bb, cmd, msg)
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
