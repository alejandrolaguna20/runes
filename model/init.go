package model

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	viewport viewport.Model
	ready    bool
}

func (m Model) Init() tea.Cmd {
	return nil
}
