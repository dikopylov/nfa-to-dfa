package nfa

import (
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

func (nfa *Nfa) ConstructNfaFromFile(nfaTxt []string, transitionFunction transitionFunction.TransitionFunction) {
	nfa.NumStates, _ = strconv.Atoi(nfaTxt[0])

	nfa.InitStates()
	nfa.Symbols = strings.TrimSpace(nfaTxt[1])

	acceptingStatesLine := strings.Split(nfaTxt[2], " ")

	nfa.NumAcceptingStates = len(acceptingStatesLine)
	for _, value := range acceptingStatesLine {
		state, _ := strconv.Atoi(value)

		nfa.AcceptingStates = append(nfa.AcceptingStates, state)
	}

	nfa.StartState, _ = strconv.Atoi(nfaTxt[3])

	for line := 4; line < len(nfaTxt); line++ {
		transitionFuncLine := strings.Split(nfaTxt[line], " ")

		transitionFunction.StartingState, _ = strconv.Atoi(transitionFuncLine[0])
		transitionFunction.TransitionSymbol = transitionFuncLine[1]
		transitionFunction.EndingState, _ = strconv.Atoi(transitionFuncLine[2])

		nfa.TransitionFunctions = append(nfa.TransitionFunctions, transitionFunction)
	}
}
