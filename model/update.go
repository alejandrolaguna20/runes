package model

import (
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
	if m.Cards[m.CurrentCard].IsFlipped {
		m.CardColor = "96"
	} else {
		m.CardColor = "116"
	}
	m.Cards[m.CurrentCard].IsFlipped = !m.Cards[m.CurrentCard].IsFlipped
	return m, nil
}

func changeCard(m Model, direction string) (tea.Model, tea.Cmd) {
	var nextPos int
	nextPos = m.CurrentCard + 1
	if direction == "left" {
		nextPos = m.CurrentCard - 1
	}

	if nextPos > len(m.Cards)-1 {
		nextPos = 0
	}

	if nextPos < 0 {
		nextPos = len(m.Cards) - 1
	}

	m.CurrentCard = nextPos
	return m, nil
}

func handleKeys(m Model, msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q":
		return m, tea.Quit
	case "left", "right":
		return changeCard(m, msg.String())
	case " ":
		return flipCard(m)
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
