package model

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mitchellh/go-wordwrap"
)

func (m Model) View() string {
	var content string

	if m.CurrentCard.IsFlipped {
		content = m.CurrentCard.Back
	} else {
		content = m.CurrentCard.Front
	}

	wrapped := wordwrap.WrapString(content, uint(m.viewport.Width)/3)

	mainContent := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(m.CardColor)).
		Padding(2, 4).
		Align(lipgloss.Center).
		Render(wrapped)

	bottomText := lipgloss.NewStyle().
		Foreground(lipgloss.Color("241")).
		Render("Press q to quit")

	combined := lipgloss.JoinVertical(
		lipgloss.Center,
		mainContent,
		"",
		bottomText,
	)

	return lipgloss.Place(
		m.viewport.Width,
		m.viewport.Height,
		lipgloss.Center,
		lipgloss.Center,
		combined,
	)
}
