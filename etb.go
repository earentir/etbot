package main

import (
	"fmt"
	"time"
)

func main() {

	LoadJSONFileTOStruct("etb-settings.json", &settings)

	etb := BasicBot{
		Channel:     settings.General.Twitch.Channel,
		Name:        settings.General.Twitch.BotUserName,
		Port:        settings.General.Twitch.IRCPort,
		PrivatePath: settings.General.CredentialFile,
		Server:      settings.General.Twitch.Server,
		MsgRate:     time.Duration(settings.General.Twitch.MSGRate),
	}

	LoadJSONFileTOStruct(etb.PrivatePath, &creds)

	fmt.Println(jokesJSON("oil"))

	if settings.Servers.BotServers.Chat {
		etb.Start()
	}
}
