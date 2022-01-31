package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func inArray(array []string, lookup string) bool {
	sort.Strings(array)
	i := sort.SearchStrings(array, lookup)
	if i < len(array) && array[i] == lookup {
		return true
	}
	return false
}

func getAttributedUser(msg string) string {
	if strings.Contains(msg, "@") {
		return msg[strings.Index(msg, "@"):]
	} else {
		return ""
	}
}

func getCleanMessage(cmd, msg string) string {
	if strings.Contains(msg, "@") {
		return msg[len(cmd)+1 : strings.Index(msg, "@")-1]
	} else {
		return msg[len(cmd)+1:]
	}
}

func readTextFile(filename string) []string {
	textfile, err := os.Open(filename)
	if err != nil {
		fmt.Print(err)
	}
	defer textfile.Close()

	var lines []string
	scanner := bufio.NewScanner(textfile)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func IsItOnTimeout(command, userName string) bool {
	if inArray(MODUsers, userName) || inArray(VIPUsers, userName) {
		//store now epoch on call if empty
		//convert diff(epoch, now) to ms > timeout(ms)
		return true
	} else {
		return false
	}
}

func getItchIOProfile(userName string) string {
	itchioprofile := fmt.Sprintf("https://%s.itch.io", userName)
	if HTTPCheckResponse(itchioprofile) {
		return fmt.Sprintf(" Check their itch.io profile @ %s", itchioprofile)
	} else {
		return ""
	}
}
