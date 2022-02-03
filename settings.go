package main

import (
	"fmt"
	"strings"
)

type Settings struct {
	Acl struct {
		Users []User `json:"users"`
	} `json:"acl"`
	Commands []Command  `json:"commands"`
	Lurklist []LurkList `json:"lurklist"`
	General  struct {
		Servers struct {
			Spotify struct {
				Currentlyplaying bool `json:"currentlyplaying,omitempty"`
			} `json:"spotify"`
			Twitch struct {
				Goals bool `json:"goals,omitempty"`
			} `json:"twitch"`
		} `json:"servers"`
		Weather struct {
			DefaultCity string `json:"default,omitempty"`
		} `json:"weather"`
		Curency struct {
			DefaultCurrency string `json:"default,omitempty"`
			CurrencyTo      string `json:"to,omitempty"`
		} `json:"currency"`
		UserLevels []UserLevelList `json:"userlevels"`
	} `json:"general"`
}

type UserLevelList struct {
	// Level []struct {
	Level    int    `json:"level"`
	Name     string `json:"name"`
	Cooldown int    `json:"cooldown"`
	// } `json:"userlevels"`
}

type LurkList struct {
	Lurkers []struct {
		Lurker   string `json:"lurker"`
		LurkedOn int64  `json:"lurkedon,omitempty"`
	} `json:"lurklist"`
}

type User struct {
	Name    string   `json:"name"`
	Nick    string   `json:"nick,omitempty"`
	Type    string   `json:"type"`
	Socials []Social `json:"socials"`
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

type Command struct {
	Commands []struct {
		Msg                string `json:"msg,omitempty"`
		Atmsg              string `json:"atmsg,omitempty"`
		Help               string `json:"help,omitempty"`
		Alias              string `json:"alias,omitempty"`
		Lastuse            int    `json:"lastuse,omitempty"`
		Counter            int    `json:"counter,omitempty"`
		UserLevel          int    `json:"userlevel"`
		CooldownMultiplier int    `json:"cooldownmultiplier"`
		Enabled            bool   `json:"enabled,omitempty"`
	} `json:"commands"`
}

func getUserSocials(userName string) []string {
	setusers := settings.Acl.Users
	found := []string{}

	for _, usr := range setusers {
		if usr.Name == userName {
			usrsoc := usr.Socials
			for _, soc := range usrsoc {
				found = append(found, strings.ReplaceAll(fmt.Sprintf("%v", soc), " ", ""))
			}
		}
	}

	return found
}

func SearchUser(userName string) bool {
	var found bool = false

	setusers := settings.Acl.Users
	for _, usr := range setusers {
		if userName == usr.Name {
			found = true
		}
	}
	return found
}

func convertLevelType(LevelType string) int {
	var level int = 10
	setusers := settings.General.UserLevels
	for _, lvl := range setusers {
		if LevelType == lvl.Name {
			level = lvl.Level
		}
	}
	return level
}

func getLevelCoolDown(level int) int {
	var cooldown int = 30000
	setusers := settings.General.UserLevels
	for _, lvl := range setusers {
		if level == lvl.Level {
			cooldown = lvl.Cooldown
		}
	}
	return cooldown
}

func UserLevel(userName string) int {
	var level int = 10
	setusers := settings.Acl.Users
	for _, usr := range setusers {
		if userName == usr.Name {
			level = convertLevelType(usr.Type)
		}
	}
	return level
}
