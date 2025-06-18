package model

import (
	"github.com/charmbracelet/lipgloss"
)

func ViewKeys(m Model) string {
	var keys string
	okText := lipgloss.NewStyle().Margin(0, 1).Foreground(OkColor).Render("[o] to mark it as correct")
	koText := lipgloss.NewStyle().Margin(0, 1).Foreground(KoColor).Render("[k] to mark it as incorrect")
	nextText := lipgloss.NewStyle().Margin(0, 1).Foreground(CardBlueColor).Render("[space] to flip it")
	keys = lipgloss.JoinHorizontal(lipgloss.Center, koText, okText)
	keys = lipgloss.JoinVertical(lipgloss.Center, nextText, keys)
	return keys
}

func (m Model) View() string {

	if m.SelectedDeck == nil {
		return selectionDeckView(m)
	} else {
		return deckView(m)
	}
}
