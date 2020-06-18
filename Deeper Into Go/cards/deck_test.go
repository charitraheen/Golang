package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 36 {
		t.Errorf("Expected deck length of 36, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Nine of Clubs" {
		t.Errorf("Expected first card of Nine of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndDeckFromFile(t *testing.T) {
	os.Remove("_decktesting.txt")

	deck := newDeck()
	deck.saveToFile("_decktesting.txt")

	loadedDeck := newDeckFromFile("_decktesting.txt")

	if len(loadedDeck) != 36 {
		t.Errorf("Expect 36 cards in deck, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting.txt")
}
