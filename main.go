package main

import (
	"log"

	"github.com/alejandrolaguna20/runes/decks"
	"github.com/alejandrolaguna20/runes/model"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	decks.ListDecks()
	return
	p := tea.NewProgram(model.CreateModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
