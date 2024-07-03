package functions

import (
	"bufio"
	"log"
	"os"
)

func OpenFile(asciiChars map[int][]string) {
	file, err := os.Open(".idea/standard.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Création du scanner
	scanned := bufio.NewScanner(file)

	// Configuration du scanner
	scanned.Split(bufio.ScanLines)

	// Stockage des lignes du fichier
	code := 31
	for scanned.Scan() {
		line := scanned.Text()
		if line == "" {
			code++
		} else {
			asciiChars[code] = append(asciiChars[code], line)
		}
	}

	// Si une erreur est rencontrer lors de scanning
	if err := scanned.Err(); err != nil {
		log.Fatal(err)
	}
}
