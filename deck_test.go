package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 56 {
		t.Errorf("Expected deck length of 56, but got %d", len(d))
	}
	if d[0] != "As Wajik" {
		t.Errorf("Expected deck first is As wajik, but got %s", d[0])
	}
	if d[len(d)-1] != "King Keriting" {
		t.Errorf("Expected deck first is King Keriting, but got %s", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktest")
	d := newDeck()
	d.saveToFile("_decktest")
	loadedDeck := newDeckFromFile("_decktest")
	if len(loadedDeck) != 56 {
		t.Errorf("Expected deck length of 56, but got %d", len(d))
	}
	os.Remove("_decktest")
}
