package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

// Optional: We can use Go struct tags to modify the property names in the JSON response
type Customer struct {
	Name    string `json:"full_name" xml:"full_name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zip_code"`
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

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
