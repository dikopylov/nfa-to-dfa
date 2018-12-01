package dfa

import (
	"../nfa"
	"../transitionFunction"
	"strconv"
)

type Dfa struct {
	NumStates           int
	States              []*[]int
	Symbols             string
	NumAcceptingStates  int
	AcceptingStates     []int
	StartState          int
	TransitionFunctions []transitionFunction.TransitionFunction
}

func (dfa *Dfa) isStateInStatesArray(state *[]int) bool {
	for _, value := range dfa.States {
		if value == state {
			return true
		}
	}
	return false
}

/**
# Convert NFA transitions to DFA transitions
        for dfa_state in self.q:
            for symbol in nfa.symbols:
                if len(dfa_state) == 1 and (dfa_state[0], symbol) in nfa_transition_dict:
                    dfa_transition_dict[(dfa_state, symbol)] = nfa_transition_dict[(dfa_state[0], symbol)]

                    if tuple(dfa_transition_dict[(dfa_state, symbol)]) not in self.q:
                        self.q.append(tuple(dfa_transition_dict[(dfa_state, symbol)]))
                else:
                    destinations = []
                    final_destination = []

*/

func (dfa *Dfa) convertFromNfa(nfa nfa.Nfa) {
	dfa.Symbols = nfa.Symbols
	dfa.StartState = nfa.StartState

	var nfaTransitionKey transitionFunction.TransitionKey
	var dfaTransitionKey transitionFunction.TransitionKey

	nfaTransition := make(map[*transitionFunction.TransitionKey][]int)
	dfaTransition := make(map[*transitionFunction.TransitionKey][]int)

	for _, transition := range nfa.TransitionFunctions {

		nfaTransitionKey.StartingState = []int{transition.StartingState}
		nfaTransitionKey.TransitionSymbol = transition.TransitionSymbol

		nfaTransition[&nfaTransitionKey] = append(nfaTransition[&nfaTransitionKey], transition.EndingState)
	}

	start := new([]int)
	*start = []int{1}
	dfa.States = append(dfa.States, start)

	for _, dfaState := range dfa.States {
		for _, symbol := range nfa.Symbols {
			if len(*dfaState) == 1 {
				firstState := *dfaState
				nfaTransitionKey := transitionFunction.TransitionKey{StartingState: []int{firstState[0]}, TransitionSymbol: strconv.Itoa(int(symbol))}
				if _, ok := nfaTransition[&nfaTransitionKey]; ok {
					dfaTransitionKey := transitionFunction.TransitionKey{StartingState: firstState, TransitionSymbol: strconv.Itoa(int(symbol))}
					dfaTransition[&dfaTransitionKey] = nfaTransition[&nfaTransitionKey]

				}
			}
		}
	}
}
