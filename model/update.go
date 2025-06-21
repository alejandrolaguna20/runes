package model

import (
	"github.com/alejandrolaguna20/runes/decks"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

func resize(m Model, msg tea.WindowSizeMsg) (tea.Model, tea.Cmd) {
	if !m.ready {
		m.viewport = viewport.New(msg.Width, msg.Height)
		m.ready = true
	} else {
		m.viewport.Width = msg.Width
		m.viewport.Height = msg.Height
	}
	return m, nil
}

func flipCard(m Model) (tea.Model, tea.Cmd) {
	currentCard := m.GetCurrentCard()
	if currentCard.IsFlipped {
		m.CardColor = CardDefaultColor
	} else {
		m.CardColor = CardBlueColor
	}
	currentCard.IsFlipped = !currentCard.IsFlipped
	return m, nil
}

func selectDeck(m Model) (tea.Model, tea.Cmd) {
	var err error
	m.SelectedDeck, err = decks.LoadDeck(m.Decks[m.CurrentCard])
	if err != nil {
		panic(err.Error())
	}
	m.Cards = &m.SelectedDeck.Cards
	m.CurrentCard = 0
	return m, nil
}

func updateCard(m Model, status string) (tea.Model, tea.Cmd) {
	currentCard := m.GetCurrentCard()
	if status == "o" {
		currentCard.TimesCorrect++
	} else {
		currentCard.TimesIncorrect++
	}

	currentCard.Hidden = true

	if m.CurrentCard >= len((*m.Cards)) && len((*m.Cards)) > 0 {
		m.CurrentCard = 0
	}

	return m, nil
}

func handleKeys(m Model, msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q":
		if m.SelectedDeck != nil {
			m.CardColor = CardDefaultColor
			m.SelectedDeck = nil
			return m, nil
		}
		return m, tea.Quit
	case "left", "right":
		return changeCard(m, msg.String())
	case "k", "o":
		return updateCard(m, msg.String())
	case " ":
		return actOnCard(m)
	}
	return m, nil
}

func (m Model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.WindowSizeMsg:
		return resize(m, msg)
	case tea.KeyMsg:
		return handleKeys(m, msg)
	}
	return m, nil
}
