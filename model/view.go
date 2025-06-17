package model

import "github.com/charmbracelet/lipgloss"

func (m Model) View() string {
	style := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		Padding(2, 2).
		Align(lipgloss.Center)
	content := style.Render("hello, runes")
	centeredContent := lipgloss.Place(
		m.viewport.Width,
		m.viewport.Height,
		lipgloss.Center,
		lipgloss.Center,
		content,
	)
	return centeredContent
}
