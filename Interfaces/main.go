package main

import "fmt"

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

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

// There's an error, because Go doesn't support method overloading
func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}
