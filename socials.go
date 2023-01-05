package main

import (
	"fmt"
	"strings"
)

func cmdUPDSoc(bb *BasicBot, cmd, userName, msg string) {
	if UserLevel(userName).Level <= levelNameTolvl("mod") {
		if isCMD(cmd, msg) {
			botSay(bb, "Set a users social. ex. !updsoc @earentir github https://github.com/earentir")
		} else {
			fields := strings.Fields(msg[len(cmd)+2:])
			if len(fields) == 3 {
				for i := 0; i < len(userlist.Users); i++ {

					if strings.EqualFold(userlist.Users[i].Name, strings.ToLower(fields[0][1:])) {
						fmt.Println(userlist.Users[i].Name, strings.ToLower(fields[0][1:]))
						var socexists = false
						for j := 0; j < len(userlist.Users[i].Socials); j++ {
							if userlist.Users[i].Socials[j].SocNet == fields[1] {
								userlist.Users[i].Socials[j].Link = fields[2]
								botSay(bb, fmt.Sprintf("%s's %s profile is now %s", fields[0][1:], fields[1], fields[2]))
								socexists = true
							}
						}

						if !socexists {
							socs := Social{
								SocNet: fields[1],
								Link:   fields[2],
							}

							userlist.Users[i].Socials = append(userlist.Users[i].Socials, socs)

							botSay(bb, fmt.Sprintf("%s's %s profile is now %s", fields[0][1:], fields[1], fields[2]))
						}

						saveData("Users", userlist)
					}
				}
			} else {
				botSay(bb, "Set a users social. ex. !updsoc @earentir github https://github.com/earentir")
			}
		}
	} else {
		botSay(bb, "Only mods can update socials manually, please ask a mod for help.")
	}
}

func cmdILOVE(bb *BasicBot, cmd, userName, msg string) {
	if isCMD(cmd, msg) {
		botSay(bb, "Add what you love and it will be included in your SO, ex. !love coffee")
	} else {
		for i := 0; i < len(userlist.Users); i++ {
			if strings.EqualFold(userlist.Users[i].Name, userName) || strings.EqualFold(userlist.Users[i].Nick, userName) {
				userlist.Users[i].Love = msg[len(cmd)+2:]
				botSay(bb, fmt.Sprintf("You love %s, thank you for letting me know.", userlist.Users[i].Love))
			}
		}
	}
}
