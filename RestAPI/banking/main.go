package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Optional: We can use Go struct tags to modify the property names in the JSON response
type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	Zipcode string `json:"zip_code"`
}

func main() {
	// Define routes
	// - Register the '/greet' endpoint with the Default Multiplexer
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	// To specify where to start the server (machine, port, etc)
	// We wrap ListenAndServe in log.Fatal in case the server fails to start
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// - ResponseWriter is going to help us send the response to the client
// - Request is the one coming into the server (request boy, parameters, etc)
func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Ashish", "New Delhi", "110075"},
		{"Rob", "New Delhi", "110075"},
	}

	// To set the right content type in the response.
	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(customers)
}
