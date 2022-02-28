package main

import (
	"fmt"
	"strings"

	"github.com/agnivade/levenshtein"
	"github.com/ryanbradynd05/go-tmdb"
)

func tmdbSearch(mediaTitle string) TMDBSearch {
	var (
		tmdbAPI           *tmdb.TMDb
		searchresults     TMDBSearchResults
		search, jsonreply TMDBSearch
		cannonical        string
		found             bool = false
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

	LoadJSONTOStruct([]byte(mediaInfoJSON), &jsonreply)

	for _, j := range jsonreply.Results {
		cannonical = getUniqueMediaTitle(j.Title, j.OriginalTitle, j.Name, j.OriginalName)

		mediaFound, commonWords := checkAllstrings(cannonical, strings.Fields(mediaTitle))
		if (mediaFound && commonWords >= len(strings.Fields(cannonical))) || levenshtein.ComputeDistance(strings.ToLower(cannonical), strings.ToLower(mediaTitle)) < 3 {
			found = true

			j.Overview = limitOverview(j.Overview)

			switch j.MediaType {
			case "movie":
				search.Results = append(search.Results, j)
				fmt.Println("found mv", mediaFound, commonWords, cannonical, levenshtein.ComputeDistance(strings.ToLower(cannonical), strings.ToLower(mediaTitle)))
			case "tv":
				search.Results = append(search.Results, j)
				fmt.Println("found tv", mediaFound, commonWords, cannonical, levenshtein.ComputeDistance(strings.ToLower(cannonical), strings.ToLower(mediaTitle)))
			}
		}
	}

	var maxresults int = 0
	if len(jsonreply.Results) > 3 {
		maxresults = 3
	} else {
		maxresults = len(jsonreply.Results)
	}

	if !found {
		for i := 0; i < maxresults; i++ {
			var relDate string = ""

			cannonical = getUniqueMediaTitle(jsonreply.Results[i].Title, jsonreply.Results[i].OriginalTitle, jsonreply.Results[i].Name, jsonreply.Results[i].OriginalName)
			fmt.Println("not found accurate ", cannonical)
			if jsonreply.Results[i].FirstAirDate != "" {
				relDate = jsonreply.Results[i].FirstAirDate
			} else {
				relDate = jsonreply.Results[i].ReleaseDate
			}

			searchresults.Title = cannonical
			searchresults.ID = jsonreply.Results[i].ID
			searchresults.FirstAirDate = relDate
			searchresults.MediaType = jsonreply.Results[i].MediaType
			searchresults.Overview = limitOverview(jsonreply.Results[i].Overview)

			search.Results = append(search.Results, searchresults)
		}
	}

	return search
}

func checkAllstrings(str string, subs []string) (bool, int) {

	matches := 0
	isCompleteMatch := true

	for _, sub := range subs {
		if strings.Contains(strings.ToLower(str), strings.ToLower(sub)) {
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

func tmdbMovie(mediaID int) TMDBMovie {
	var (
		tmdbAPI       *tmdb.TMDb
		tmdbmovieData TMDBMovie
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

	return tmdbmovieData
}

func tmdbTV(mediaID int) TMDBTV {
	var (
		tmdbAPI    *tmdb.TMDb
		tmdbtvData TMDBTV
	)

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

	LoadJSONTOStruct([]byte(mediaInfoJSON), &tmdbtvData)

	return tmdbtvData
}
