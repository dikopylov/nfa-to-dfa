package main

import (
	dfaSrc "./dfa"
	nfaSrc "./nfa"
	"bufio"
	"log"
	"os"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	lines, err := readLines("C:/Users/arman/Desktop/Portfolio/nfa-to-dfa/src/test4.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	nfa := nfaSrc.Nfa{}
	nfa.ConstructNfaFromFile(lines)

	dfa := dfaSrc.Dfa{}
	dfa.ConvertFromNfa(nfa)

	dfa.Print()
}
