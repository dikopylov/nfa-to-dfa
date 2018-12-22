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

func KeysIsExist(transition map[*TransitionKey][]int, key TransitionKey) bool {
	for index := range transition {
		if index.TransitionSymbol == key.TransitionSymbol {
			if IntArrayEquals(key.StartingState, index.StartingState) {
				return true
			}
		}
	}
	return false
}
