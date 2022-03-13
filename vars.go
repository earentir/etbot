package main

import (
	"regexp"
	"time"
)

const (
	UTCFormat = time.RFC3339
	etbver    = "20220312A"
)

type stubMapping map[string]interface{}

type Bot interface {
	// Opens a connection to the Twitch.tv IRC chat server.
	Connect()

	// Closes a connection to the Twitch.tv IRC chat server.
	Disconnect()

	// Listens to chat messages and PING request from the IRC server.
	HandleChat() error

	// Joins a specific chat channel.
	JoinChannel()

	// Parses credentials needed for authentication.
	ReadCredentials() error

	// Sends a message to the connected channel.
	Say(msg string) error

	// Attempts to keep the bot connected and handling chat.
	Start()
}

var (
	// Regex for parsing PRIVMSG strings.
	// First matched group is the user's name and the second matched group is the content of the
	// user's message.
	MsgRegex *regexp.Regexp = regexp.MustCompile(`^:(\w+)!\w+@\w+\.tmi\.twitch\.tv (PRIVMSG) #\w+(?: :(.*))?$`)

	// Regex for parsing user commands, from already parsed PRIVMSG strings.
	CommandRegex *regexp.Regexp = regexp.MustCompile(`!(.*?)[^\s]+`)

	// Store credentials from etb-auth.json
	creds Credentials

	// Store In Memory accessible settings from et-settings.json
	settings Settings

	//Store User Commands in memory
	usercommands   UserCommands
	systemcommands CommandList
	userlist       UserList
	petlist        PetList
	chatlog        ChatLog

	//Store Maps
	StubStorage = stubMapping{}
)
