package app

import (
	"log"
	"net/http"
)

func Start() {
	// Delcare our own Multiplexer
	mux := http.NewServeMux()

	// Define routes
	// - Register the '/greet' endpoint with the Default Multiplexer
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomers)

	// To specify where to start the server (machine, port, etc)
	// We wrap ListenAndServe in log.Fatal in case the server fails to start
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
