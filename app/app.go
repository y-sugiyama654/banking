package app

import (
	"github.com/gorilla/mux"
	"github.com/y-sugiyama654/banking/domain"
	"github.com/y-sugiyama654/banking/service"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()

	// wiring
	ch := CustomerHandlers{service.NewCustomerRepository(domain.NewCustomerRepositoryStub())}
	// define routes
	router.HandleFunc("/customers", ch.getAllCustomer).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
