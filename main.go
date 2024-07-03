package main

import (
	"fmt"
	"os"
	"regexp"

	utils "asciiart/functions"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("\n USAGE: go run main.go 'your string'\n ")
		return
	}

	userInput := os.Args[1]

	// Contrôle des nouvelles lignes
	if userInput == "\\n" {
		fmt.Println()
		return
	}

	pattern := regexp.MustCompile(`\A((\\n)+)\\n`)
	userInput = pattern.ReplaceAllString(userInput, "$1")

	asciiChars := make(map[int][]string)

	// Collection de ASCII Art de standard.txt
	utils.OpenFile(asciiChars)
	// Impression de ASCII Art correspondant
	utils.PrintChar(userInput, asciiChars)
}
