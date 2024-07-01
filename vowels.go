package goreloaded

import (
	"regexp"
	"strings"
)

func SortVowels(word string) string {
	slice := strings.Fields(word)
	// regex pattern for checking if a word starts with the words inside the square brackets
	pattern := regexp.MustCompile(`(?i)^[aeiouh].*`)
	for i, str := range slice {
		if str == "a" || str == "A" {
			if pattern.MatchString(slice[i+1]) {
				slice[i] += "n"
			}
		}
	}
	final := strings.Join(slice, " ")
	return final
}
