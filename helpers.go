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
	if strings.Contains(msg, "@") {
		if at {
			return msg[strings.Index(msg, "@"):]
		} else {
			return msg[strings.Index(msg, "@")+1:]
		}
	} else {
		return ""
	}
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

// getCleanMessage removes attr and cmd from msg
func getCleanMessage(cmd, msg string) string {
	if strings.Contains(msg, "@") {
		return msg[len(cmd)+1 : strings.Index(msg, "@")-1]
	} else {
		return msg[len(cmd)+1:]
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

func LoadJSONFileTOStruct(jsonFileName string, onTo interface{}) {
	//read json here
	jsonFile, err := ioutil.ReadFile(jsonFileName)
	if nil != err {
		fmt.Println(err)
	}
	json.Unmarshal([]byte(jsonFile), &onTo)
}

func LoadJSONTOStruct(jsondata []byte, onTo interface{}) {
	//read json here
	json.Unmarshal(jsondata, &onTo)
}

func saveSettings() {
	file, _ := json.MarshalIndent(settings, "", "\t")
	_ = ioutil.WriteFile("testsave.json", file, 0644)
}

func cleanup() {
	saveSettings()
	fmt.Println("Saved Data")
}
