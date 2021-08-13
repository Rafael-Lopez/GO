package main

import "fmt"

func main() {
	// Many ways to declare a map

	// a)
	// Declare a map where the keys are of type string, and the values are of type string
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4b5hjt",
	// }

	// b)
	// var colors2 map[string]string

	// c) - equivalent to b)
	// colors := make(map[string]string)
	// colors["white"] = "#ffffff"

	// fmt.Println(colors)

	// delete(colors, "white")
	// fmt.Println(colors)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4b5hjt",
		"white": "#ffffff",
	}

	printMap(colors)
}

// Pass a map as a parameter
func printMap(c map[string]string) {
	// Iterate over a map
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
