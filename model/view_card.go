package model

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mitchellh/go-wordwrap"
)

func viewCard(m Model) string {
	var content string

	if (*m.Cards)[m.CurrentCard].IsFlipped {
		content = (*m.Cards)[m.CurrentCard].Back
	} else {
		content = (*m.Cards)[m.CurrentCard].Front
	}

	wrapped := wordwrap.WrapString(content, uint(m.viewport.Width)/3)

	card := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(m.CardColor).
		Padding(2, 4).
		Align(lipgloss.Center).
		Render(wrapped)
	return card
}
