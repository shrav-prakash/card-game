package main

import (
	"os"
	"testing"
)

func TestCreateDeck(t *testing.T) {
	d := createDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck of length 52, got deck of length %v", len(d))
	}

	if d[0].suit != "Spades" || d[0].number != "Ace" {
		t.Errorf("Expected first card to be Ace of Spades, instead got %v of %v", d[0].number, d[0].suit)
	}

	if d[len(d) - 1].number != "King" || d[len(d) - 1].suit != "Clubs" {
		t.Errorf("Expected last card to be King of Clubs, instead got %v of %v",d[len(d) - 1].number, d[len(d) - 1].suit)
	}
}

func TestStoreDeckAndReadDeckFromFile(t *testing.T) {
	os.Remove("decks/_decktesting")

	d := createDeck()
	d.storeDeck("_decktesting")
	readD := readDeckFromFile("_decktesting")

	var equalDecks = func (d1 deck, d2 deck) bool {
		if len(d1) != len(d2) {
			return false
		}

		for i, card := range d1 {
			if d2[i] != card {
				return false
			}
		}

		return true
	}

	if !equalDecks(d, readD) {
		t.Errorf("Initial deck and deck read from file after storage are different from each other")
	}

	os.Remove("decks/_decktesting")
}