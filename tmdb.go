package main

import (
	"fmt"
	"strings"

	"github.com/ryanbradynd05/go-tmdb"
)

func tmdbSearch(mediaTitle string) MediaDATAResults {
	var (
		tmdbAPI      *tmdb.TMDb
		searchresult TMDBSearch
		foundtitles  MediaDATAResults
		foundtitle   MediaData
		cannonical   string
		relDate      string
		found        bool = false
	)

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

	LoadJSONTOStruct([]byte(mediaInfoJSON), &searchresult)

	for _, j := range searchresult.Results {
		cannonical = getUniqueMediaTitle(j.Title, j.OriginalTitle, j.Name, j.OriginalName)

		if strings.EqualFold(cannonical, mediaTitle) {
			found = true
			switch j.MediaType {
			case "movie":
				foundtitles.MediaDATAResults = append(foundtitles.MediaDATAResults, tmdbMovie(j.ID))
			case "tv":
				// foundtitles = append(foundtitles, tmdbTV(j.ID))
			}
		}
	}

	for i := 0; i < 3; i++ {

		if !found {
			cannonical = getUniqueMediaTitle(searchresult.Results[i].Title, searchresult.Results[i].OriginalTitle, searchresult.Results[i].Name, searchresult.Results[i].OriginalName)

			fmt.Println(checkAllstrings(cannonical, strings.Fields(mediaTitle)))

			foundtitle.Title = cannonical
			foundtitle.ID = searchresult.Results[i].ID
			foundtitle.Date = relDate
			foundtitle.Type = searchresult.Results[i].MediaType
			foundtitle.Overview = searchresult.Results[i].Overview

			if searchresult.Results[i].FirstAirDate != "" {
				relDate = searchresult.Results[i].FirstAirDate
			} else {
				relDate = searchresult.Results[i].ReleaseDate
			}

			foundtitles.MediaDATAResults = append(foundtitles.MediaDATAResults, foundtitle)
		}
	}

	return foundtitles
}

func checkAllstrings(str string, subs []string) (bool, int) {

	matches := 0
	isCompleteMatch := true

	str = strings.ToLower(str)

	// fmt.Printf("Search in: \"%s\", subs: %s\n", str, subs)

	for _, sub := range subs {
		if strings.Contains(str, sub) {
			matches += 1
		} else {
			isCompleteMatch = false
		}
	}

	return isCompleteMatch, matches
}

func getUniqueMediaTitle(comparestrings ...string) string {
	var name string = ""

	for i := 0; i < len(comparestrings); i++ {
		if comparestrings[i] != "" {
			name = comparestrings[i]
		}
	}

	return name
}

func tmdbMovie(mediaID int) MediaData {
	var (
		tmdbAPI       *tmdb.TMDb
		tmdbmovieData TMDBMovie
		mediadata     MediaData
	)

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

	LoadJSONTOStruct([]byte(mediaInfoJSON), &tmdbmovieData)

	mediadata.Adult = tmdbmovieData.Adult
	mediadata.ID = tmdbmovieData.ID
	mediadata.Title = tmdbmovieData.Title
	mediadata.Type = "movie"
	mediadata.Date = tmdbmovieData.ReleaseDate
	mediadata.Runtime = tmdbmovieData.Runtime
	if len(tmdbmovieData.Overview) < 233 {
		mediadata.Overview = tmdbmovieData.Overview
	} else {
		mediadata.Overview = tmdbmovieData.Overview[0:233]
	}
	mediadata.Homepage = tmdbmovieData.Homepage
	mediadata.Votes = int(tmdbmovieData.VoteAverage)
	mediadata.VoteCount = tmdbmovieData.VoteCount

	return mediadata
}

func tmdbTV(mediaID int) string {
	var tmdbAPI *tmdb.TMDb
	// var tmdbtvData TMDBTV

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

	// LoadJSONTOStruct([]byte(mediaInfoJSON), &tmdbtvData)

	return mediaInfoJSON
}
