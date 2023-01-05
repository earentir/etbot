package main

import (
	"fmt"
	"strings"
	"time"
)

func cmdLurk(bb *BasicBot, userName, cmd, msg string) {
	var (
		lurkers  []string
		lurklist LurkList
	)

	loadData("Lurkers", &lurklist)

	if isCMD(cmd, msg) {
		msgOut := fmt.Sprintf(getCMDOptions("lurk").Msg, userName)
		addLurker(userName, cmd, msg)
		botSay(bb, msgOut)
	} else {
		switch strings.Fields(msg)[1] {
		case "list":
			for i := 0; i < len(lurklist.Lurkers); i++ {
				lurkers = append(lurkers, lurklist.Lurkers[i].Name)
			}
			botSay(bb, fmt.Sprintf("Current Lurkers %s", lurkers))
		case "help":
			botSay(bb, "!lurk | !lurk [optional reasons | !lurk list]")
		default:
			addLurker(userName, cmd, msg)
			msgOut := fmt.Sprintf(getCMDOptions("lurk").Atmsg, userName, getCleanMessage(msg))
			botSay(bb, msgOut)
		}
	}
}

func cmdUnlurk(bb *BasicBot, userName string) {
	var lurklist LurkList

	loadData("Lurkers", &lurklist)

	for i := 0; i < len(lurklist.Lurkers); i++ {
		if strings.EqualFold(userName, lurklist.Lurkers[i].Name) {
			now := time.Unix(time.Now().Unix(), 0)
			lurkedon := time.Unix(lurklist.Lurkers[i].Date, 0)

			if lurklist.Lurkers[i].Message == "" {
				botSay(bb, fmt.Sprintf("Welcome back @%s, you have been gone for %v", userName, now.Sub(lurkedon)))
			} else {
				botSay(bb, fmt.Sprintf("Welcome back @%s, how was your %s, you have been gone for %v", userName, lurklist.Lurkers[i].Message, now.Sub(lurkedon)))
			}
			removeLurker(userName)
		}
	}
}
