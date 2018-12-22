package dfa

import (
	"../nfa"
	"../transitionFunction"
	"strconv"
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

type Destinations struct {
	Ways []*[]int
}

//func (destination *Destinations) isExistWay(way []int) bool {
//	for _, value := range destination.Ways {
//		if intArrayEquals(*value, way) {
//			return true
//		}
//	}
//	return false
//}

//func (dfa *Dfa) isExistState(state []int) bool {
//	for _, value := range dfa.States {
//		if intArrayEquals(*value, state) {
//			return true
//		}
//	}
//	return false
//}

//func (dfa *Dfa) getStateIndex(key transitionFunction.TransitionKey) bool {
//	for i, value := range dfa.States {
//		value = value
//		if intArrayEquals(*value, []int{key.StartingState[0], key.TransitionSymbol}) {
//			return true
//		}
//	}
//	return false
//}


func isElementInArray(element int, array []int) bool {
	for _, value := range array {
		if value == element {
			return true
		}
	}
	return false
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

	//var dfaTransitionKey transitionFunction.TransitionKey

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

	for _, dfaState := range dfa.States {
		for _, symbol := range nfa.Symbols {
			if len(dfaState) == 1 {
				nfaTransitionKey := transitionFunction.TransitionKey{StartingState: []int{dfaState[0]}, TransitionSymbol: string(symbol)}
				if isExist, key := getKey(nfaTransition, nfaTransitionKey); isExist {
					dfaTransitionKey := transitionFunction.TransitionKey{StartingState: dfaState, TransitionSymbol: strconv.Itoa(int(symbol))}
					dfaTransition[&dfaTransitionKey] = nfaTransition[key]

					if _, ok := nfaTransition[nfaTransitionKey]; ok {
					dfaTransitionKey := transitionFunction.TransitionKey{StartingState: dfaState, TransitionSymbol: strconv.Itoa(int(symbol))}
					dfaTransition[&dfaTransitionKey] = nfaTransition[&nfaTransitionKey]

					if !dfa.isExistState(nfaTransition[&nfaTransitionKey]) {
						pState := new([]int)
						*pState = nfaTransition[&nfaTransitionKey]
						dfa.States = append(dfa.States, pState)
					}
				}
				//			} else {
				//				var destinations Destinations
				//				var finalDestination []int
				//
				//				start := new([]int)
				//				*start = []int{0} // {1}
				//				destinations.Ways = append(destinations.Ways, start)
				//
				//				for _, nfaState := range *dfaState {
				//					nfaTransitionKey = transitionFunction.TransitionKey{StartingState: []int{nfaState}, TransitionSymbol: strconv.Itoa(int(symbol))}
				//					if _, ok := nfaTransition[&nfaTransitionKey]; ok {
				//						if destinations.isExistWay(nfaTransition[&nfaTransitionKey]); !ok {
				//							pNfaTransition := new([]int)
				//							*pNfaTransition = nfaTransition[&nfaTransitionKey]
				//							destinations.Ways = append(destinations.Ways, pNfaTransition)
				//						}
				//					}
				//				}
				//				if len(destinations.Ways) == 0 {
				//					finalDestination = append(finalDestination, -1)
				//				} else {
				//					for _, ways := range destinations.Ways {
				//						for _, value := range *ways {
				//							if !isElementInArray(value, finalDestination) {
				//								finalDestination = append(finalDestination, value)
				//							}
				//						}
				//					}
				//				}
				//				dfaTransition[&dfaTransitionKey] = finalDestination
				//
				//				if !dfa.isExistState(finalDestination) {
				//					dfa.States = append(dfa.States, &finalDestination)
				//				}
				//			}
				//		}
				//	}
			}
			/**
		# Convert NFA states to DFA states
	        for key in dfa_transition_dict:
	            self.transition_functions.append(
	                (self.q.index(tuple(key[0])), key[1], self.q.index(tuple(dfa_transition_dict[key]))))

	        for q_state in self.q:
	            for nfa_accepting_state in nfa.accepting_states:
	                if nfa_accepting_state in q_state:
	                    self.accepting_states.append(self.q.index(q_state))
	                    self.num_accepting_states += 1
	*/
			// Convert NFA states to DFA states
			//for key, value := range dfaTransition {
			//	transitionFunction := transitionFunction.TransitionFunction{key.StartingState[0], key.TransitionSymbol, value.StartingState[0]}
			//	dfa.TransitionFunctions = append(dfa.TransitionFunctions, )
			//}
		}
	}
}
