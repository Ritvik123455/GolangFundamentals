package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %d", len(d))
	}
	if d[0] != "Spades of Ace" {
		t.Errorf("Expected first card to be 'Spades of Ace', but got %s", d[0])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	err := d.saveToFile("_decktesting")
	if err != nil {
		t.Errorf("Error saving deck to file: %v", err)
	}
	os.Remove("_decktesting")
}
