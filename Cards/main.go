package main

func main() {
	// cards := newDeck()
	// When you run this code, a file named "my_cards" will be created in this same folder.
	// cards.saveToFile("my_cards")

	cards := newDeckFromFile("my_cards")
	cards.print()
}
