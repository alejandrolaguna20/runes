package model

import (
	"math/rand"

	"github.com/alejandrolaguna20/runes/cards"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	viewport    viewport.Model
	CardColor   string
	ready       bool
	Cards       []cards.Card
	CurrentCard cards.Card
}

func CreateModel() Model {
	m := Model{
		CardColor: "96",
		Cards:     cards.GenerateMockCards(),
	}
	m.CurrentCard = m.Cards[rand.Intn(len(m.Cards))]
	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}
