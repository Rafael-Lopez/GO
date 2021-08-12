package main

import (
	"fmt"
	"strings"
)

// Create a new type 'deck'
// whcih is basically a slice of strings
type deck []string

// The '(d deck)' part is called Receiver.
// By specifying this receiver, now every any variable of type 'deck' gets access to the 'print' method
// The 'd' is basically a reference to the 'deck' variable calling the 'print' method
// Why we called it 'd'? By convention we usually refer to the receiver with a one or two letter abbreviation that
// matches the type of the receiver. In this case, because our receiver is of type 'deck', we refer to the actual copy of the reference
// to this object as 'd', just the first letter. You don't have to do this if you want to. But this is the convention.
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// _(underscore) in Golang is known as the Blank Identifier. Golang has a special feature to define and use the unused variable
	// using Blank Identifier. These variables make the program almost unreadable. Since Golang aims to be concise and readable, it
	// doesn’t allow you to define an unused variable. If you do so, the compiler will throw an error. The real use of Blank Identifier
	// comes when a function returns multiple values, but we need only some of them. Basically, it tells the compiler that this
	// variable is not needed, so it should be ignored without any error. This hides the variable’s values and makes the program
	// readable. So whenever you will assign a value to Bank Identifier it becomes unusable.
	// In this case, we are not using the indexes returned by the 'range' keyword, so we just ignore them.
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Go allows to return multiple values
func deal(d deck, handSize int) (deck, deck) {
	// Slices support a “slice” operator with the syntax slice[lowIncludingElement:highUpToButNotIncludingElement]
	// You can also opt to not specify those values, in which case Go will assume the following:
	//	a) if low not present: you want to start from the beginning
	//	b) if high not present: you want to go till the last value
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// Type conversion. We can do it because 'deck' is basically a slice of strings (line 7)
	return strings.Join([]string(d), ",")
}
