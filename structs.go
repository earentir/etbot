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
	} `json:"general"`
	Users      []User          `json:"users"`
	Commands   []Command       `json:"commands"`
	Lurklists  []LurkList      `json:"lurklist"`
	UserLevels []UserLevelList `json:"userlevels"`
}

type User struct {
	Name    string   `json:"name"`
	Nick    string   `json:"nick,omitempty"`
	Type    string   `json:"type"`
	Socials []Social `json:"socials"`
}

type Command struct {
	CommandName    string        `json:"name"`
	CommandOptions CommandOption `json:"options"`
}

type Social struct {
	Twitch     string `json:"twitch,omitempty"`
	Github     string `json:"github,omitempty"`
	Itchio     string `json:"itchio,omitempty"`
	Youtube    string `json:"youtube,omitempty"`
	Twitter    string `json:"twiter,omitempty"`
	Instagram  string `json:"instagram,omitempty"`
	Facebook   string `json:"facebook,omitempty"`
	Artstation string `json:"artstation,omitempty"`
}

type CommandOption struct {
	Msg                string `json:"msg,omitempty"`
	Atmsg              string `json:"atmsg,omitempty"`
	Help               string `json:"help,omitempty"`
	Alias              string `json:"alias,omitempty"`
	Lastuse            int    `json:"lastuse,omitempty"`
	Counter            int    `json:"counter,omitempty"`
	UserLevel          int    `json:"userlevel"`
	CooldownMultiplier int    `json:"cooldownmultiplier"`
	Enabled            bool   `json:"enabled"`
}

type UserLevelList struct {
	Level    int    `json:"level"`
	Name     string `json:"name"`
	Cooldown int    `json:"cooldown"`
}

type LurkList struct {
	Lurkers []struct {
		Lurker   string `json:"lurker"`
		LurkedOn int64  `json:"lurkedon,omitempty"`
	} `json:"lurklist"`
}
