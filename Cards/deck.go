package main

import "fmt"

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
