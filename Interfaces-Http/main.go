package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// https://pkg.go.dev/net/http@go1.16.7#Get
	resp, error := http.Get("http://google.com")

	if error != nil {
		fmt.Println("Error:", error)
		os.Exit(1)
	}

	// - Another way to create a Slice.
	// - 99999 is the number of empty spaces we want to initially allocate.
	// - Since the Read function is coded to read data until the Slice is full or no more data to read, and
	//   it doesn't resize the Slice, that's why we use 99999
	bs := make([]byte, 99999)
	resp.Body.Read(bs)

	fmt.Println(string(bs))
}
