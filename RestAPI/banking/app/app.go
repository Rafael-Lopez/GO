package app

import (
	"log"
	"net/http"

	"github.com/Rafael-Lopez/GO/RESTAPI/banking/domain"
	"github.com/Rafael-Lopez/GO/RESTAPI/banking/service"
	"github.com/gorilla/mux"
)

func Start() {
	// Delcare our own Multiplexer
	//mux := http.NewServeMux()
	// Use gorilla/mux Multiplexer
	router := mux.NewRouter()

	// Wiring
	//ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// Define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	// To specify where to start the server (machine, port, etc)
	// We wrap ListenAndServe in log.Fatal in case the server fails to start
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
