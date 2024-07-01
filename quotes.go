package goreloaded

import (
	"regexp"
	"strings"
)

func SortQuotes(word string) string {
	// split the string to a slice
	slice := strings.Fields(word)
	for i, v := range slice {
		// checks if the values found is a slice and if it the start of the array
		if v == "'" && i != 0 {
			// if not the first value the single quote is merged with the previous word
			slice[i-1] = slice[i-1] + slice[i]
			slice[i] = ""
		}
	}

	pattern := regexp.MustCompile(`'\s*([^']+)'`)
	initial := strings.Join(slice, " ")
	initial2 := pattern.ReplaceAllString(initial, " '$1'")
	check := strings.Fields(initial2)
	final := strings.Join(check, " ")
	return final
}
