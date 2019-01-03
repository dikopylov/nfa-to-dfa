package nfa

import (
	"../enums"
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

func (nfa *Nfa) isLastAcceptingState(state int) bool {
	result := false
	for _, acceptingState := range nfa.AcceptingStates {
		if acceptingState == state {
			result = true
		}
	}

	if result {
		for _, transitionFunc := range nfa.TransitionFunctions {
			if transitionFunc.StartingState == state {
				result = false
			}
		}
	}
	return result
}

func getKey(haystack []int, value int) int {
	for key, val := range haystack {
		if val == value {
			return key
		}
	}

	return -1
}

func getKeyFromTransitions(haystack []transitionFunction.TransitionFunction, value transitionFunction.TransitionFunction) int {
	for key, val := range haystack {
		if val.StartingState == value.StartingState &&
			val.TransitionSymbol == value.TransitionSymbol &&
			val.EndingState == value.EndingState {
			return key
		}
	}

	return -1
}

func (nfa *Nfa) InitStates() {
	for i := 0; i < nfa.NumStates; i++ {
		nfa.States = append(nfa.States, i)
	}
}

func (nfa *Nfa) Print() {
	fmt.Println(nfa.NumStates)
}

func (nfa *Nfa) addTransitionFunction(startState int, endState int, transitionSymbol string) {
	transitionFunc := transitionFunction.TransitionFunction{}

	transitionFunc.StartingState = startState
	transitionFunc.TransitionSymbol = transitionSymbol
	transitionFunc.EndingState = endState

	nfa.TransitionFunctions = append(nfa.TransitionFunctions, transitionFunc)
}

func (nfa *Nfa) ConstructNfaFromFile(nfaTxt []string) {
	nfa.NumStates, _ = strconv.Atoi(nfaTxt[enums.NumStatesLine])
	epsTransitKey := 0
	epsTransitions := make(map[int]transitionFunction.TransitionFunction)
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

		if transitionSymbol == enums.EpsSymbol && startState != endState {
			epsTransitions[epsTransitKey] = transitionFunction.TransitionFunction{
				StartingState:    startState,
				TransitionSymbol: enums.EpsSymbol,
				EndingState:      endState,
			}
			epsTransitKey++
		} else {
			nfa.addTransitionFunction(startState, endState, transitionSymbol)
		}
	}

	if len(epsTransitions) > 0 {
		nfaTransitions := nfa.TransitionFunctions

		for len(epsTransitions) > 0 {
			for key, epsTransition := range epsTransitions {
				for _, nfaTransition := range nfaTransitions {
					if nfaTransition.EndingState == epsTransition.StartingState {
						nfa.addTransitionFunction(nfaTransition.StartingState, epsTransition.EndingState, nfaTransition.TransitionSymbol)
						delete(epsTransitions, key)
					} else if nfaTransition.StartingState == nfa.StartState &&
						epsTransition.StartingState == nfaTransition.StartingState &&
						nfa.isLastAcceptingState(epsTransition.EndingState) {
						acceptingStateKey := getKey(nfa.AcceptingStates, epsTransition.EndingState)
						nfa.AcceptingStates = append(nfa.AcceptingStates[:acceptingStateKey], nfa.AcceptingStates[acceptingStateKey+1:]...)
						nfa.NumStates--
						nfa.AcceptingStates = append(nfa.AcceptingStates, nfaTransition.StartingState)
						delete(epsTransitions, key)
					}
				}
			}
		}
	}

	nfaTransitions := make([]transitionFunction.TransitionFunction, len(nfa.TransitionFunctions))
	for _, nfaTransition := range nfa.TransitionFunctions {
		nfaTransitions = append(nfaTransitions, nfaTransition)
	}

	for _, nfaTransition := range nfaTransitions {
		if len(nfaTransition.TransitionSymbol) > 1 {
			transitionStateKey := getKeyFromTransitions(nfa.TransitionFunctions, nfaTransition)
			nfa.TransitionFunctions = append(nfa.TransitionFunctions[:transitionStateKey], nfa.TransitionFunctions[transitionStateKey+1:]...)

			transitionSymbols := strings.Split(nfaTransition.TransitionSymbol, "")
			nfa.addTransitionFunction(nfaTransition.StartingState, nfa.NumStates, transitionSymbols[0])
			nfa.NumStates++

			lastIndex := len(transitionSymbols) - 1
			for _, symbol := range transitionSymbols[1:lastIndex] {
				nfa.addTransitionFunction(nfa.NumStates-1, nfa.NumStates, symbol)
				nfa.NumStates++
			}

			nfa.addTransitionFunction(nfa.NumStates-1, nfaTransition.EndingState, transitionSymbols[lastIndex])

		}
	}

	nfa.InitStates()
}
