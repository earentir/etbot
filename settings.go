package main

import (
	"strings"
	"time"
)

func getCMDS(userName string) string {
	cmds := settings.Commands
	var allcommands string = ""

	for _, cm := range cmds {
		if CMDCanRun(userName, cm.CommandName) {
			allcommands = allcommands + ", " + cm.CommandName
		}
	}

	return allcommands
}

func CMDCanRun(userName, cmd string) bool {
	var ourcmdopts CommandOption
	itcan := false

	ourcmdopts = getCMDOptions(cmd)

	if ourcmdopts.Enabled && (IsItOnTimeout(cmd, userName) || ourcmdopts.Lastuse == 0) {
		itcan = ourcmdopts.UserLevel >= UserLevel(userName).Level
		setCMDUsed(cmd)
	}

	return itcan
}

func getCMDOptions(cmd string) CommandOption {
	var commandOption CommandOption

	for i := 0; i <= len(settings.Commands)-1; i++ {
		if (cmd == settings.Commands[i].CommandName) || (cmd == settings.Commands[i].CommandOptions.Alias) {
			return settings.Commands[i].CommandOptions
		}
	}

	return commandOption
}

func setCMDUsed(cmd string) {
	for i := 0; i <= len(settings.Commands)-1; i++ {
		if (cmd == settings.Commands[i].CommandName) || (cmd == settings.Commands[i].CommandOptions.Alias) {
			settings.Commands[i].CommandOptions.Lastuse = int(time.Now().Unix())
			settings.Commands[i].CommandOptions.Counter++
		}
	}
}

func getUserSocials(userName string) []string {
	setusers := settings.Users
	var found []string

	for _, usr := range setusers {
		if usr.Name == userName {
			for _, k := range usr.Socials {
				found = append(found, k.Link)
			}
		}
	}
	return found
}

func SearchUser(userName string) bool {
	var found bool = false

	setusers := settings.Users
	for _, usr := range setusers {
		if userName == usr.Name {
			found = true
		}
	}
	return found
}

func UserLevel(userName string) UserLevelList {
	var userLevelReturn UserLevelList
	setusers := settings.Users

	for _, usr := range setusers {
		if userName == usr.Name {
			userLevel := settings.UserLevels
			for _, lvl := range userLevel {
				if usr.Type == lvl.Name {
					userLevelReturn = lvl
				}
			}
		}
	}

	return userLevelReturn
}

func levelNameTolvl(levelName string) int {
	var found int = -1
	for i := 0; i < len(settings.UserLevels); i++ {
		if levelName == settings.UserLevels[i].Name {
			found = settings.UserLevels[i].Level
		}
	}
	return found
}

	return found
}

func getUserData(userName string) User {
	var outUser User
	for i := 0; i < len(settings.Users); i++ {
		if strings.EqualFold(settings.Users[i].Name, userName) {
			outUser = settings.Users[i]
		}
	}
func addLurker(userName, cmd, msg string) {
	var (
		lurker LurkerList
		found  bool = false
	)
	for i := 0; i < len(settings.Lurklists); i++ {
		if strings.EqualFold(userName, settings.Lurklists[i].Lurker) {
			found = true
			settings.Lurklists[i].LurkedOn = int(time.Now().Unix())
			settings.Lurklists[i].LurkMessage = msg[len(cmd)+2:]
		}
	}

	if !found {
		lurker.Lurker = userName
		lurker.LurkedOn = int(time.Now().Unix())
		if msg == "" {
			lurker.LurkMessage = msg[len(cmd)+2:]
		} else {
			lurker.LurkMessage = ""
		}
		settings.Lurklists = append(settings.Lurklists, lurker)
	}
}

func isUserLurking(userName string) bool {
	var arethey bool = false
	for i := 0; i < len(settings.Lurklists); i++ {
		if strings.EqualFold(userName, settings.Lurklists[i].Lurker) {
			arethey = true
		}
	}
	return arethey
}

func removeLurker(userName string) {
	var newLurkList []LurkerList

	for i := len(settings.Lurklists) - 1; i >= 0; i-- {
		if !strings.EqualFold(userName, settings.Lurklists[i].Lurker) {
			{
				newLurkList = append(newLurkList, settings.Lurklists[i])
			}
		}
	}

	settings.Lurklists = newLurkList
}
