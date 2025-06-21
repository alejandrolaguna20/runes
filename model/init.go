package model

import (
	//	"math/rand"

	"github.com/alejandrolaguna20/runes/cards"
	"github.com/alejandrolaguna20/runes/decks"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	viewport     viewport.Model
	CardColor    lipgloss.Color
	ready        bool
	Decks        []string
	SelectedDeck *decks.Deck
	Cards        *[]cards.Card
	CurrentCard  int
}

var OkColor = lipgloss.Color("108")
var KoColor = lipgloss.Color("167")
var CardBlueColor = lipgloss.Color("116")
var CardDefaultColor = lipgloss.Color("96")
var MuteColor = lipgloss.Color("241")

func CreateModel() Model {
	decks, err := decks.ListDeckNames()
	if err != nil {
		panic("could not load decks!")
	}

	m := Model{
		CardColor: CardDefaultColor,
		Decks:     decks,
		Cards:     nil,
	}
	m.CurrentCard = 0
	//m.CurrentCard = rand.Intn(len(m.Cards) - 1)
	return m
}

func (m Model) GetCurrentCard() *cards.Card {
	return &(*m.Cards)[m.CurrentCard]
}

func (m Model) Init() tea.Cmd {
	return nil
}
