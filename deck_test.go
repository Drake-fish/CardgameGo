package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 28 {
		t.Errorf("Expected deck length of 28 but got, %v", len(d))
	}
	if d[0] != "One of Spades" {
		t.Errorf("Expected One of Spades but got %v", d[0])
	}
	if d[len(d)-1] != "Seven of Clubs" {
		t.Errorf("Expected last card of Seven of Clubs but got, %v", d[len(d)-1])
	}
}

func TestSaveToDeckaAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()

	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 28 {
		t.Errorf("Expected 28 Cards in deck got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
