package nfa

import (
	enums "../enums"
	"../transitionFunction"
	"fmt"
	"strconv"
	"strings"
)

type Nfa struct {
	NumStates           int
	States              []int
	Symbols             string
	NumAcceptingStates  int
	AcceptingStates     []int
	StartState          int
	TransitionFunctions []transitionFunction.TransitionFunction
}

func (nfa *Nfa) InitStates() {
	for i := 0; i < nfa.NumStates; i++ {
		nfa.States = append(nfa.States, i)
	}
}

func (nfa *Nfa) Print() {
	fmt.Println(nfa.NumStates)
}

func (nfa *Nfa) AddTransitionFunction(startState int, endState int, transitionSymbol string) {
	transitionFunc := transitionFunction.TransitionFunction{}

	transitionFunc.StartingState = startState
	transitionFunc.TransitionSymbol = transitionSymbol
	transitionFunc.EndingState = endState

	nfa.TransitionFunctions = append(nfa.TransitionFunctions, transitionFunc)
}

func (nfa *Nfa) ConstructNfaFromFile(nfaTxt []string) {
	nfa.NumStates, _ = strconv.Atoi(nfaTxt[enums.NumStatesLine])

	nfa.Symbols = strings.TrimSpace(nfaTxt[enums.ValidCharactersLine])

	acceptingStatesLine := strings.Split(nfaTxt[enums.AcceptingStatesLine], " ")

	nfa.NumAcceptingStates = len(acceptingStatesLine)
	for _, value := range acceptingStatesLine {
		state, _ := strconv.Atoi(value)

		nfa.AcceptingStates = append(nfa.AcceptingStates, state)
	}

	nfa.StartState, _ = strconv.Atoi(nfaTxt[enums.StartStateLine])

	for line := enums.FunctionsStartLine; line < len(nfaTxt); line++ {
		transitionFuncLine := strings.Split(nfaTxt[line], " ")
		startState, _ := strconv.Atoi(transitionFuncLine[0])
		transitionSymbol := transitionFuncLine[1]
		endState, _ := strconv.Atoi(transitionFuncLine[2])

		if len(transitionSymbol) > 1 {
			transitionSymbols := strings.Split(transitionSymbol, "")
			nfa.AddTransitionFunction(startState, nfa.NumStates, transitionSymbols[0])
			nfa.NumStates++

			lastIndex := len(transitionSymbols) - 1
			for _, symbol := range transitionSymbols[1:lastIndex] {
				nfa.AddTransitionFunction(nfa.NumStates-1, nfa.NumStates, symbol)
				nfa.NumStates++
			}

			nfa.AddTransitionFunction(nfa.NumStates-1, endState, transitionSymbols[lastIndex])
		} else if transitionFuncLine[1] == "E" {

		} else {
			nfa.AddTransitionFunction(startState, endState, transitionSymbol)
		}
	}

	nfa.InitStates()
}
