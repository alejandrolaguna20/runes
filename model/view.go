package model

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/mitchellh/go-wordwrap"
)

func viewCard(m Model) string {
	var content string

	if m.Cards[m.CurrentCard].IsFlipped {
		content = m.Cards[m.CurrentCard].Back
	} else {
		content = m.Cards[m.CurrentCard].Front
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

func ViewKeys(m Model) string {
	var keys string
	okText := lipgloss.NewStyle().Margin(0, 1).Foreground(OkColor).Render("[o] to mark it as correct")
	koText := lipgloss.NewStyle().Margin(0, 1).Foreground(KoColor).Render("[k] to mark it as incorrect")
	nextText := lipgloss.NewStyle().Margin(0, 1).Foreground(CardBlueColor).Render("[space] to flip it")
	keys = lipgloss.JoinHorizontal(lipgloss.Center, koText, okText)
	keys = lipgloss.JoinVertical(lipgloss.Center, nextText, keys)
	return keys
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

func (m Model) View() string {

	if m.SelectedDeck == nil {
		return selectionDeckView(m)
	} else {
		return deckView(m)
	}
}
