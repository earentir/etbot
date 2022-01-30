package main

import (
	"sort"
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
