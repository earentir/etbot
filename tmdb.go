package main

import (
	"fmt"

	"github.com/ryanbradynd05/go-tmdb"
)

func tmdbSearch(mediaTitle string) {
	var tmdbAPI *tmdb.TMDb

	tmdbconfig := tmdb.Config{
		APIKey:   creds.TMDBToken,
		Proxies:  nil,
		UseProxy: false,
	}

	tmdbAPI = tmdb.Init(tmdbconfig)

	mediaInfo, err := tmdbAPI.SearchMulti(mediaTitle, nil)
	if err != nil {
		fmt.Println(err)
	}
	mediaInfoJSON, err := tmdb.ToJSON(mediaInfo)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(mediaInfoJSON)
}

func tmdbMovie(mediaID int) string {
	var tmdbAPI *tmdb.TMDb

	config := tmdb.Config{
		APIKey:   creds.TMDBToken,
		Proxies:  nil,
		UseProxy: false,
	}

	tmdbAPI = tmdb.Init(config)

	mediaInfo, err := tmdbAPI.GetMovieInfo(mediaID, nil)
	if err != nil {
		fmt.Println(err)
	}

	mediaInfoJSON, err := tmdb.ToJSON(mediaInfo)
	if err != nil {
		fmt.Println(err)
	}

	return mediaInfoJSON
}
func tmdbTV(mediaID int) string {
	var tmdbAPI *tmdb.TMDb

	config := tmdb.Config{
		APIKey:   creds.TMDBToken,
		Proxies:  nil,
		UseProxy: false,
	}

	tmdbAPI = tmdb.Init(config)

	mediaInfo, err := tmdbAPI.GetTvInfo(mediaID, nil)
	if err != nil {
		fmt.Println(err)
	}

	mediaInfoJSON, err := tmdb.ToJSON(mediaInfo)
	if err != nil {
		fmt.Println(err)
	}

	return mediaInfoJSON
}
