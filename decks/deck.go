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

func SaveDeck(deck *Deck, filename string) error {
	err := os.MkdirAll("_decks", 0755)
	if err != nil {
		return fmt.Errorf("failed to create decks directory: %w", err)
	}
	deck.UpdatedAt = time.Now()
	filePath := filepath.Join("_decks", filename+".json")
	data, err := json.MarshalIndent(deck, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal deck to JSON: %w", err)
	}
	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write deck file %s: %w", filePath, err)
	}
	return nil
}
