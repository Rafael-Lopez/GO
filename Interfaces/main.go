package main

import "fmt"

// What this interface statement right here is doing is: if there is any other type inside of our program that
// has a function called getGreeting() associated with it that returns a string, then that type is also of type bot.
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	// This is how you declare a struct when it has no properties
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (eb englishBot) getGreeting() string {
	return "Hello there!"
}

// If you don't use the receiver value (like in this chatbot example), you can simply remove
// the variable name and just leave the receiver type. This is optional.
func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
