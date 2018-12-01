/**
class NFA:
    def __init__(self):
        self.num_states = 0
        self.states = []
        self.symbols = []
        self.num_accepting_states = 0
        self.accepting_states = []
        self.start_state = 0
        self.transition_functions = []

    def init_states(self):
        self.states = list(range(self.num_states))

    def print_nfa(self):
        print(self.num_states)
        print(self.states)
        print(self.symbols)
        print(self.num_accepting_states)
        print(self.accepting_states)
        print(self.start_state)
        print(self.transition_functions)
*/
package NfaToDfa

import (
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
	TransitionFunctions []int
}

func (nfa Nfa) initStates() {
	for i := 0; i < nfa.NumStates; i++ {
		nfa.States = append(nfa.States, i)
	}
}

func (nfa Nfa) print() {
	fmt.Println(nfa)
}

/**
  def construct_nfa_from_file(self, lines):
      self.num_states = int(lines[0])
      self.init_states()
      self.symbols = list(lines[1].strip())

      accepting_states_line = lines[2].split(" ")
      for index in range(len(accepting_states_line)):
          if index == 0:
              self.num_accepting_states = int(accepting_states_line[index])
          else:
              self.accepting_states.append(int(accepting_states_line[index]))

      self.startState = int(lines[3])

      for index in range(4, len(lines)):
          transition_func_line = lines[index].split(" ")

          starting_state = int(transition_func_line[0])
          transition_symbol = transition_func_line[1]
          ending_state = int(transition_func_line[2])

          transition_function = (starting_state, transition_symbol, ending_state);
          self.transition_functions.append(transition_function)
*/

func (nfa Nfa) constructNfaFromFile(nfaTxt []string) {
	nfa.NumStates, _ = strconv.Atoi(nfaTxt[0])

	nfa.initStates()
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

		// Тут int/string/int кладется в один список
		//startingState, _ := strconv.Atoi(transitionFuncLine[0])
		//transitionSymbol := transitionFuncLine[1]
		//endingState, _ := strconv.Atoi(transitionFuncLine[2])

		startingState := transitionFuncLine[0]
		transitionSymbol := transitionFuncLine[1]
		endingState := transitionFuncLine[2]

		transitionFunction := [...]string{startingState, transitionSymbol, endingState}
		nfa.TransitionFunctions = append(nfa.TransitionFunctions, transitionFunction)
	}
}