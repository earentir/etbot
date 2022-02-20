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

type TMDBSearch struct {
	Results      []TMDBSearchResults `json:"results"`
	TotalPages   int                 `json:"total_pages"`
	TotalResults int                 `json:"total_results"`
}

type TMDBSearchResults struct {
	OriginalName     string   `json:"original_name,omitempty"`
	OriginalTitle    string   `json:"original_title,omitempty"`
	Name             string   `json:"Name,omitempty"`
	Title            string   `json:"title,omitempty"`
	OriginalLanguage string   `json:"original_language,omitempty"`
	Overview         string   `json:"overview,omitempty"`
	FirstAirDate     string   `json:"first_air_date,omitempty"`
	ReleaseDate      string   `json:"release_date,omitempty"`
	OriginCountry    []string `json:"origin_country,omitempty"`
	PosterPath       string   `json:"poster_path,omitempty"`
	VoteAverage      float64  `json:"vote_average,omitempty"`
	VoteCount        int      `json:"vote_count,omitempty"`
	ID               int      `json:"ID"`
	MediaType        string   `json:"media_type"`
	Adult            bool     `json:"Adult,omitempty"`
}

type TMDBMovie struct {
	Adult            bool                  `json:"Adult"`
	Budget           int                   `json:"Budget"`
	Genres           []TMDBGenres          `json:"Genres"`
	Homepage         string                `json:"Homepage"`
	ID               int                   `json:"ID"`
	ImdbID           string                `json:"imdb_id"`
	OriginalLanguage string                `json:"original_language"`
	OriginalTitle    string                `json:"original_title"`
	Overview         string                `json:"Overview"`
	PosterPath       string                `json:"poster_path"`
	ReleaseDate      string                `json:"release_date"`
	Revenue          int                   `json:"Revenue"`
	Runtime          int                   `json:"Runtime"`
	SpokenLanguages  []TMDBSpokenLanguages `json:"spoken_languages"`
	Status           string                `json:"Status"`
	Tagline          string                `json:"Tagline"`
	Title            string                `json:"Title"`
	VoteAverage      float64               `json:"vote_average"`
	VoteCount        int                   `json:"vote_count"`
}

type TMDBGenres struct {
	ID   int    `json:"ID"`
	Name string `json:"Name"`
}

type TMDBSpokenLanguages struct {
	Iso6391 string `json:"iso_639_1"`
	Name    string `json:"Name"`
}

type MediaDATAResults struct {
	MediaDATAResults []MediaData
}
