package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
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
				return j
			} else {
				return j[1:]
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

// IstItOnTimeout return the user current timeout
func IsItOnTimeout(cmd, userName string) bool {
	var allowed bool = false

	if (UserLevel(userName).Cooldown * (getCMDOptions(cmd).Cooldown / 1000)) < int(time.Now().Unix())-getCMDOptions(cmd).Lastuse {
		allowed = true
	}
	return allowed
}

// CPrint print out the msg in the selected colour
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

// return time in utc format
func timeStamp() string {
	return TimeStamp(UTCFormat)
}

// create time in the selected format
func TimeStamp(format string) string {
	return time.Now().Format(format)
}

// read json file to struct here
func LoadJSONFileTOStruct(jsonFileName string, onTo interface{}) {
	jsonFile, err := ioutil.ReadFile(jsonFileName)
	if nil != err {
		fmt.Println(err)
	}
	json.Unmarshal([]byte(jsonFile), &onTo)
}

// read json data to struct here
func LoadJSONTOStruct(jsondata []byte, onTo interface{}) {
	if err := json.Unmarshal(jsondata, &onTo); err != nil {
		fmt.Println("LoadJSONTOStruct>\n", err)
	}

}

// save settings (default data only)
func saveSettings() {
	saveData("Settings", settings)
	saveData("UserCommands", usercommands)
	saveChatLog()
}

// we cleanup when user terminates
func cleanup() {
	saveSettings()
	fmt.Println("Saved Data")
}

// we cut off the overview to fit inside the minimum 512 bytes message
func limitOverview(overview string) string {
	var newOverview string

	if len(overview) <= 233 {
		newOverview = overview
	} else {
		newOverview = overview[:233]
	}

	return newOverview
}

// calculate the timezone of current time by using the tz capital city name
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

// convert tz to time now
func tzNow(locationToLookup string) string {
	loc, err := time.LoadLocation(locationToLookup)
	outMsg := ""
	if err == nil {
		now := time.Now()
		outMsg = fmt.Sprintf("%v", now.In(loc))
	}

	return outMsg
}

// simple progress bar for use in the year commands
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

// check if cmd is in the users commands
func isUsrCmd(cmd string) bool {
	var found bool = false
	for i := 0; i < len(usercommands); i++ {
		if strings.EqualFold(usercommands[i].UserCmdName, cmd) || isUsrCmdAlias(i, cmd) {
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

// loop over user commands aliase and return true if its in the array
func isUsrCmdAlias(index int, cmd string) bool {
	var found bool = false
	for i := 0; i < len(usercommands[index].Alias); i++ {
		if usercommands[index].Alias[i] == cmd {
			found = true
		}
	}

	return found
}

// check if cmd is in the system commands
func isSysCmd(cmd string) int {
	var found int = -1
	for i := 0; i < len(systemcommands.Commands); i++ {
		if strings.EqualFold(systemcommands.Commands[i].Name, cmd) || isSysCmdAlias(i, cmd) {
			found = i
		}
	}

	return found
}

// loop over system command aliases and return true if its in the array
func isSysCmdAlias(index int, cmd string) bool {
	var found bool = false
	for i := 0; i < len(systemcommands.Commands[index].Alias); i++ {
		if systemcommands.Commands[index].Alias[i] == cmd {
			found = true
		}
	}

	return found
}

// clean everything from a msg
func getCleanMessage(msg string) string {
	var (
		cmd       string = getCommand(msg)
		msgOut    string = ""
		attrUser  string = ""
		cmdIndex  int    = strings.Index(msg, cmd)
		attrIndex int    = 0
	)

	if len(cmd) > len(msg) {
		msgOut = msg[cmdIndex+len(cmd)+1:]
	} else {
		msgOut = msg
	}

	attrUser = getAttributedUser(msgOut, false)
	attrIndex = strings.Index(msgOut, attrUser)

	if len(attrUser) > 0 {
		msgOut = msgOut[attrIndex+len(attrUser)+1:]
	}

	return msgOut
}

// parse msg to get a command from it (if it exists)
func getCommand(msg string) string {
	var cmdmatch []string = CommandRegex.FindStringSubmatch(msg)
	var cmd string = ""

	if len(cmdmatch) > 0 {
		if len(cmdmatch[0]) > 0 {
			cmd = cmdmatch[0][1:]
		} else {
			cmd = cmdmatch[0]
		}
	}

	return cmd
}

// parse message for a keyword and get its starting possition
func getKeyWord(keyword, msg string) int {
	var intOut int = -1
	intOut = strings.Index(msg, keyword)
	return intOut
}

// save data by giving it the path in the settings file and the struct that holds the data
func saveData(settingsName string, thestruct interface{}) {
	var settingsPath string = settings.FilePaths.SettingsDir
	var fileName string = filepath.Join(settingsPath, getFieldFP(settingsName))

	if datafile, err := json.MarshalIndent(thestruct, "", "\t"); err == nil {
		if err = ioutil.WriteFile(fileName, datafile, 0644); err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}

// load data by giving it the path in the settings file and the struct that holds the data
func loadData(settingsName []string, thestruct interface{}) {
	var settingsPath string = settings.FilePaths.SettingsDir
	var fileName string = filepath.Join(settingsPath, getField(&settings, settingsName))

	if _, err := os.Stat(fileName); err == nil {
		LoadJSONFileTOStruct(fileName, thestruct)
	}
}

// from struct field name using reflection
func getField(v *Settings, fields []string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(fields[0]).FieldByName(fields[1])
	return f.String()
}

// call a func using its name stored in a string
func Call(funcName string, params ...interface{}) (result interface{}, err error) {
	StubStorage = map[string]interface{}{
		"cmdHi": cmdHi,
		"cmdSO": cmdSO,
	}

	f := reflect.ValueOf(StubStorage[funcName])
	if len(params) != f.Type().NumIn() {
		fmt.Println("too many parameters")
		return
	}

	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}

	res := f.Call(in)
	if len(res) > 0 {
		result = res[0].Interface()
	}
	return
}

// check if we can load the bot
func checkLoadStatus() bool {
	return true
}
