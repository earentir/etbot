package main

type Settings struct {
	Acl struct {
		Users []User `json:"users"`
	} `json:"acl"`
	Commands []Command `json:"commands"`
	Lurklist []struct {
		Username string `json:"username"`
		Lurkedon int    `json:"lurkedon,omitempty"`
	} `json:"lurklist"`
	General struct {
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
	} `json:"general"`
}

type User struct {
	Users []struct {
		Name    string `json:"name"`
		Nick    string `json:"nick,omitempty"`
		Type    string `json:"type"`
		Socials struct {
			Twitch     string `json:"twitch,omitempty"`
			Github     string `json:"github,omitempty"`
			Itchio     string `json:"itchio,omitempty"`
			Youtube    string `json:"youtube,omitempty"`
			Twitter    string `json:"twiter,omitempty"`
			Instagram  string `json:"instagram,omitempty"`
			Facebook   string `json:"facebook,omitempty"`
			Artstation string `json:"artstation,omitempty"`
		} `json:"socials"`
	} `json:"users"`
}

type Command struct {
	Commands []struct {
		Msg     string `json:"msg,omitempty"`
		Atmsg   string `json:"atmsg,omitempty"`
		Help    string `json:"help,omitempty"`
		Lastuse int    `json:"lastuse,omitempty"`
	}
}
