package main

import (
	"bytes"
	"fmt"
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
					line := str[0] + " " + strings.ReplaceAll(str[2], "from "+userName, "") + " on " + date[0] + " " + date[1]
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
