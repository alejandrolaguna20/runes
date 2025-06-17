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
	m.CardColor = "164"
	return m, nil
}

func handleKeys(m Model, msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q":
		return m, tea.Quit
	case "left", "right":
		return m, nil
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
