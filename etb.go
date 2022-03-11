package main

import (
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	cli "github.com/jawher/mow.cli"
)

func main() {

	app := cli.App("etbot", "Twitch & Discord & OBS Bot")
	app.Version("v", etbver)
	// app.Spec = ""

	app.Command("twitchbot", "Manage Twitch Bot Funtions", func(twitchbot *cli.Cmd) {
		twitchbot.Command("start", "Start Bot", cliTBStart)
	})

	if _, err := os.Stat("settings/etb-settings.json"); err == nil {

		LoadJSONFileTOStruct("settings/etb-settings.json", &settings)

		etb := BasicBot{
			Channel:     settings.General.Twitch.Channel,
			Name:        settings.General.Twitch.BotUserName,
			Port:        settings.General.Twitch.IRCPort,
			PrivatePath: settings.General.CredentialFile,
			Server:      settings.General.Twitch.Server,
			MsgRate:     time.Duration(settings.General.Twitch.MSGRate),
		}

		if _, err := os.Stat("settings/systemcommands.json"); err == nil {
			LoadJSONFileTOStruct("settings/systemcommands.json", &systemcommands)
		}

		if _, err := os.Stat("settings/etb-users.json"); err == nil {
			LoadJSONFileTOStruct("settings/etb-users.json", &userlist)
		}

		if _, err := os.Stat("settings/usr-cmd.json"); err == nil {
			LoadJSONFileTOStruct("settings/usr-cmd.json", &usercommands)
		}

		if _, err := os.Stat("settings/pets.json"); err == nil {
			LoadJSONFileTOStruct("settings/pets.json", &petlist)
		}

		if _, err := os.Stat(etb.PrivatePath); err == nil {
			LoadJSONFileTOStruct(etb.PrivatePath, &creds)

			chatlog.Channel = settings.General.Twitch.Channel
			chatlog.Date = strconv.Itoa(int(time.Now().Unix()))

			//ffs we catch oob interupt, cause he keeps ctrl+c
			c := make(chan os.Signal)
			signal.Notify(c, os.Interrupt, syscall.SIGTERM)
			go func() {
				<-c
				cleanup()
				os.Exit(1)
			}()

			if settings.Servers.WebServers.Enabled {
				go startWebServer()
			}

			if settings.Servers.BotServers.Chat {
				etb.Start()
			}
		} else {
			saveSettings()
		}

	} else {
		saveSettings()
	}
}
