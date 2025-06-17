package model

import "github.com/charmbracelet/lipgloss"

func (m Model) View() string {
	mainContent := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(m.CardColor)).
		Padding(2, 4).
		Align(lipgloss.Center).
		Render("hello, runes\ntest")

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
