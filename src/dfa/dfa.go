package dfa

import (
	"../nfa"
	"../transitionFunction"
	"fmt"
)

type Dfa struct {
	NumStates           int
	States              [][]int
	Symbols             string
	NumAcceptingStates  int
	AcceptingStates     []int
	StartState          int
	TransitionFunctions []transitionFunction.TransitionFunction
}

func (destination *Destinations) isExistWay(way []int) bool {
	for _, value := range destination.Ways {
		if transitionFunction.IntArrayEquals(value, way) {
			return true
		}
	}
	return false
}

type Destinations struct {
	Ways [][]int
}

func (dfa *Dfa) isExistState(state []int) bool {
	for _, value := range dfa.States {
		if transitionFunction.IntArrayEquals(value, state) {
			return true
		}
	}
	return false
}

func isElementInArray(element int, array []int) bool {
	for _, value := range array {
		if value == element {
			return true
		}
	}
	return false
}

func indexOf(element []int, data [][]int) (int, bool) {
	for index, value := range data {
		if transitionFunction.IntArrayEquals(value, element) {
			return index, true
		}
	}
	return -1, false
}

func getKey(transition map[*transitionFunction.TransitionKey][]int, transitionKey transitionFunction.TransitionKey) (bool, *transitionFunction.TransitionKey) {
	for key := range transition {
		if key.TransitionSymbol == transitionKey.TransitionSymbol {
			if transitionFunction.IntArrayEquals(key.StartingState, transitionKey.StartingState) {
				return true, key
			}
		}
	}
	return false, nil
}

func (dfa *Dfa) ConvertFromNfa(nfa nfa.Nfa) {
	dfa.Symbols = nfa.Symbols
	dfa.StartState = nfa.StartState

	nfaTransition := make(map[*transitionFunction.TransitionKey][]int)
	dfaTransition := make(map[*transitionFunction.TransitionKey][]int)

	for _, transition := range nfa.TransitionFunctions {

		nfaTransitionKey := transitionFunction.TransitionKey{[]int{transition.StartingState}, transition.TransitionSymbol}

		if isExist, key := getKey(nfaTransition, nfaTransitionKey); isExist {
			nfaTransition[key] = append(nfaTransition[key], transition.EndingState)
		} else {
			nfaTransition[&nfaTransitionKey] = []int{transition.EndingState}
		}
	}

	dfa.States = append(dfa.States, []int{0})

	dfaStatesSize := len(dfa.States)

	for i := 0; i < dfaStatesSize; i++ {
		for _, symbol := range nfa.Symbols {
			if len(dfa.States[i]) == 1 {
				nfaTransitionKey := transitionFunction.TransitionKey{StartingState: []int{dfa.States[i][0]}, TransitionSymbol: string(symbol)}
				if isExist, key := getKey(nfaTransition, nfaTransitionKey); isExist {
					dfaTransitionKey := transitionFunction.TransitionKey{StartingState: dfa.States[i], TransitionSymbol: string(symbol)}
					dfaTransition[&dfaTransitionKey] = nfaTransition[key]

					if !dfa.isExistState(nfaTransition[key]) {
						dfa.States = append(dfa.States, nfaTransition[key])
					}
				}
			} else {
				var destinations Destinations
				var finalDestination []int

				start := []int{0}
				destinations.Ways = append(destinations.Ways, start)

				for _, nfaState := range dfa.States[i] {
					nfaTransitionKey := transitionFunction.TransitionKey{StartingState: []int{nfaState}, TransitionSymbol: string(symbol)}

					if isExist, key := getKey(nfaTransition, nfaTransitionKey); isExist {
						if _, ok := nfaTransition[key]; ok {
							if !destinations.isExistWay(nfaTransition[key]) {
								destinations.Ways = append(destinations.Ways, nfaTransition[key])
							}
						}
					}
					if len(destinations.Ways) == 0 {
						finalDestination = append(finalDestination, -1)
					} else {
						for _, ways := range destinations.Ways {
							for _, value := range ways {
								if !isElementInArray(value, finalDestination) {
									finalDestination = append(finalDestination, value)
								}
							}
						}
					}
					dfaTransitionKey := transitionFunction.TransitionKey{StartingState: dfa.States[i], TransitionSymbol: string(symbol)}

					if isExist, key := getKey(dfaTransition, dfaTransitionKey); isExist {
						dfaTransition[key] = finalDestination
					} else {
						dfaTransition[&dfaTransitionKey] = finalDestination
					}

					if !dfa.isExistState(finalDestination) {
						dfa.States = append(dfa.States, finalDestination)
					}
				}
			}
		}
		dfaStatesSize = len(dfa.States)
	}

	// Convert NFA states to DFA states
	for key, value := range dfaTransition {
		if startingStateIndex, ok := indexOf(key.StartingState, dfa.States); ok {
			if valueOfIndex, isTrue := indexOf(value, dfa.States); isTrue {
				transFunc := transitionFunction.TransitionFunction{startingStateIndex, key.TransitionSymbol, valueOfIndex}
				dfa.TransitionFunctions = append(dfa.TransitionFunctions, transFunc)
			}
		}
	}
	for _, qState := range dfa.States {
		for _, nfaAcceptingState := range nfa.AcceptingStates {
			if isElementInArray(nfaAcceptingState, qState) {
				index, _ := indexOf(qState, dfa.States)
				dfa.AcceptingStates = append(dfa.AcceptingStates, index)
				dfa.NumAcceptingStates += 1
			}
		}
	}
}

func (dfa Dfa) Print() {
	fmt.Println(len(dfa.States))

	for _, symbol := range dfa.Symbols {
		fmt.Printf("%s ", string(symbol))
	}
	fmt.Println()

	for _, symbol := range dfa.AcceptingStates {
		fmt.Printf("%d ", symbol)
	}
	fmt.Println()

	fmt.Println(dfa.StartState)

	for _, transition := range dfa.TransitionFunctions {
		fmt.Printf("%d %s %d \n", transition.StartingState, transition.TransitionSymbol, transition.EndingState)
	}
}
