package model

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/mitchellh/go-wordwrap"
)

func viewDeck(m Model) string {
	content := m.Decks[m.CurrentCard]
	wrapped := wordwrap.WrapString(content, uint(m.viewport.Width)/3)

	card := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(m.CardColor).
		Padding(2, 4).
		Align(lipgloss.Center).
		Render(wrapped)
	return card
}

func selectionDeckView(m Model) string {
	bottomText := lipgloss.NewStyle().
		Foreground(MuteColor).
		Render("Press q to quit")

	decksStats := lipgloss.NewStyle().
		Foreground(MuteColor).
		Render(fmt.Sprintf("Deck %d out of %d decks", m.CurrentCard+1, len(m.Decks)))

	combined := lipgloss.JoinVertical(
		lipgloss.Center,
		decksStats,
		viewDeck(m),
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

func deckView(m Model) string {
	bottomText := lipgloss.NewStyle().
		Foreground(MuteColor).
		Render("Press q to quit")

	cardStats := lipgloss.NewStyle().
		Foreground(MuteColor).
		Render(fmt.Sprintf("Card %d out of %d cards", m.CurrentCard+1, len(m.Cards)))
	combined := lipgloss.JoinVertical(
		lipgloss.Center,
		cardStats,
		viewCard(m),
		ViewKeys(m),
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
