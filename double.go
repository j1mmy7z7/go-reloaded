package goreloaded

import (
	"strconv"
)

// handles the conversions which have digits with them

func DoubleInput(index int, command string, times string, slice []string) []string {
	var num string
	// extract the number of times the loop will go behind
	for _, v := range times {
		if v >= '0' && v <= '9' {
			num += string(v)
		}
	}
	// change the value captured to an integer
	new, _ := strconv.Atoi(num)
	// keep the number of available words to change since we don't any index out of range errors
	checker := index - 1
	// checks for commands and call for the appropriate functions
	if command == "cap" {
		// this for loops go behinds the captured array changes the word as long it does not go over the array|slice, if it does the loop stops
		for i := 0; i < new && i-checker <= 0; i++ {
			word := ToCapitalize(slice[index-1])
			slice[index-1] = word
			index = index - 1
		}
	} else if command == "up" {
		for i := 0; i < new && i-checker <= 0; i++ {
			word := ToUpper(slice[index-1])
			slice[index-1] = word
			index = index - 1
		}
	} else if command == "low" {
		for i := 0; i < new && i-checker <= 0; i++ {
			word := ToLower(slice[index-1])
			slice[index-1] = word
			index = index - 1
		}
	}
	return slice
}
