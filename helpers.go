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
		// fmt.Printf("found \"%s\" at array index [%d]\n", array[i], i)
		return true
	}
	return false
}

func getAttributedUser(msg string) string {
	return msg[strings.Index(msg, "@")+1:]
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
