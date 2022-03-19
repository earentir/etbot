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
		loadData([]string{"FilePaths", "Settings"}, settings)

		etb := BasicBot{
			Channel:     settings.General.Twitch.Channel,
			Name:        settings.General.Twitch.BotUserName,
			Port:        settings.General.Twitch.IRCPort,
			PrivatePath: settings.General.CredentialFile,
			Server:      settings.General.Twitch.Server,
			MsgRate:     time.Duration(settings.General.Twitch.MSGRate),
		}

		if !checkLoadStatus() {
			return
		}

		loadData([]string{"FilePaths", "Settings"}, systemcommands)
		loadData([]string{"FilePaths", "Users"}, userlist)
		loadData([]string{"FilePaths", "UserCommands"}, usercommands)
		loadData([]string{"FilePaths", "Pets"}, petlist)
		loadData([]string{"FilePaths", "CredentialFile"}, creds)

		if _, err := os.Stat(etb.PrivatePath); err == nil {
			loadData([]string{"FilePaths", "CredentialFile"}, creds)

			chatlog.Channel = settings.General.Twitch.Channel
			chatlog.Date = strconv.Itoa(int(time.Now().Unix()))

			//ffs we catch oob interupt, cause the noop keeps ctrl+c
			c := make(chan os.Signal)
			signal.Notify(c, os.Interrupt, syscall.SIGTERM)
			go func() {
				<-c
				cleanup() //we cleanup here
				os.Exit(1)
			}()

			saveData([]string{"FilePaths", "Settings"}, settings)

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
