package main

import "testing"

// Had to run "go mod init Cards" at ~/Documents/GitHub/Go/Cards to be able to run tests
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected first card of Four of Clubs, but got %v", d[len(d)-1])
	}
}
