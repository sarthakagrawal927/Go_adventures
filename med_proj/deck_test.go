package main

import (
	"os"
	"testing"
)

func TestCreateDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 28 {
		t.Errorf("Expected deck length of 28 but got %v", len(d))
	}

	if d[0] != "AceSpades" {
		t.Errorf("Expected first card as AceSpades but got %v", d[0])
	}

	if d[len(d)-1] != "SixClub" {
		t.Errorf("Expected last card as Six of Club but got %v", d[len(d)-1])
	}
}

func TestSaveDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckfromFile(("_decktesting"))

	if len(loadedDeck) != 28 {
		t.Errorf("Expected 28 cards but got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
