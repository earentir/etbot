package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

func getCMDS(userName string) string {
	cmds := settings.Commands
	usrcmds := usercommands
	var allcommands []string

	for _, cm := range cmds {
		if CMDCanRun(userName, cm.CommandName) {
			allcommands = append(allcommands, cm.CommandName)
		}
	}

	for _, cm := range usrcmds {
		allcommands = append(allcommands, cm.UserCmdName)
	}

	sort.Strings(allcommands)
	return fmt.Sprintf("%s", allcommands)
}

func CMDCanRun(userName, cmd string) bool {
	var ourcmdopts CommandOption
	itcan := false

	ourcmdopts = getCMDOptions(cmd)

	if ourcmdopts.Enabled && (IsItOnTimeout(cmd, userName) || ourcmdopts.Lastuse == 0) {
		itcan = ourcmdopts.UserLevel >= UserLevel(userName).Level
		setCMDUsed(cmd)
	}

	fmt.Printf("Usr: %s | CMD: %s | Perm: %v\n", userName, cmd, itcan)

	return itcan
}

func getCMDOptions(cmd string) CommandOption {
	var commandOption CommandOption
	var cmdFound bool = false

	if isUsrCmd(cmd) {
		for i := 0; i <= len(usercommands)-1; i++ {
			if (cmd == usercommands[i].UserCmdName) || isUsrCmdAlias(i, cmd) {
				commandOption = usercommands[i].UserCmdOptions
				cmdFound = true
			}
		}
	} else {
		for i := 0; i <= len(settings.Commands)-1; i++ {
			if (cmd == settings.Commands[i].CommandName) || (cmd == settings.Commands[i].CommandOptions.Alias) {
				commandOption = settings.Commands[i].CommandOptions
				cmdFound = true
			}
		}
	}

	if !cmdFound {
		if getUserSocial(cmd) != "" {
			commandOption.UserLevel = 10
			commandOption.Enabled = true
			commandOption.Cooldown = 4000
			commandOption.Lastuse = 0
		}
	}

	return commandOption
}

func setCMDUsed(cmd string) {
	if isUsrCmd(cmd) {
		for i := 0; i <= len(usercommands)-1; i++ {
			if (cmd == usercommands[i].UserCmdName) || isUsrCmdAlias(i, cmd) {
				usercommands[i].UserCmdOptions.Lastuse = int(time.Now().Unix())
				usercommands[i].UserCmdOptions.Counter++
			}
		}
	} else {
		for i := 0; i <= len(settings.Commands)-1; i++ {
			if (cmd == settings.Commands[i].CommandName) || (cmd == settings.Commands[i].CommandOptions.Alias) {
				settings.Commands[i].CommandOptions.Lastuse = int(time.Now().Unix())
				settings.Commands[i].CommandOptions.Counter++
			}
		}
	}
}

func getUserSocials(userName string) string {
	setusers := settings.Users
	var found string

	for _, usr := range setusers {
		if usr.Name == userName {
			for _, k := range usr.Socials {
				found = found + k.Link + " "
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
	var (
		userLevelReturn UserLevelList
		found           int = -1
	)

	for i := 0; i < len(settings.Users); i++ {
		if userName == settings.Users[i].Name {
			found = i
			for _, lvl := range settings.UserLevels {
				if settings.Users[i].Type == lvl.Name {
					userLevelReturn = lvl
				}
			}
		}
	}

	if found == -1 {
		userLevelReturn.Cooldown = 10
		userLevelReturn.Name = "viewer"
		userLevelReturn.Level = 10
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

func lvlToLevelName(lvl int) string {
	var found string
	for i := 0; i < len(settings.UserLevels); i++ {
		if lvl == settings.UserLevels[i].Level {
			found = settings.UserLevels[i].Name
		}
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
	return outUser
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

func addQuote(userName, attrUser, cleanmsg string) string {
	var (
		quotelist QuoteList
		found     bool = false
	)

	for i := 0; i < len(settings.Quotes); i++ {
		if settings.Quotes[i].QuotedMessage == cleanmsg {
			found = true
		}
	}

	if !found {
		quotelist.Quoter = userName
		quotelist.QuotedMessage = cleanmsg
		quotelist.AtributedUser = attrUser
		quotelist.QuoteDate = time.Now().Unix()
		settings.Quotes = append(settings.Quotes, quotelist)
	}

	return fmt.Sprintf("Quote from %s. \"%s\" added ", quotelist.Quoter, cleanmsg)
}

//add check for duplicates
func addUser(userToAdd, UserType string) string {
	var (
		found   bool = false
		newUser User
		msgOut  string = ""
	)

	for i := 0; i < len(settings.Users); i++ {
		if userToAdd == settings.Users[i].Name {
			found = true
		}
	}

	if !found {
		newUser.Name = userToAdd
		newUser.Type = UserType

		settings.Users = append(settings.Users, newUser)
		msgOut = fmt.Sprintf("User %s was added as a %s", userToAdd, UserType)
	} else {
		msgOut = fmt.Sprintf("User %s already exists", userToAdd)
	}

	return msgOut
}

func delUser(userToDelete string) string {

	var (
		newUserList []User
		msgOut      string = ""
		found       bool   = false
	)

	for i := len(settings.Users) - 1; i >= 0; i-- {
		if !strings.EqualFold(userToDelete, settings.Users[i].Name) {
			newUserList = append(newUserList, settings.Users[i])
		} else {
			found = true
		}
	}

	if found {
		msgOut = fmt.Sprintf("User %s deleted", userToDelete)
	} else {
		msgOut = fmt.Sprintf("User %s not found, nothing deleted", userToDelete)
	}

	settings.Users = newUserList
	return msgOut
}
