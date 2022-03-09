package main

import (
	"encoding/json"
	"fmt"
)

func JokesAPI(msg string) []string {
	jokejson := jokes{}
	var retuarr []string

	err := json.Unmarshal([]byte(msg), &jokejson)
	if err == nil {
		fmt.Print(err)
	}

	if jokejson.JOKE.BOFHline != "" {
		retuarr = append(retuarr, jokejson.JOKE.BOFHline)
	} else {
		retuarr = append(retuarr, jokejson.JOKE.Q)
		if jokejson.JOKE.A != "" {
			retuarr = append(retuarr, jokejson.JOKE.A)
		}
	}
	return retuarr
}
