package main

import "time"

var creds Credentials
var settings Settings

//Credentials struct
type Credentials struct {
	TwitchPassword     string `json:"twitch_password,omitempty"`
	TwitchAppToken     string `json:"twitch_apptoken,omitempty"`
	TwitchClientID     string `json:"twitch_client_id,omitempty"`
	TwitchClientsecret string `json:"twitch_client_secret,omitempty"`
	OpenWeatherAPIKey  string `json:"openweathermapapi,omitempty"`
	CurrencyAPIKey     string `json:"currconvapi,omitempty"`
	TMDBToken          string `json:"tmdb_token,omitempty"`
}

type TwitchUserData []struct {
	ID              string    `json:"id"`
	Login           string    `json:"login"`
	DisplayName     string    `json:"display_name"`
	Type            string    `json:"type"`
	BroadcasterType string    `json:"broadcaster_type"`
	Description     string    `json:"description"`
	ProfileImageURL string    `json:"profile_image_url"`
	OfflineImageURL string    `json:"offline_image_url"`
	ViewCount       int       `json:"view_count"`
	Email           string    `json:"email"`
	CreatedAt       time.Time `json:"created_at"`
}

type TwitchChannelData []struct {
	BroadcasterID       string `json:"broadcaster_id"`
	BroadcasterName     string `json:"broadcaster_name"`
	BroadcasterLanguage string `json:"broadcaster_language"`
	GameID              string `json:"game_id"`
	GameName            string `json:"game_name"`
	Title               string `json:"title"`
	Delay               int    `json:"delay"`
}

//Settings struct
type Settings struct {
	General struct {
		CredentialFile string `json:"credentialfile"`
		Twitch         struct {
			Channel     string `json:"channel"`
			BotUserName string `json:"botusername"`
			IRCPort     string `json:"ircport"`
			Server      string `json:"server"`
			MSGRate     int    `json:"msgrate"`
		} `json:"twitch"`
		Project struct {
			Description string `json:"description,omitempty"`
		} `json:"project"`
	} `json:"general"`
	Servers struct {
		WebServers struct {
			Spotify struct {
				Currentlyplaying bool `json:"currentlyplaying,omitempty"`
			} `json:"spotify"`
			Twitch struct {
				Goals bool `json:"goals,omitempty"`
			} `json:"twitch"`
		} `json:"web"`
		BotServers struct {
			Chat bool `json:"chat"`
		} `json:"bot"`
	} `json:"servers"`
	Weather struct {
		DefaultCity string `json:"default,omitempty"`
	} `json:"weather"`
	Curency struct {
		DefaultCurrency string `json:"default,omitempty"`
		CurrencyTo      string `json:"to,omitempty"`
	} `json:"currency"`
	Users      []User          `json:"users"`
	Commands   []Command       `json:"commands"`
	Lurklists  []LurkerList    `json:"lurklist"`
	UserLevels []UserLevelList `json:"userlevels"`
	Pets       []Pet           `json:"pets"`
}

type Pet struct {
	Name      string `json:"name,omitempty"`
	Pet       int    `json:"pet,omitempty"`
	Feed      int    `json:"feed,omitempty"`
	FeedLimit int    `json:"feedlimit,omitempty"`
}

type User struct {
	Name    string   `json:"name"`
	Nick    string   `json:"nick,omitempty"`
	Type    string   `json:"type"`
	Love    string   `json:"love"`
	Socials []social `json:"socials"`
}

type social struct {
	SocNet string `json:"socnet"`
	Link   string `json:"link"`
}

type Command struct {
	CommandName    string        `json:"name"`
	CommandOptions CommandOption `json:"options"`
}
type CommandOption struct {
	Msg       string `json:"msg,omitempty"`
	Atmsg     string `json:"atmsg,omitempty"`
	Help      string `json:"help,omitempty"`
	Alias     string `json:"alias,omitempty"`
	Lastuse   int    `json:"lastuse,omitempty"`
	Counter   int    `json:"counter,omitempty"`
	UserLevel int    `json:"userlevel"`
	Cooldown  int    `json:"cooldown"`
	Enabled   bool   `json:"enabled"`
}

type UserLevelList struct {
	Level    int    `json:"level"`
	Name     string `json:"name"`
	Cooldown int    `json:"cooldown"`
}
type LurkerList struct {
	Lurker   string `json:"lurker"`
	LurkedOn int    `json:"lurkedon,omitempty"`
}
