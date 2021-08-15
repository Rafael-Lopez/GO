package main

import (
	"fmt"
	"net/http"
	"time"
)

// Program that checks if certain websites are up
func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// Create a channel of type string
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// Infinite loop
	// for {
	// 	// <- c is blocking code
	// 	go checkLink(<-c, c)
	// }

	// Alternative FOR syntax
	// Use 'range' with a channel. This is the same as above, but a lot clearer as to what this for is doing.
	for l := range c {
		// Function Literal - In Java this is an anonymous function
		// We pass 'l' as an argument to the function literal so that the function literal uses a copy of it,
		// instead of the one that main it's using at the same time.
		go func(link string) {
			time.Sleep(5 * time.Second)
			// <- c is blocking code
			checkLink(link, c)
		}(l)
	}
}

// When passing a channel as a parameter, you need to specify the type for the channel. In this case, string.
func checkLink(link string, c chan string) {
	_, error := http.Get(link)

	if error != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up")
	c <- link
}
