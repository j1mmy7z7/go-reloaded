package goreloaded

import (
	"regexp"
	"strconv"
	"strings"
)

// this file handles all the conversion done to the strings

func ToCapitalize(letter string) string {
	new := strings.Title(letter)
	return new
}

func ToUpper(letter string) string {
	new := strings.ToUpper(letter)
	return new
}

func ToLower(letter string) string {
	new := strings.ToLower(letter)
	return new
}

func HexToDecimal(letter string) string {
	decimal, _ := strconv.ParseInt(letter, 16, 64)
	new := strconv.Itoa(int(decimal))
	return new
}

func BinToDecimal(letter string) string {
	binary, _ := strconv.ParseInt(letter, 2, 64)
	new := strconv.Itoa(int(binary))
	return new
}

func WhiteSpace(word string) string {
	white := regexp.MustCompile(`\s+`)
	new := white.ReplaceAllString(word, " ")
	return new
}
