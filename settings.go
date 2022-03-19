package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

func getCMDS(userName string) string {
	var allcommands []string

	for _, cm := range systemcommands.Commands {
		if CMDCanRun(userName, cm.Name) {
			allcommands = append(allcommands, cm.Name)
		}
	}

	for _, cm := range usercommands {
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

	fmt.Printf("usr: %s [%v] | cmd: %s [%v]\nen: %v | time: %v | usrCMD: %v\nperm: %v\n", userName, levelNameTolvl(getUserData(userName).Type), cmd, ourcmdopts.UserLevel, ourcmdopts.Enabled, IsItOnTimeout(cmd, userName), isUsrCmd(cmd), itcan)
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
		for i := 0; i <= len(systemcommands.Commands)-1; i++ {
			if cmd == systemcommands.Commands[i].Name || isSysCmdAlias(i, cmd) {
				commandOption = systemcommands.Commands[i].Options
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
	if cmd != "commands" {
		if isUsrCmd(cmd) {
			for i := 0; i <= len(usercommands)-1; i++ {
				if (cmd == usercommands[i].UserCmdName) || isUsrCmdAlias(i, cmd) {
					usercommands[i].UserCmdOptions.Lastuse = int(time.Now().Unix())
					usercommands[i].UserCmdOptions.Counter++

					saveData([]string{"FilePaths", "UserCommands"}, usercommands)
				}
			}
		} else {
			cmdIndex := isSysCmd(cmd)
			if cmdIndex > -1 {
				systemcommands.Commands[cmdIndex].Options.Lastuse = int(time.Now().Unix())
				systemcommands.Commands[cmdIndex].Options.Counter++

				saveData([]string{"FilePaths", "SystemCommands"}, systemcommands)
			}
		}
	}

}

func getUserSocials(userName string) string {
	var found string

	for _, usr := range userlist.Users {
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

	for _, usr := range userlist.Users {
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

	if len(userlist.Users) > 0 {
		for i := 0; i < len(userlist.Users); i++ {
			if userName == userlist.Users[i].Name {
				found = i
				for _, lvl := range settings.UserLevels {
					if userlist.Users[i].Type == lvl.Name {
						userLevelReturn = lvl
					}
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
	for i := 0; i < len(userlist.Users); i++ {
		if strings.EqualFold(userlist.Users[i].Name, userName) {
			outUser = userlist.Users[i]
		}
	}
	return outUser
}

func addLurker(userName, cmd, msg string) {
	var (
		lurklist LurkList

		lurker Lurker
		found  bool = false
	)

	loadData([]string{"FilePaths", "Lurkers"}, lurklist)

	for i := 0; i < len(lurklist.Lurkers); i++ {
		if strings.EqualFold(userName, lurklist.Lurkers[i].Name) {
			found = true
			lurklist.Lurkers[i].Date = int(time.Now().Unix())
			lurklist.Lurkers[i].Message = msg[len(cmd)+1:]
		}
	}

	if !found {
		lurker.Name = userName
		lurker.Date = int(time.Now().Unix())
		if msg != "" {
			lurker.Message = msg[len(cmd)+1:]
		} else {
			lurker.Message = ""
		}
		lurklist.Lurkers = append(lurklist.Lurkers, lurker)
	}

	saveData([]string{"FilePaths", "Lurkers"}, lurklist)
}

func removeLurker(userName string) {
	var newLurkList []Lurker
	var lurklist LurkList

	loadData([]string{"FilePaths", "Lurkers"}, lurklist)

	for i := len(lurklist.Lurkers) - 1; i >= 0; i-- {
		if !strings.EqualFold(userName, lurklist.Lurkers[i].Name) {
			{
				newLurkList = append(newLurkList, lurklist.Lurkers[i])
			}
		}
	}

	lurklist.Lurkers = newLurkList

	saveData([]string{"FilePaths", "Lurkers"}, lurklist)
}

func addQuote(userName, attrUser, cleanmsg string) string {
	var (
		quotelist QuoteList
		qitem     QuoteItem
		found     bool = false
	)

	loadData([]string{"FilePaths", "Quotes"}, quotelist)

	for i := 0; i < len(quotelist.QuoteItems); i++ {
		if quotelist.QuoteItems[i].QuotedMessage == cleanmsg {
			found = true
		}
	}

	if !found {
		qitem.Quoter = userName
		qitem.QuotedMessage = cleanmsg
		qitem.AtributedUser = attrUser
		qitem.QuoteDate = time.Now().Unix()
		quotelist.QuoteItems = append(quotelist.QuoteItems, qitem)
	}

	saveData([]string{"FilePaths", "Quotes"}, quotelist)
	return fmt.Sprintf("Quote from %s. \"%s\" added ", userName, cleanmsg)
}

func addJoke(userName, attrUser, cleanmsg string) string {
	var (
		jokelist JokeList
		jitem    JokeItem
		found    bool = false
	)

	loadData([]string{"FilePaths", "Jokes"}, jokelist)

	for i := 0; i < len(jokelist.JokeItems); i++ {
		if jokelist.JokeItems[i].JokeMessage == cleanmsg {
			found = true
		}
	}

	if !found {
		jitem.Jokster = userName
		jitem.JokeMessage = cleanmsg
		jitem.AtributedUser = attrUser
		jitem.JokeDate = time.Now().Unix()
		jokelist.JokeItems = append(jokelist.JokeItems, jitem)
	}

	saveData([]string{"FilePaths", "Jokes"}, jokelist)
	return fmt.Sprintf("Joke from %s. \"%s\" added ", userName, cleanmsg)
}

//add check for duplicates
func addUser(userToAdd, userType string) string {
	var (
		found   bool = false
		newUser User
		msgOut  string = ""
	)

	for i := 0; i < len(userlist.Users); i++ {
		if userToAdd == userlist.Users[i].Name {
			found = true
		}
	}

	if !found {
		var soc Social
		soc.SocNet = "twitch"
		soc.Link = "https://twitch.tv/" + userToAdd

		newUser.Name = strings.ToLower(userToAdd)
		newUser.Type = strings.ToLower(userType)
		newUser.Socials = append(newUser.Socials, soc)

		userlist.Users = append(userlist.Users, newUser)
		msgOut = fmt.Sprintf("User %s was added as a %s", userToAdd, userType)
	} else {
		msgOut = fmt.Sprintf("User %s already exists", userToAdd)
	}

	saveData([]string{"FilePaths", "Users"}, userlist)
	return msgOut
}

func delUser(userToDelete string) string {
	var (
		newUserList []User
		msgOut      string = ""
		found       bool   = false
	)

	for i := len(userlist.Users) - 1; i >= 0; i-- {
		if !strings.EqualFold(userToDelete, userlist.Users[i].Name) {
			newUserList = append(newUserList, userlist.Users[i])
		} else {
			found = true
		}
	}

	if found {
		msgOut = fmt.Sprintf("User %s deleted", userToDelete)
	} else {
		msgOut = fmt.Sprintf("User %s not found, nothing deleted", userToDelete)
	}

	userlist.Users = newUserList

	saveData([]string{"FilePaths", "Users"}, userlist)
	return msgOut
}

func getUserSocial(socnet string) string {
	user := getUserData(settings.General.Twitch.Channel)
	var outMessage string = ""
	for i := 0; i < len(user.Socials); i++ {
		if strings.EqualFold(socnet, user.Socials[i].SocNet) {
			outMessage = user.Socials[i].Link
		}
	}

	return outMessage
}
