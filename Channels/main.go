package main

import (
	"fmt"
	"net/http"
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
	for {
		// <- c is blocking code
		go checkLink(<-c, c)
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
