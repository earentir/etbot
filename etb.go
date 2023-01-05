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

	app.Command("twitchbot", "Manage Twitch Bot Funtions", func(twitchbot *cli.Cmd) {
		twitchbot.Command("start", "Start Bot", cliTBStart)
	})

	// app.Run(os.Args)

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

type GithubRepoIssues []struct {
	URL           string `json:"url"`
	RepositoryURL string `json:"repository_url"`
	LabelsURL     string `json:"labels_url"`
	CommentsURL   string `json:"comments_url"`
	EventsURL     string `json:"events_url"`
	HTMLURL       string `json:"html_url"`
	ID            int    `json:"id"`
	NodeID        string `json:"node_id"`
	Number        int    `json:"number"`
	Title         string `json:"title"`
	User          struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"user"`
	Labels            []interface{} `json:"labels"`
	State             string        `json:"state"`
	Locked            bool          `json:"locked"`
	Assignee          interface{}   `json:"assignee"`
	Assignees         []interface{} `json:"assignees"`
	Milestone         interface{}   `json:"milestone"`
	Comments          int           `json:"comments"`
	CreatedAt         time.Time     `json:"created_at"`
	UpdatedAt         time.Time     `json:"updated_at"`
	ClosedAt          interface{}   `json:"closed_at"`
	AuthorAssociation string        `json:"author_association"`
	ActiveLockReason  interface{}   `json:"active_lock_reason"`
	Body              interface{}   `json:"body"`
	Reactions         struct {
		URL        string `json:"url"`
		TotalCount int    `json:"total_count"`
		Num1       int    `json:"+1"`
		Num10      int    `json:"-1"`
		Laugh      int    `json:"laugh"`
		Hooray     int    `json:"hooray"`
		Confused   int    `json:"confused"`
		Heart      int    `json:"heart"`
		Rocket     int    `json:"rocket"`
		Eyes       int    `json:"eyes"`
	} `json:"reactions"`
	TimelineURL           string      `json:"timeline_url"`
	PerformedViaGithubApp interface{} `json:"performed_via_github_app"`
	StateReason           interface{} `json:"state_reason"`
}
