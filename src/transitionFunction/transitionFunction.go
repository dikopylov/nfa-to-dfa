package transitionFunction

type TransitionFunction struct {
	StartingState    int
	TransitionSymbol string
	EndingState      int
}

type TransitionKey struct {
	StartingState    []int
	TransitionSymbol string
}

func IntArrayEquals(arrayOne []int, arrayTwo []int) bool {
	if len(arrayOne) != len(arrayTwo) {
		return false
	}
	for i, v := range arrayOne {
		if v != arrayTwo[i] {
			return false
		}
	}
	return true
}
