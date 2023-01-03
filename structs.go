package main

import (
	"net"
	"time"
)

type BasicBot struct {
	Channel     string        // The channel that the bot is supposed to join. Note: The name MUST be lowercase
	conn        net.Conn      // A reference to the bot's connection to the server.
	MsgRate     time.Duration // Message Rate Guidelines: https://dev.twitch.tv/docs/irc#irc-command-and-message-limits
	Name        string        // The bot name
	Port        string        // The port of the IRC server.
	Server      string        // The domain of the IRC server.
	PrivatePath string        // credentials path (to be replaced with the Filepaths entry)
	startTime   time.Time     // Time when we connected
}

// Settings struct
type Settings struct {
	General struct {
		Twitch struct {
			Channel     string `json:"channel"`
			BotUserName string `json:"botusername"`
			IRCPort     string `json:"ircport"`
			Server      string `json:"server"`
			MSGRate     int    `json:"msgrate"`
		} `json:"twitch"`
		Project struct {
			Description string `json:"description,omitempty"`
		} `json:"project"`
		Lockdown struct {
			Enabled bool   `json:"enabled"`
			Reason  string `json:"reason,omitempty"`
		} `json:"lockdown"`
	} `json:"general"`
	FilePaths struct {
		SettingsDir    string `json:"settingsdir"`
		ChatLogDir     string `json:"chatlogdir"`
		Settings       string `json:"settings"`
		Users          string `json:"users"`
		Jokes          string `json:"jokes"`
		Lurkers        string `json:"lurkers"`
		Pets           string `json:"pets"`
		Quotes         string `json:"quotes"`
		SystemCommands string `json:"system-commands"`
		UserCommands   string `json:"user-commands"`
		// CredentialFile string `json:"credentials"`
	} `json:"filepaths"`
	Servers struct {
		WebServers struct {
			Enabled bool `json:"enabled"`
			Spotify struct {
				Currentlyplaying bool `json:"currentlyplaying,omitempty"`
			} `json:"spotify"`
			Twitch struct {
				Goals bool `json:"goals,omitempty"`
			} `json:"twitch"`
		} `json:"web"`
		BotServers struct {
			Chat         bool `json:"chat"`
			SendMessages bool `json:"sendmessages"`
			Log          bool `json:"log"`
		} `json:"bot"`
		DiscordBot struct {
			Enabled      bool `json:"enabled"`
			SendMessages bool `json:"sendmessages"`
			Log          bool `json:"log"`
		}
	} `json:"servers"`
	API struct {
		Weather struct {
			DefaultCity string `json:"default,omitempty"`
		} `json:"weather"`
		Curency struct {
			DefaultCurrency string `json:"default,omitempty"`
			CurrencyTo      string `json:"to,omitempty"`
			CryptoDefault   string `json:"cryptodefault,omitempty"`
		} `json:"currency"`
		Calendar struct {
			Country   string `json:"country"`
			DaysAhead int    `josn:"daysahead"`
		} `json:"Calendar"`
		OpenAI struct {
			Model            string  `json:"model"`
			Prompt           string  `json:"prompt"`
			Temperature      float64 `json:"temperature"`
			Max_Tokens       int     `json:"max_tokens"`
			Top_P            float64 `json:"top_p"`
			FrequencyPenalty float64 `json:"frequency_penalty"`
			PresencePenalty  float64 `json:"presence_penalty"`
			Stop             string  `json:"stop"`
		} `json:"OpenAI"`
	}
	UserLevels []UserLevelList `json:"userlevels"`
}

// Credentials struct
type Credentials struct {
	TwitchPassword     string `json:"twitch_password"`
	TwitchAppToken     string `json:"twitch_apptoken"`
	TwitchClientID     string `json:"twitch_client_id"`
	TwitchClientsecret string `json:"twitch_client_secret"`
	OpenWeatherAPIKey  string `json:"openweathermapapi,omitempty"`
	CurrencyAPIKey     string `json:"currconvapi,omitempty"`
	TMDBToken          string `json:"tmdb_token,omitempty"`
	Calendarific       string `json:"calendarific_token,omitempty"`
	OpenAI             string `json:"openai_token,omitempty"`
	Gumroad            string `json:"gumroad_access_token,omitempty"`
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

type PetList struct {
	Pets []Pet
}

type Pet struct {
	Name      string `json:"name,omitempty"`
	Pet       int    `json:"pet,omitempty"`
	Feed      int    `json:"feed,omitempty"`
	FeedLimit int    `json:"feedlimit,omitempty"`
}

type UserList struct {
	Users []User `json:"users"`
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

type LurkList struct {
	Lurkers []Lurker `json:"lurklist"`
}

type Lurker struct {
	Name    string `json:"name"`
	Date    int64  `json:"date,omitempty"`
	Message string `json:"message"`
}

type QuoteList struct {
	QuoteItems []QuoteItem `json:"quotes"`
}

type QuoteItem struct {
	QuotedMessage string `json:"quotedmessage"`
	QuoteDate     int64  `json:"quotedate"`
	AtributedUser string `json:"attibuteduser"`
	Quoter        string `json:"quoter"`
}

type JokeList struct {
	JokeItems []JokeItem `json:"jokes"`
}

type JokeItem struct {
	JokeMessage   string `json:"jokemessage"`
	JokeDate      int64  `json:"jokedate"`
	AtributedUser string `json:"attibuteduser"`
	Jokster       string `json:"jokster"`
}

type UserCommands []struct {
	UserCmdName    string        `json:"name"`
	UserCmdType    string        `json:"type"`
	Messages       []string      `json:"messages"`
	Alias          []string      `json:"alias"`
	UserCmdOptions CommandOption `json:"options"`
}

type CommandList struct {
	Commands []Command `json:"commands"`
}

type Command struct {
	Name         string        `json:"name"`
	InternalName string        `json:"internalname"`
	Alias        []string      `json:"alias"`
	Options      CommandOption `json:"options"`
}

type CommandOption struct {
	Msg       string `json:"msg,omitempty"`
	Atmsg     string `json:"atmsg,omitempty"`
	Help      string `json:"help,omitempty"`
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

type openweathermap struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

type jokes struct {
	JOKE struct {
		BOFHline string `json:"bofhline,omitempty"`
		Q        string `json:"q,omitempty"`
		A        string `json:"a,omitempty"`
	} `json:"joke"`
}

type DaysOfF struct {
	Meta struct {
		Code int `json:"code"`
	} `json:"meta"`
	Response struct {
		Holidays []struct {
			Country struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"country"`
			Date struct {
				Datetime struct {
					Day    int `json:"day"`
					Hour   int `json:"hour"`
					Minute int `json:"minute"`
					Month  int `json:"month"`
					Second int `json:"second"`
					Year   int `json:"year"`
				} `json:"datetime"`
				Iso      string `json:"iso"`
				Timezone struct {
					Offset          string `json:"offset"`
					Zoneabb         string `json:"zoneabb"`
					Zonedst         int    `json:"zonedst"`
					Zoneoffset      int    `json:"zoneoffset"`
					Zonetotaloffset int    `json:"zonetotaloffset"`
				} `json:"timezone"`
			} `json:"date"`
			Description string   `json:"description"`
			Locations   string   `json:"locations"`
			Name        string   `json:"name"`
			Type        []string `json:"type"`
			//States      string   `json:"states"`
		} `json:"holidays"`
	} `json:"response"`
}

type ChatLog struct {
	Channel       string        `json:"channel"`
	BroadcasterID string        `json:"broadcaster_id"`
	Date          string        `json:"date"`
	GameID        string        `json:"game_id"`
	GameName      string        `json:"game_name"`
	StreamTitle   string        `json:"title"`
	ChatMessages  []ChatMessage `json:"chatmessages"`
}

type ChatMessage struct {
	Date    string `json:"date"`
	User    string `json:"user"`
	Message string `json:"message"`
}
