package model

import tea "github.com/charmbracelet/bubbletea"

func actOnCard(m Model) (tea.Model, tea.Cmd) {
	if m.SelectedDeck == nil {
		return selectDeck(m)
	} else {
		return flipCard(m)
	}
}

func changeCard(m Model, direction string) (tea.Model, tea.Cmd) {
	var nextPos int
	var maxPos int

	if m.SelectedDeck == nil {
		maxPos = len(m.Decks) - 1
	} else {
		maxPos = len(m.Cards) - 1
	}

	nextPos = m.CurrentCard + 1
	if direction == "left" {
		nextPos = m.CurrentCard - 1
	}

	if nextPos > maxPos {
		nextPos = 0
	}

	if nextPos < 0 {
		nextPos = maxPos
	}

	m.CurrentCard = nextPos
	return m, nil
}
