package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

// getAttributedUser at returns @ or not in the reply
func getAttributedUser(msg string, at bool) string {
	var attrUser string = ""
	fields := strings.Fields(msg)

	for _, j := range fields {
		if strings.Contains(j, "@") {
			if at {
				return j // msg[strings.Index(msg, "@"):]
			} else {
				return j[1:] //msg[strings.Index(msg, "@")+1:]
			}
		}
	}

	return strings.ToLower(attrUser)
}

// isAttr return if there is an attr user in msg
func isAttr(msg string) bool {
	exists := false
	if getAttributedUser(msg, true) != "" {
		exists = true
	}
	return exists
}

// isCMD Checks if we only got the cmd in the prompt
func isCMD(cmd, msg string) bool {
	if cmd == strings.ReplaceAll(msg, "!", "") {
		return true
	} else {
		return false
	}
}

func IsItOnTimeout(cmd, userName string) bool {
	var allowed bool = false

	if (UserLevel(userName).Cooldown * (getCMDOptions(cmd).Cooldown / 1000)) < int(time.Now().Unix())-getCMDOptions(cmd).Lastuse {
		allowed = true
	}
	return allowed
}

func CPrint(color, msg string) {
	var colorCode string
	var resetcolor string = "\033[0m"

	switch strings.ToLower(color) {
	case "r":
		fallthrough
	case "red":
		colorCode = "\033[31m"

	case "g":
		fallthrough
	case "green":
		colorCode = "\033[32m"

	case "y":
		fallthrough
	case "yellow":
		colorCode = "\033[33m"

	case "b":
		fallthrough
	case "blue":
		colorCode = "\033[34m"

	case "p":
		fallthrough
	case "purple":
		colorCode = "\033[35m"

	case "c":
		fallthrough
	case "cyan":
		colorCode = "\033[36m"

	default:
		colorCode = resetcolor
	}

	fmt.Println(string(colorCode), msg)
	fmt.Print(string(resetcolor))
}

func timeStamp() string {
	return TimeStamp(UTCFormat)
}

func TimeStamp(format string) string {
	return time.Now().Format(format)
}

//read json file to struct here
func LoadJSONFileTOStruct(jsonFileName string, onTo interface{}) {
	jsonFile, err := ioutil.ReadFile(jsonFileName)
	if nil != err {
		fmt.Println(err)
	}
	json.Unmarshal([]byte(jsonFile), &onTo)
}

//read json data to struct here
func LoadJSONTOStruct(jsondata []byte, onTo interface{}) {
	err := json.Unmarshal(jsondata, &onTo)
	if err != nil {
		fmt.Println("LoadJSONTOStruct>\n", err)
	}

}

func saveSettings() {
	settingsfile, _ := json.MarshalIndent(settings, "", "\t")
	_ = ioutil.WriteFile("settings/etb-settings.json", settingsfile, 0644)

	usrcmdfile, _ := json.MarshalIndent(usercommands, "", "\t")
	_ = ioutil.WriteFile("settings/usr-cmd.json", usrcmdfile, 0644)
}

func cleanup() {
	saveSettings()
	fmt.Println("Saved Data")
}

func limitOverview(overview string) string {
	var newOverview string

	if len(overview) <= 233 {
		newOverview = overview
	} else {
		newOverview = overview[:233]
	}

	return newOverview
}

func timeZone(location string) string {
	locs := []string{"", "Europe", "Africa", "America", "Asia", "Atlantic", "Australia", "Brazil", "Canada", "Indian", "US"}
	var tme string
	var msgOut string
	var newloc string

	if strings.Contains(location, "/") {
		msgOut = tzNow(location)
	} else {
		for _, l := range locs {
			fields := strings.Fields(location)

			if l == "" {
				newloc = location
			} else {
				if len(fields) > 1 {
					newloc = l + "/" + strings.Title(strings.ToLower(fields[0])) + "_" + strings.Title(strings.ToLower(fields[1]))
				} else {
					newloc = l + "/" + strings.Title(strings.ToLower(location))
				}
			}
			tme = tzNow(newloc)
			if tme != "" {
				msgOut = tme
			}
		}
	}

	return msgOut
}

func tzNow(locationToLookup string) string {
	loc, err := time.LoadLocation(locationToLookup)
	outMsg := ""
	if err == nil {
		now := time.Now()
		outMsg = fmt.Sprintf("%v", now.In(loc))
	}

	return outMsg
}

func progressbar(percent float64) string {
	var equals int = int(percent / 10)
	var msgOUt string = ""

	for i := 0; i < 10; i++ {
		if i <= equals {
			msgOUt = msgOUt + "="
		} else {
			msgOUt = msgOUt + "-"
		}
	}

	return msgOUt
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

func isUsrCmdAlias(index int, cmd string) bool {
	var found bool = false
	for i := 0; i < len(usercommands[index].Alias); i++ {
		if usercommands[index].Alias[i] == cmd {
			found = true
		}
	}

	return found
}

func cleanMessage(msg string) string {
	var (
		cleanmsg string   = ""
		cmd      string   = getCommand(msg)
		attrUser string   = getAttributedUser(msg, false)
		fields   []string = strings.Fields(msg)
	)

	if attrUser != "" {
		cleanmsg = msg[len(cmd)+1+1+len(fields[1])+1+len(attrUser)+1:]
	} else {
		if len(fields) > 1 {
			if len(msg) >= len(cmd)+1+1+len(fields[1])+1 {
				cleanmsg = msg[len(cmd)+1+1+len(fields[1])+1:]
			} else {
				cleanmsg = msg[len(cmd)+1+1+len(fields[1]):]
			}
		}
	}

	return cleanmsg
}

func getCommand(msg string) string {
	cmdmatch := CommandRegex.FindStringSubmatch(msg)

	var cmd string = ""

	if len(cmdmatch[0]) > 1 {
		cmd = cmdmatch[0][1:]
	} else {
		cmd = cmdmatch[0]
	}
	return cmd
}

func getKeyWord(keyword, msg string) int {
	var intOut int = -1
	intOut = strings.Index(msg, keyword)
	return intOut
}
