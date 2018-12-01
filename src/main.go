package main

import (
	nfaSrc "./nfa"
	transitionFunctionSrc "./transitionFunction"
	"bufio"
	"fmt"
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

	lines, err := readLines("H:/Documents/GoLandProjects/NfaToDfa/src/test.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	nfa := nfaSrc.Nfa{}
	transitionFunction := transitionFunctionSrc.TransitionFunction{}
	nfa.ConstructNfaFromFile(lines, transitionFunction)

	fmt.Print(nfa)
}
