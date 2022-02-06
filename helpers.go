package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// func inArray(array []string, lookup string) bool {
// 	sort.Strings(array)
// 	i := sort.SearchStrings(array, lookup)
// 	if i < len(array) && array[i] == lookup {
// 		return true
// 	}
// 	return false
// }

func getAttributedUser(msg string, at bool) string {
	if strings.Contains(msg, "@") {
		if at {
			return msg[strings.Index(msg, "@"):]
		} else {
			return msg[strings.Index(msg, "@")+1:]
		}

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
	if true {
		//store now epoch on call if empty
		//convert diff(epoch, now) to ms > timeout(ms)
		return true
	} else {
		return false
	}
}

func getItchIOProfile(userName string) string {
	if userName != "" {
		itchioprofile := fmt.Sprintf("https://%s.itch.io", userName)
		if HTTPCheckResponse(itchioprofile) {
			return fmt.Sprintf(" Check their itch.io profile @ %s", itchioprofile)
		} else {
			return ""
		}
	} else {
		return ""
	}
}

func CPrint(color, msg string) {
	var colorCode string
	var resetcolor string = "\033[0m"

	switch strings.ToLower(color) {
	case "r":
		fallthrough
	case "red":
		colorCode = "\033[31m"

	case "g":
		fallthrough
	case "green":
		colorCode = "\033[32m"

	case "y":
		fallthrough
	case "yellow":
		colorCode = "\033[33m"

	case "b":
		fallthrough
	case "blue":
		colorCode = "\033[34m"

	case "p":
		fallthrough
	case "purple":
		colorCode = "\033[35m"

	case "c":
		fallthrough
	case "cyan":
		colorCode = "\033[36m"

	default:
		colorCode = resetcolor
	}

	fmt.Println(string(colorCode), msg)
	fmt.Print(string(resetcolor))
}
