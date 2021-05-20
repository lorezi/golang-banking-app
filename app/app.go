package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/lorezi/golang-bank-app/db"
	"github.com/lorezi/golang-bank-app/handlers"
	"github.com/lorezi/golang-bank-app/repositories"
	"github.com/lorezi/golang-bank-app/service"
	"github.com/subosito/gotenv"
)

func sanitizeConfigs() {
	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variable not defined...")
	}

}

func Start() {
	gotenv.Load()

	sanitizeConfigs()
	// created multiplexer
	router := mux.NewRouter()

	dbClient := db.Connect()

	customerRepo := repositories.NewCustomerRepositoryDb(dbClient)
	// accountRepo := repositories.NewAccountRepositoryDb(dbClient)

	// Testing
	// ch := handlers.CustomerHandlers{
	// 	CustomerService: service.NewCustomerService(mocks.NewCustomerRepositoryStub()),
	// }

	// wiring
	ch := handlers.CustomerHandlers{
		CustomerService: service.NewCustomerService(customerRepo),
	}

	// defining routes

	router.HandleFunc("/customers", ch.GetAllCustomers).Methods("GET")
	// allow customer id with only alpha numeric and underscore character
	router.HandleFunc("/customers/{customer_id:[a-zA-Z0-9_]+}", ch.GetCustomer).Methods("GET")
	// router.HandleFunc("/customers", ch.GetAllCustomers).Methods("GET")

	// starting serve
	addr := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", addr, port), router))
}
