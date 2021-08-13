package main

import "testing"

// Had to run "go mod init Cards" at ~/Documents/GitHub/Go/Cards to be able to run tests
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
}
