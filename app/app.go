package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/y-sugiyama654/banking/domain"
	"github.com/y-sugiyama654/banking/service"
	"log"
	"net/http"
	"os"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variable not defined.")
	}
}

func Start() {
	sanityCheck()

	router := mux.NewRouter()

	// wiring
	ch := CustomerHandlers{service.NewCustomerRepository(domain.NewCustomerRepositoryDb())}
	// define routes
	router.HandleFunc("/customers", ch.getAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	// starting server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}
