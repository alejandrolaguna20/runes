package cards

type Card struct {
	Front          string
	Back           string
	IsFlipped      bool
	timesCorrect   int
	timesIncorrect int
}
