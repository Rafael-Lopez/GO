package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Define routes
	// - Register the '/greet' endpoint with the Default Multiplexer
	http.HandleFunc("/greet", greet)

	// To specify where to start the server (machine, port, etc)
	// We wrap ListenAndServe in log.Fatal in case the server fails to start
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// - ResponseWriter is going to help us send the response to the client
// - Request is the one coming into the server (request boy, parameters, etc)
func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}
