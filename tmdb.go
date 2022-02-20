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

		if j.Title != "" {
			cannonical = j.Title
		} else {
			if j.OriginalTitle != "" {
				cannonical = j.OriginalTitle
			} else {
				if j.Name != "" {
					cannonical = j.Name
				} else {
					if j.OriginalName != "" {
						cannonical = j.OriginalName
					}
				}
			}
		}

		if j.FirstAirDate != "" {
			relDate = j.FirstAirDate
		} else {
			relDate = j.ReleaseDate
		}

		if strings.ToLower(cannonical) == strings.ToLower(mediaTitle) {
			if !j.Adult {
				switch j.MediaType {
				case "movie":
					foundtitles.MediaDATAResults = append(foundtitles.MediaDATAResults, tmdbMovie(j.ID))
					// fmt.Println(tmdbMovie(j.ID))
				case "tv":
					// foundtitles = append(foundtitles, tmdbTV(j.ID))
				}
			} else {
				foundtitle.Adult = true
				foundtitles.MediaDATAResults = append(foundtitles.MediaDATAResults, foundtitle)
			}
		} else {
			foundtitle.Title = cannonical
			foundtitle.ID = j.ID
			foundtitle.Date = relDate
			foundtitle.Type = j.MediaType
			foundtitle.Overview = j.Overview

			foundtitles.MediaDATAResults = append(foundtitles.MediaDATAResults, foundtitle)
			fmt.Printf("%#v\n", foundtitle)
		}
	}

	return foundtitles
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

	mediadata.Title = tmdbmovieData.Title
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
