package main

import (
	"os"
	"os/signal"
	"path"
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

	if !checkLoadStatus() {
		return
	}

	if _, err := os.Stat("settings/settings.json"); err == nil {
		LoadJSONFileTOStruct("settings/settings.json", &settings)

		etb := BasicBot{
			Channel:     settings.General.Twitch.Channel,
			Name:        settings.General.Twitch.BotUserName,
			Port:        settings.General.Twitch.IRCPort,
			PrivatePath: settings.General.CredentialFile,
			Server:      settings.General.Twitch.Server,
			MsgRate:     time.Duration(settings.General.Twitch.MSGRate),
		}

		loadData("SystemCommands", &systemcommands)
		loadData("Users", &userlist)
		loadData("UserCommands", &usercommands)
		loadData("Pets", &petlist)

		if _, err := os.Stat(path.Join(settings.FilePaths.SettingsDir, settings.FilePaths.CredentialFile)); err == nil {
			loadData("CredentialFile", &creds)

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

			saveData("Settings", settings)

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
