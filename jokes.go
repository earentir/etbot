package main

import (
	"encoding/json"
	"fmt"
)

type jokes struct {
	JOKE struct {
		BOFHline string `json:"bofhline,omitempty"`
		Q        string `json:"q,omitempty"`
		A        string `json:"a,omitempty"`
	} `json:"joke"`
}

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
