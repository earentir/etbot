package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
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

	retuarr = append(retuarr, jokejson.JOKE.BOFHline)
	retuarr = append(retuarr, jokejson.JOKE.Q)
	retuarr = append(retuarr, jokejson.JOKE.A)

	return retuarr
}

func oliveoil() string {
	rand.Seed(time.Now().UnixNano())
	lines := readTextFile("oliveoillines.txt")
	return lines[rand.Intn(len(lines))]
}

func yogurt() string {
	rand.Seed(time.Now().UnixNano())
	lines := readTextFile("yogurt.txt")
	ouryogurt := lines[rand.Intn(len(lines))]
	if ouryogurt == "Greek" {
		return ouryogurt + " yogurt is a good Yogurt"
	} else {
		return ouryogurt + " yogurt is disgusting"
	}
}
