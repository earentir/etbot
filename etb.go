package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
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

	// Setup the bot settings
	etb := BasicBot{
		Channel: settings.General.Twitch.Channel,
		Name:    settings.General.Twitch.BotUserName,
		Port:    settings.General.Twitch.IRCPort,
		Server:  settings.General.Twitch.Server,
		MsgRate: time.Duration(settings.General.Twitch.MSGRate),
	}

	//Load the data from the json files
	loadData("SystemCommands", &systemcommands)
	loadData("Users", &userlist)
	loadData("UserCommands", &usercommands)
	loadData("Pets", &petlist)

	// get the channel info
	var twitchChannelData TwitchChannelData = getChannelInfo(getTwitchUser(strings.ToLower(settings.General.Twitch.Channel))[0].ID)

	// Setup chatlog and read channel data
	chatlog = ChatLog{
		Channel:       settings.General.Twitch.Channel,
		BroadcasterID: twitchChannelData[0].BroadcasterID,
		Date:          strconv.Itoa(int(time.Now().Unix())),
		GameID:        twitchChannelData[0].GameID,
		GameName:      twitchChannelData[0].GameName,
		StreamTitle:   twitchChannelData[0].Title,
	}

	// ffs we catch oob interupt, cause the noop keeps ctrl+c
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup() // we cleanup here
		os.Exit(1)
	}()

	if settings.Servers.WebServers.Enabled {
		go startWebServer()
	}

	if settings.Servers.BotServers.Chat {
		etb.Start()
	}

}
