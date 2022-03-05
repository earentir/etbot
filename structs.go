package main

import (
	"net"
	"time"
)

type BasicBot struct {

	// The channel that the bot is supposed to join. Note: The name MUST be lowercase, regardless
	// of how the username is displayed on Twitch.tv.
	Channel string

	// A reference to the bot's connection to the server.
	conn net.Conn

	// A forced delay between bot responses. This prevents the bot from breaking the message limit
	// rules. A 20/30 millisecond delay is enough for a non-modded bot. If you decrease the delay
	// make sure you're still within the limit!
	//
	// Message Rate Guidelines: https://dev.twitch.tv/docs/irc#irc-command-and-message-limits
	MsgRate time.Duration

	// The name that the bot will use in the chat that it's attempting to join.
	Name string

	// The port of the IRC server.
	Port string

	// A path to a limited-access directory containing the bot's OAuth credentials.
	PrivatePath string

	// The domain of the IRC server.
	Server string

	// The time at which the bot achieved a connection to the server.
	startTime time.Time
}

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
			Chat         bool `json:"chat"`
			AllowedToSay bool `json:"allowedtosay"`
		} `json:"bot"`
	} `json:"servers"`
	Weather struct {
		DefaultCity string `json:"default,omitempty"`
	} `json:"weather"`
	Curency struct {
		DefaultCurrency string `json:"default,omitempty"`
		CurrencyTo      string `json:"to,omitempty"`
		CryptoDefault   string `json:"cryptodefault,omitempty"`
	} `json:"currency"`
	Users      []User          `json:"users"`
	Commands   []Command       `json:"commands"`
	Lurklists  []LurkerList    `json:"lurklist"`
	UserLevels []UserLevelList `json:"userlevels"`
	Pets       []Pet           `json:"pets"`
	Quotes     []QuoteList     `json:"quotes"`
}

type UserCommands []struct {
	UserCmdName    string        `json:"name"`
	UserCmdType    string        `json:"type"`
	Messages       []string      `json:"messages"`
	Alias          []string      `json:"alias"`
	UserCmdOptions CommandOption `json:"options"`
}

type QuoteList struct {
	QuotedMessage string `json:"quotedmessage"`
	QuoteDate     int64  `json:"quotedate"`
	AtributedUser string `json:"attibuteduser"`
	Quoter        string `json:"quoter"`
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
	Socials []Social `json:"socials"`
}

type Social struct {
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
	Lurker      string `json:"lurker"`
	LurkedOn    int    `json:"lurkedon,omitempty"`
	LurkMessage string `json:"lurkmessage"`
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

type TMDBNetworks struct {
	ID            int64  `json:"ID"`
	Name          string `json:"Name"`
	LogoPath      string `json:"logo_path"`
	OriginCountry string `json:"origin_country"`
}

type TMDBTVSeasons struct {
	ID           int64  `json:"ID"`
	Name         string `json:"Name"`
	Overview     string `json:"Overview"`
	AirDate      string `json:"air_date"`
	EpisodeCount int64  `json:"episode_count"`
	PosterPath   string `json:"poster_path"`
	SeasonNumber int64  `json:"season_number"`
}

type TMDBProductionCompanies struct {
	ID            int64  `json:"ID"`
	Name          string `json:"Name"`
	LogoPath      string `json:"logo_path"`
	OriginCountry string `json:"origin_country"`
}

type TMDBTV struct {
	CreatedBy           interface{}               `json:"CreatedBy"`
	Genres              []TMDBGenres              `json:"Genres"`
	Homepage            string                    `json:"Homepage"`
	ID                  int                       `json:"ID"`
	Languages           []string                  `json:"Languages"`
	Name                string                    `json:"Name"`
	Networks            []TMDBNetworks            `json:"Networks"`
	Overview            string                    `json:"Overview"`
	Popularity          float64                   `json:"Popularity"`
	Seasons             []TMDBTVSeasons           `json:"Seasons"`
	Status              string                    `json:"Status"`
	Type                string                    `json:"Type"`
	BackdropPath        string                    `json:"backdrop_path"`
	EpisodeRunTime      []int                     `json:"episode_run_time"`
	FirstAirDate        string                    `json:"first_air_date"`
	InProduction        bool                      `json:"in_production"`
	LastAirDate         string                    `json:"last_air_date"`
	NumberOfEpisodes    int                       `json:"number_of_episodes"`
	NumberOfSeasons     int                       `json:"number_of_seasons"`
	OriginCountry       []string                  `json:"origin_country"`
	OriginalLanguage    string                    `json:"original_language"`
	OriginalName        string                    `json:"original_name"`
	PosterPath          string                    `json:"poster_path"`
	ProductionCompanies []TMDBProductionCompanies `json:"production_companies"`
	VoteAverage         float64                   `json:"vote_average"`
	VoteCount           int                       `json:"vote_count"`
}

type BinanceData struct {
	AskPrice           string `json:"askPrice"`
	AskQty             string `json:"askQty"`
	BidPrice           string `json:"bidPrice"`
	BidQty             string `json:"bidQty"`
	CloseTime          int    `json:"closeTime"`
	Count              int    `json:"count"`
	FirstID            int    `json:"firstId"`
	HighPrice          string `json:"highPrice"`
	LastID             int    `json:"lastId"`
	LastPrice          string `json:"lastPrice"`
	LastQty            string `json:"lastQty"`
	LowPrice           string `json:"lowPrice"`
	OpenPrice          string `json:"openPrice"`
	OpenTime           int    `json:"openTime"`
	PrevClosePrice     string `json:"prevClosePrice"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	QuoteVolume        string `json:"quoteVolume"`
	Symbol             string `json:"symbol"`
	Volume             string `json:"volume"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
}
