package main

type Settings struct {
	Acl struct {
		Users []User `json:"users"`
	} `json:"acl"`
	Commands []Command `json:"commands"`
	Lurklist []struct {
		Username string `json:"username"`
		Lurkedon int    `json:"lurkedon"`
	} `json:"lurklist"`
	General struct {
		Servers struct {
			Spotify struct {
				Currentlyplaying bool `json:"currentlyplaying"`
			} `json:"spotify"`
			Twitch struct {
				Goals bool `json:"goals"`
			} `json:"twitch"`
		} `json:"servers"`
		Weather struct {
			DefaultCity string `json:"defaultcity"`
		} `json:"weather"`
	} `json:"general"`
}

type User struct {
	Users []struct {
		Name    string `json:"name"`
		Nick    string `json:"nick"`
		Type    string `json:"type"`
		Socials struct {
			Twitch string `json:"twitch"`
			Github string `json:"github"`
			Itchio string `json:"itchio"`
		} `json:"socials"`
	} `json:"users"`
}

type Command struct {
	Commands []struct {
		Msg     string `json:"msg"`
		Atmsg   string `json:"atmsg"`
		Help    string `json:"help"`
		Lastuse int    `json:"lastuse"`
	}
}
