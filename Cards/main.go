package main

import "fmt"

func main() {
	// Delcare a Slice
	cards := []string{"Ace of Diamonds", newCard()}
	// Add values to the slice. The append() function doesn't modify the original slice, that's why we are assigning the result to cards
	cards = append(cards, "Six of Spades")

	// Iterate through the slice
	// When ranging over a slice, two values are returned for each iteration.
	// The first is the index, and the second is a copy of the element at that index.
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
