package cards

type Card struct {
	Front          string
	Back           string
	IsFlipped      bool
	TimesCorrect   int
	TimesIncorrect int
	Hidden         bool
}
