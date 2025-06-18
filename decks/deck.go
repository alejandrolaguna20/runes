package decks

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

func ListDeckNames() ([]string, error) {
	pattern := filepath.Join("_decks", "*.json")
	filePaths, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	var deckNames []string
	for _, filePath := range filePaths {
		name := strings.TrimSuffix(filepath.Base(filePath), ".json")
		deckNames = append(deckNames, name)
	}
	return deckNames, nil
}

func LoadDeck(filePath string) (*Deck, error) {
	filePath = "_decks/" + filePath + ".json"
	data, err := os.ReadFile(filePath)
	if err != nil {
		return &Deck{}, fmt.Errorf("failed to read deck file %s: %w", filePath, err)
	}

	var deck Deck
	err = json.Unmarshal(data, &deck)
	if err != nil {
		return &Deck{}, fmt.Errorf("failed to parse deck JSON in %s: %w", filePath, err)
	}

	return &deck, nil
}

// Then load individual deck when selected
//func LoadDeck(name string) (Deck, error) {
//	filePath := filepath.Join("_decks", name+".json")
//	return loadDeckFromFile(filePath)
//}
