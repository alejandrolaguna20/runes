package decks

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/alejandrolaguna20/runes/cards"
)

type Deck struct {
	Cards        []cards.Card
	Name         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	LastUsed     time.Time
	TimesTrained int
}

// TODO: think about the way/location decks are stored, maybe?
func ListDecks() ([]string, error) {
	pattern := filepath.Join("_decks/*.json")
	decks, err := filepath.Glob(pattern)
	fmt.Println(decks[0])
	return decks, err
}
