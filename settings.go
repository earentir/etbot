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

	ourcmdopts = CmdOpts(cmd)

	if ourcmdopts.Enabled {
		itcan = ourcmdopts.UserLevel >= UserLevel(userName)
	}

	return itcan
}

func CmdOpts(cmd string) CommandOption {
	var commandOption CommandOption

	for i := 0; i <= len(settings.Commands); i++ {
		if (cmd == settings.Commands[i].CommandName) || (cmd == settings.Commands[i].CommandOptions.Alias) {
			settings.Commands[i].CommandOptions.Counter++
			settings.Commands[i].CommandOptions.Lastuse = int(time.Now().Unix())

			return settings.Commands[i].CommandOptions
		}
	}

	return commandOption
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

func convertLevelType(LevelType string) int {
	var level int = 10
	setusers := settings.UserLevels
	for _, lvl := range setusers {
		if LevelType == lvl.Name {
			level = lvl.Level
		}
	}
	return level
}

func getLevelCoolDown(level int) int {
	var cooldown int = 30000
	setusers := settings.UserLevels
	for _, lvl := range setusers {
		if level == lvl.Level {
			cooldown = lvl.Cooldown
		}
	}
	return cooldown
}

func UserLevel(userName string) int {
	var level int = 10
	setusers := settings.Users
	for _, usr := range setusers {
		if userName == usr.Name {
			level = convertLevelType(usr.Type)
		}
	}
	return level
}
