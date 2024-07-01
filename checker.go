package goreloaded

import (
	"regexp"
	"strings"
)

func CheckString(word string) string {
	// convert the word into a slice
	slice := strings.Fields(word)
	// checks for combinatons of (cap), (low) ,(up) , (bin), (hex)
	caps := regexp.MustCompile(`\(cap\)`)
	lower := regexp.MustCompile(`\(low\)`)
	upper := regexp.MustCompile(`\(up\)`)
	bin := regexp.MustCompile(`\(bin\)`)
	hex := regexp.MustCompile(`\(hex\)`)

	// patterns for removing different combinations of text with brackets or brackets with spaces
	pattern1 := regexp.MustCompile(`(\s*)\((\w+)\)`)
	pattern2 := regexp.MustCompile(`(\s*)\(\w+\s*,\s*\d+\)`)

	for i, str := range slice {
		if caps.MatchString(str) {
			slice[i-1] = ToCapitalize(slice[i-1])
		} else if lower.MatchString(str) {
			slice[i-1] = ToLower(slice[i-1])
		} else if upper.MatchString(str) {
			slice[i-1] = ToUpper(slice[i-1])
		} else if hex.MatchString(str) {
			slice[i-1] = HexToDecimal(slice[i-1])
		} else if bin.MatchString(str) {
			slice[i-1] = BinToDecimal(slice[i-1])
		}
	}
	// patterns for broken digits
	capdig := regexp.MustCompile(`\(cap,`)
	uppdig := regexp.MustCompile(`\(up,`)
	lowdig := regexp.MustCompile(`\(low,`)

	for i, str := range slice {
		if capdig.MatchString(str) {
			slice = DoubleInput(i, "cap", slice[i+1], slice)
		} else if uppdig.MatchString(str) {
			slice = DoubleInput(i, "up", slice[i+1], slice)
		} else if lowdig.MatchString(str) {
			slice = DoubleInput(i, "low", slice[i+1], slice)
		}
	}

	check1 := strings.Join(slice, " ")
	check2 := pattern1.ReplaceAllString(check1, "")
	check3 := pattern2.ReplaceAllString(check2, "")
	final := WhiteSpace(check3)

	return final
}
