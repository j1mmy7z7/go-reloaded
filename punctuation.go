package goreloaded

import "regexp"

func Punct(word string) string {
	// pattern checks for a one or more whitespace followed by one or more punctuation
	pattern := regexp.MustCompile(`(\s+)([,.:;!?]+)`)
	// pattern2 checks for punctuations followed by words
	pattern2 := regexp.MustCompile(`([,.;:!?]+)(\w+)`)
	// replace the pattern with the 2nd captured group which is the punctuations hence eliminating spaces
	initial := pattern.ReplaceAllString(word, "$2")
	// add space where there is a punctuation and a word following it
	sorted := pattern2.ReplaceAllString(initial, "$1 $2")
	// eliminate extra whitespaces
	final := WhiteSpace(sorted)

	return final
}
