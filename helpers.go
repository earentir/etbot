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
	if strings.Index(msg, "@") > -1 {
		return msg[strings.Index(msg, "@"):]
	} else {
		return ""
	}
}

func getCleanMessage(cmd, msg string) string {
	if strings.Index(msg, "@") > -1 {
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

// func IsItOnTimeout(command string) bool {

// }
