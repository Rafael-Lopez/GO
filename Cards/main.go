package main

import "fmt"

func main() {
	// Two ways to declare a new variable

	// 1) Typical way
	// var card string = "Ace of Spades"

	// 2) Go infers from the right side what data type to use
	card := "Ace of Spades"

	fmt.Println(card)

	// Assign a new value to existing variables
	card = "New Value"
	fmt.Println(card)
}
