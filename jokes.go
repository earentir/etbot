package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
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

func jokesJSON(jkcmd string) string {
	var outMessage string = ""
	var jsonMap map[string]map[string][]string

	jsonFileDesc, err := os.Open("jokes.json")
	if err != nil {
		fmt.Println(err)
	}

	jsonData, err := ioutil.ReadAll(jsonFileDesc)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(jsonData, &jsonMap)

	for key0 := range jsonMap {
		if key0 == jkcmd {
			for key1 := range jsonMap[key0] {
				rand.Seed(time.Now().UnixNano())
				if len(jsonMap[key0]) > 1 {
					// fmt.Println(key1)
					//if we get here we need to make logic for this to work

				} else {
					rand.Seed(time.Now().UnixNano())
					outMessage = jsonMap[key0][key1][rand.Intn(len(jsonMap[key0][key1]))]
				}
			}
		}
	}

	return outMessage
}
