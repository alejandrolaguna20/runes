package model

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	SomeProperty string
}

func CreateModel() Model {
	return Model{
		SomeProperty: "test",
	}
}

func (m Model) Init() tea.Cmd {
	return tea.SetWindowTitle("runes")
}
