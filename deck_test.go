package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 12 {
		t.Errorf("Expected deck length of 12, but got %v", len(d))
	}

	if d[0] == "One of Spades" {
		t.Errorf("One of Spades was expected, but got %v", d[0])
	}

	if d[len(d)-1] == "Three of Heart" {
		t.Errorf("Three of Heart was expected, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {

	os.Remove("_decktesting")

	deck := newDeck()

	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 12 {
		t.Errorf("Expected 12 cards but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
