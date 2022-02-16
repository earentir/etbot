package main

import (
	"fmt"
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
	found := []string{}

	for _, usr := range setusers {
		if usr.Name == userName {
			soc := usr.Socials
			found = strings.Fields(strings.ReplaceAll(strings.ReplaceAll(fmt.Sprintf("%v", soc), "{", ""), "}", ""))
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
