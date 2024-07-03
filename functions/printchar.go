package functions

import (
	"fmt"
	"strings"
)

func PrintChar(userInput string, asciiChar map[int][]string) {
	if userInput == "" {
		return
	}

	// Contrôle des nouvelles lignes
	replaceNewline := strings.ReplaceAll(userInput, "\\n", "\n")
	lines := strings.Split(replaceNewline, "\n")

	// Impression de representations ASCII art
	for _, line := range lines {
		if line == "" {
			fmt.Println()
			continue
		}
		for j := 0; j < len(asciiChar[32]); j++ {
			for _, letter := range line {
				if int(letter) < 32 || int(letter) > 126 {
					fmt.Println("\n NOTICE: no printable character!\n ")
					return
				} else {
					fmt.Print(asciiChar[int(letter)][j])
				}
			}
			fmt.Println()
		}
	}
}
