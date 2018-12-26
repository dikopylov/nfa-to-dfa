package dfa

import (
	"../nfa"
	"../transitionFunction"
	"fmt"
	"sort"
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

func (dfa *Dfa) sortTransitionFunctions() {
	sort.SliceStable(dfa.TransitionFunctions,
		func(i, j int) bool {
			current := dfa.TransitionFunctions[j]
			following := dfa.TransitionFunctions[i]

			result := following.StartingState < current.StartingState
			if following.StartingState == current.StartingState {
				result = following.TransitionSymbol < current.TransitionSymbol
			}

			return result
		})
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

func initCondition(dfaStates []int, nfaTransition map[*transitionFunction.TransitionKey][]int, transitionSymbol string) (bool, *transitionFunction.TransitionKey) {
	if len(dfaStates) == 1 {
		nfaTransitionKey := transitionFunction.TransitionKey{StartingState: []int{dfaStates[0]}, TransitionSymbol: string(transitionSymbol)}
		if isExist, key := getKey(nfaTransition, nfaTransitionKey); isExist {
			return isExist, key
		}
	}

	return false, nil
}

/**
Преобразуем исходные функции перехода НКА в карту nfaTransition,
где ключ - Структура TransitionKey, а значение - срез состояний,
в которые осуществляется переход из StartingState по символу TransitionSymbol
*/
func makeNFATransitionMap(transitionFunctions []transitionFunction.TransitionFunction) map[*transitionFunction.TransitionKey][]int {
	nfaTransition := make(map[*transitionFunction.TransitionKey][]int)

	for _, transition := range transitionFunctions {

		nfaTransitionKey := transitionFunction.TransitionKey{[]int{transition.StartingState}, transition.TransitionSymbol}

		if isExist, key := getKey(nfaTransition, nfaTransitionKey); isExist {
			nfaTransition[key] = append(nfaTransition[key], transition.EndingState)
		} else {
			nfaTransition[&nfaTransitionKey] = []int{transition.EndingState}
		}
	}

	return nfaTransition
}

func makeDFATransitionMap(dfa *Dfa, nfaTransition map[*transitionFunction.TransitionKey][]int) map[*transitionFunction.TransitionKey][]int {
	dfaTransition := make(map[*transitionFunction.TransitionKey][]int)

	dfa.States = append(dfa.States, []int{0})

	dfaStatesSize := len(dfa.States)

	for i := 0; i < dfaStatesSize; i++ {
		for _, symbol := range dfa.Symbols {
			if ok, key := initCondition(dfa.States[i], nfaTransition, string(symbol)); ok {
				dfaTransitionKey := transitionFunction.TransitionKey{StartingState: dfa.States[i], TransitionSymbol: string(symbol)}
				dfaTransition[&dfaTransitionKey] = nfaTransition[key]

				if !dfa.isExistState(nfaTransition[key]) {
					dfa.States = append(dfa.States, nfaTransition[key])
				}
			} else {
				var destinations Destinations
				var finalDestination []int

				for _, nfaState := range dfa.States[i] {
					nfaTransitionKey := transitionFunction.TransitionKey{StartingState: []int{nfaState}, TransitionSymbol: string(symbol)}

					if isExist, key := getKey(nfaTransition, nfaTransitionKey); isExist {
						if !destinations.isExistWay(nfaTransition[key]) {
							destinations.Ways = append(destinations.Ways, nfaTransition[key])
						}
					}
				}
				if len(destinations.Ways) == 0 {
					finalDestination = nil
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
		dfaStatesSize = len(dfa.States)
	}

	return dfaTransition
}

func (dfa *Dfa) ConvertFromNfa(nfa nfa.Nfa) {
	dfa.Symbols = nfa.Symbols
	dfa.StartState = nfa.StartState

	nfaTransition := makeNFATransitionMap(nfa.TransitionFunctions)
	dfaTransition := makeDFATransitionMap(dfa, nfaTransition)

	// Формируем функции перехода для ДКА
	for key, value := range dfaTransition {
		if startingStateIndex, ok := indexOf(key.StartingState, dfa.States); ok {
			if valueOfIndex, isTrue := indexOf(value, dfa.States); isTrue {
				transFunc := transitionFunction.TransitionFunction{startingStateIndex, key.TransitionSymbol, valueOfIndex}
				dfa.TransitionFunctions = append(dfa.TransitionFunctions, transFunc)
			}
		}
	}

	dfa.sortTransitionFunctions()

	for _, qState := range dfa.States {
		for _, nfaAcceptingState := range nfa.AcceptingStates {
			if isElementInArray(nfaAcceptingState, qState) {
				index, _ := indexOf(qState, dfa.States)
				if !isElementInArray(index, dfa.AcceptingStates) {
					dfa.AcceptingStates = append(dfa.AcceptingStates, index)
					dfa.NumAcceptingStates += 1
				}
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
