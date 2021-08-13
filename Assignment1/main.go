package main

import "fmt"

func main() {
	// Create a Slice with numbers from 0 to 10, then iterate and determine if each number is even or odd
	n := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, number := range n {
		if number%2 == 0 {
			fmt.Println(number, " is even")
		} else {
			fmt.Println(number, " is odd")
		}
	}
}
