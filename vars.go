package main

import (
	"regexp"
	"time"
)

const (
	UTCFormat = time.RFC3339
	etbver    = "20220219A"
)

var (
	// Regex for parsing PRIVMSG strings.
	// First matched group is the user's name and the second matched group is the content of the
	// user's message.
	MsgRegex *regexp.Regexp = regexp.MustCompile(`^:(\w+)!\w+@\w+\.tmi\.twitch\.tv (PRIVMSG) #\w+(?: :(.*))?$`)

	// Regex for parsing user commands, from already parsed PRIVMSG strings.
	// First matched group is the command name and the second matched group is the argument for the
	// command.
	CmdRegex *regexp.Regexp = regexp.MustCompile(`^!(\w+)\s?(\w+)?`)
)
