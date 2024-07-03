package main

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
	"testing"
)

/*
Chaque clé de la map testCases contient le nom du fichier contenant
le résultat attendu pour chaque scénario de test, la valeur de chaque clé
est une tranche de chaînes, le premier élément contient la chaîne à donner
en argument au moment de l'exécution, la seconde contiendra la chaîne
équivalente à celle attendue sortir
*/
var testCases = map[int][]string{
	1:  {"hello", ""},
	2:  {"HELLO", ""},
	3:  {"HeLlo HuMaN", ""},
	4:  {"1Hello 2There", ""},
	5:  {"Hello\\nThere", ""},
	6:  {"{Hello & There #}", ""},
	7:  {"hello There 1 to 2!", ""},
	8:  {"MaD3IrA&LiSboN", ""},
	9:  {"1a\"#FdwHywR&/()=", ""},
	10: {"{|}~", ""},
	11: {"[\\]^_ 'a", ""},
	12: {"RGB", ""},
	13: {":;<=>?@", ""},
	14: {"\\!\" #$%&'()*+,-./", ""},
	15: {"ABCDEFGHIJKLMNOPQRSTUVWXYZ", ""},
	16: {"abcdefghijklmnopqrstuvwxyz", ""},
}

/*
Ce fichier de test teste le projet ascii-art par rapport aux 16 premiers
cas de test sur page de vérification
*/
func TestAsciiArt(t *testing.T) {
	getTestCases()

	/*	Parcourez chaque cas de test et démarrez une goroutine pour chacun,
		ceci est fait, au lieu d'attendre la fin du test précédent,
		ils peuvent tout être vérifié simultanément	*/
	var wg sync.WaitGroup
	for i := 1; i <= len(testCases); i++ {
		wg.Add(1)
		go func(current []string, w *sync.WaitGroup, ti *testing.T) {
			defer w.Done()
			result := getResult(current)
			/*	Le projet échoue si le résultat attendu des cas de test
				ne correspond pas la sortie réelle	*/
			if string(result) != current[1] {
				ti.Errorf("\nTest fails when given the test case:\n\t\"%s\","+
					"\nexpected:\n%s\ngot:\n%s\n\n",
					current[0], current[1], string(result))
			}
		}(testCases[i], &wg, t)
	}
	wg.Wait()
}

/*
Cette fonction imite l'exécution de "go run . string", qu'elle redirige ensuite
dans une deuxième fonction "cat -e" pour imiter puis renvoie le résultat
*/
func getResult(testCase []string) string {
	first := exec.Command("go", "run", ".", testCase[0])
	second := exec.Command("cat", "-e")
	reader, writer := io.Pipe()
	first.Stdout = writer
	second.Stdin = reader
	var buffer bytes.Buffer
	second.Stdout = &buffer
	first.Start()
	second.Start()
	first.Wait()
	writer.Close()
	second.Wait()
	return buffer.String()
}

/*
Cette fonction lit chacun des cas de test attendus en sortie du "testcases.txt"
et les ajoute à la tranche de cas de test correspondante dans la map testCases
*/
func getTestCases() {
	file, err := os.Open(".idea/testcases.txt")
	if err != nil {
		panic(err)
	}

	stats, _ := file.Stat()
	contents := make([]byte, stats.Size())
	file.Read(contents)
	lines := strings.Split(string(contents), "\n")

	start := 0
	number := 0
	for i, line := range lines {
		if i == len(lines)-1 {
			testCases[number][1] = strings.Join(lines[start:], "\n") + "\n"
			break
		}
		if line[0] == '#' && line[len([]rune(line))-1] == '#' {
			if i > 0 {
				testCases[number][1] = strings.Join(lines[start:i], "\n") + "\n"
			}
			start = i + 1
			number++
		}
	}
}
