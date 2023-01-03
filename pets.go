package main

import (
	"fmt"
	"strconv"
	"strings"
)

func cmdZoe(bb *BasicBot, cmd, userName, msg string) {
	if isCMD(cmd, msg) {
		var petData = ""

		if len(petlist.Pets) > 0 {
			petData = fmt.Sprintf(" | Treat: %v/%v  Petting Minutes: %v", petlist.Pets[0].Feed, petlist.Pets[0].FeedLimit, petlist.Pets[0].Pet)
		}

		botSay(bb, fmt.Sprintf("!zoe pet or !zoe feed or !zoe name %s", petData))
	} else {
		cmdFields := strings.Fields(msg)
		if len(cmdFields) > 1 {
			if len(petlist.Pets) < 1 {
				var pet Pet
				pet.Feed = 0
				pet.Pet = 0

				petlist.Pets = append(petlist.Pets, pet)
				botSay(bb, "Pet Registered")
			} else {
				switch cmdFields[1] {
				case "pet":
					if len(cmdFields) == 3 {
						if userName == settings.General.Twitch.Channel {
							rem, err := strconv.Atoi(cmdFields[2])
							if err != nil {
								petlist.Pets[0].Pet = 0
								botSay(bb, "Pet Reset")
							} else {
								petlist.Pets[0].Pet = petlist.Pets[0].Pet - rem
								botSay(bb, fmt.Sprintf("%s needs to be peted for %v minutes", petlist.Pets[0].Name, petlist.Pets[0].Pet))
							}
						}
					} else {
						petlist.Pets[0].Pet++
						botSay(bb, fmt.Sprintf("%s needs to be peted for %v minutes", petlist.Pets[0].Name, petlist.Pets[0].Pet))
					}
				case "treat":
					fallthrough
				case "feed":
					if len(cmdFields) == 3 {
						if userName == settings.General.Twitch.Channel {
							rem, err := strconv.Atoi(cmdFields[2])
							if err != nil {
								petlist.Pets[0].Feed = 0
								botSay(bb, "Feed Reset")
							} else {
								petlist.Pets[0].Feed = petlist.Pets[0].Feed - rem
								botSay(bb, fmt.Sprintf("%s needs to be given %v treats", petlist.Pets[0].Name, petlist.Pets[0].Feed))
							}
						}
					} else {
						if petlist.Pets[0].FeedLimit < petlist.Pets[0].Feed {
							botSay(bb, "No more feed can be added, feed the pet first")
						} else {
							petlist.Pets[0].Feed++
							botSay(bb, fmt.Sprintf("%s needs to be given %v treats", petlist.Pets[0].Name, petlist.Pets[0].Feed))
						}
					}
				case "name":
					if len(cmdFields) == 3 {
						if userName == settings.General.Twitch.Channel {
							petlist.Pets[0].Name = cmdFields[2]
						}
					} else {
						botSay(bb, fmt.Sprintf("Pet name is %s", petlist.Pets[0].Name))
					}
				case "limit":
					if len(cmdFields) == 3 {
						flimit, err := strconv.Atoi(cmdFields[2])
						if err != nil {
						} else {
							petlist.Pets[0].FeedLimit = flimit
						}
					}
				}
			}
			//save the pets
			saveData("Pets", petlist)
		} else {
			botSay(bb, fmt.Sprintf("%s | Treat: %v (%v)  Petting Minutes: %v", "!zoe pet or !zoe feed or !zoe name", petlist.Pets[0].Feed, petlist.Pets[0].FeedLimit, petlist.Pets[0].Pet))
		}
	}
}
