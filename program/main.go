package main

import (
	"fmt"
	"os"

	"goreloaded"
)

func main() {
	// captures the file names in the terminal
	filename := os.Args[1:]
	if len(filename) != 2 {
		fmt.Println("Error with the number of arguments present")
		return
	}
	// captures the first file with the text , reads the data to a variable named file
	file, err := os.ReadFile(filename[0])
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}

	// create and empyty string and ranges the file content which are in bytes into the string by converting them to strings then adding them
	var sentence string
	for _, v := range file {
		sentence = sentence + string(v)
	}
	// correct punctuation
	newsentence := goreloaded.Punct(sentence)
	// sort for vowels
	newlsice := goreloaded.SortVowels(newsentence)
	// sort for the patterns in the string like (cap) or (cap, 4)
	squotes := goreloaded.CheckString(newlsice)
	// this section sorts out the single quotes
	sorted := goreloaded.SortQuotes(squotes)
	// adding a newline
	sorted = sorted + "\n"
	convert := []byte(sorted)
	// we write the result to the next file captured
	os.WriteFile(filename[1], convert, 0o644)
}
