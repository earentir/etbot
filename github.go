package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"regexp"
	"strings"
)

func cmdFr(bb *BasicBot, userName, cmd, msg string) {
	var title string
	if !isCMD(cmd, msg) {
		title = getCleanMessage(msg)
		cmdres := exec.Command("gh", "issue", "create", fmt.Sprintf("-t %s from %s", title, userName), "-b \"\" ", "-lchat-bot")

		var out, errbuf bytes.Buffer

		cmdres.Stderr = &errbuf
		cmdres.Stdout = &out

		cmdres.Run()
		if len(errbuf.String()) == 0 {
			msgOut := fmt.Sprintf("Feature Request Ticket Opened by %s with a title of \"%s\"", userName, title)
			botSay(bb, msgOut)
		}
		fmt.Println(errbuf.String())
	} else {
		cmdFrList(bb, userName, cmd, msg)
	}
}

func cmdFrList(bb *BasicBot, userName, cmd, msg string) {
	cmdres := exec.Command("gh", "issue", "list")

	var errbuf bytes.Buffer
	var frs []string
	cmdres.Stderr = &errbuf

	out, err := cmdres.Output()
	if err == nil {
		msgOut := fmt.Sprintf("Feature Request List, Opened by %s", userName)
		botSay(bb, msgOut)
		outputLines := strings.Split(string(out), "\n")
		for _, j := range outputLines {
			if strings.Contains(j, "chat-bot") {
				if strings.Contains(j, userName) {
					str := regexp.MustCompile("\t").Split(j, -1)
					date := strings.Fields(str[4])
					line := fmt.Sprintf("%s %s on %s %s", str[0], strings.ReplaceAll(str[2], fmt.Sprintf("from %s", userName), ""), date[0], date[1])
					frs = append(frs, line)
				}
			}
		}
		if len(frs) > 0 {
			for _, j := range frs {
				botSay(bb, j)
			}
		} else {
			msgOut := "Please type your feature request after the cmd, ex: !fr add more stuff FFS"
			botSay(bb, msgOut)
		}
	}
}

func cmdGithub(bb *BasicBot, userName, cmd, msg string) {
	var msgOut string
	if isCMD(cmd, msg) {
		msgOut = getGists(settings.General.Twitch.Channel)
	} else {
		if isAttr(msg) {
			msgOut = getGists(getAttributedUser(msg, false))
		}
	}
	botSay(bb, msgOut)
}

func getGists(ghuser string) string {
	// Make the HTTP request
	var gisturl string = fmt.Sprintf("https://api.github.com/users/%s/gists", ghuser)
	resp, err := http.Get(gisturl)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	// Decode the response
	var gists []Gist
	var msgOut string
	if err := json.NewDecoder(resp.Body).Decode(&gists); err != nil {
		fmt.Println(err)
	}

	var sep string
	if len(gists) > 0 {
		for _, gist := range gists {
			if gist.Description != "" {
				msgOut += fmt.Sprintf("%s, ", gist.Description)
			}
			msgOut += "Files: "
			var iterator int
			for _, gfile := range gist.Files {
				iterator++
				if len(gist.Files) > 1 && iterator < len(gist.Files) {
					sep = ", "
				} else {
					sep = ""
				}

				msgOut += fmt.Sprintf("%s%s", gfile.Filename, sep)
			}
			msgOut += " | "
		}

		msgOut += fmt.Sprintf("https://gist.github.com/%s", ghuser)
	} else {
		msgOut = "No Gists"
	}

	return msgOut
}
