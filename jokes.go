package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type bofh struct {
	JOKE struct {
		BOFHline string `json:"bofhline"`
	} `json:"joke"`
}

type joke struct {
	JOKE struct {
		Q string `json:"q"`
		A string `json:"a"`
	} `json:"joke"`
}

func JokesBOFH(msg string) string {

	bofhjson := bofh{}

	err := json.Unmarshal([]byte(msg), &bofhjson)
	if err != nil {
		fmt.Print(err)
	}

	return bofhjson.JOKE.BOFHline
}

func JokesJoke(msg string) string {

	jokejson := joke{}

	err := json.Unmarshal([]byte(msg), &jokejson)
	if err != nil {
		fmt.Print(err)
	}

	if jokejson.JOKE.A != "0" {
		return jokejson.JOKE.Q + " " + jokejson.JOKE.A
	} else {
		return jokejson.JOKE.Q
	}
}

func oliveoil() string {
	jokefile, err := os.Open("oliveoillines.txt")
	if err != nil {
		fmt.Print(err)
	}

	defer jokefile.Close()

	var lines []string

	scanner := bufio.NewScanner(jokefile)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	rand.Seed(time.Now().UnixNano())
	return lines[rand.Intn(len(lines))]
}
