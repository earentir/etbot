package main

import (
	"fmt"
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

	if !checkLoadStatus() {
		fmt.Println("We Cant Launch")
		return
	}

	etb := BasicBot{
		Channel: settings.General.Twitch.Channel,
		Name:    settings.General.Twitch.BotUserName,
		Port:    settings.General.Twitch.IRCPort,
		Server:  settings.General.Twitch.Server,
		MsgRate: time.Duration(settings.General.Twitch.MSGRate),
	}

	loadData("SystemCommands", &systemcommands)
	loadData("Users", &userlist)
	loadData("UserCommands", &usercommands)
	loadData("Pets", &petlist)

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

	if settings.Servers.WebServers.Enabled {
		go startWebServer()
	}

	if settings.Servers.BotServers.Chat {
		etb.Start()
	}

}
