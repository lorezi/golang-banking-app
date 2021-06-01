package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/lorezi/golang-bank-app/db"
	"github.com/lorezi/golang-bank-app/handlers"
	"github.com/lorezi/golang-bank-app/logger"
	"github.com/lorezi/golang-bank-app/middleware"
	"github.com/lorezi/golang-bank-app/repositories"
	"github.com/lorezi/golang-bank-app/service"
	"github.com/subosito/gotenv"
)

func sanitizeConfigs() {
	envProps := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
		"DB_USER",
		"DB_PWD",
		"DB_ADDR",
		"DB_PORT",
		"DB_NAME",
		"AUTH_SERVER",
		"AUTH_SERVER_PORT",
	}

	for _, v := range envProps {
		if os.Getenv(v) == "" {
			logger.Error(fmt.Sprintf("Environment variable %s not defined. Terminating application...", v))
		}
	}

}

func Start() {
	gotenv.Load()

	sanitizeConfigs()
	// created multiplexer
	router := mux.NewRouter()

	dbClient := db.Connect()

	customerRepo := repositories.NewCustomerRepositoryDb(dbClient)
	accountRepo := repositories.NewAccountRepositoryDb(dbClient)
	transactionRepo := repositories.NewTransactionRepositoryDb(dbClient)

	// Testing
	// ch := handlers.CustomerHandlers{
	// 	CustomerService: service.NewCustomerService(mocks.NewCustomerRepositoryStub()),
	// }

	// wiring
	ch := handlers.CustomerHandler{
		CustomerService: service.NewCustomerService(customerRepo),
	}

	ah := handlers.AccountHandler{
		AccountService: service.NewAccountService(accountRepo),
	}

	th := handlers.TransactionHandler{
		Service: service.NewTransactionService(transactionRepo),
	}

	// defining routes

	router.HandleFunc("/customers", ch.GetAllCustomers).Methods("GET").Name("Get_Customers")

	// allow customer id with only alpha numeric and underscore character
	router.HandleFunc("/customers/{customer_id:[a-zA-Z0-9_]+}", ch.GetCustomer).Methods("GET").Name("Get_Customer")

	router.HandleFunc("/customers/{customer_id:[a-zA-Z0-9_]+}/account", ah.CreateAccount).Methods("POST").Name("New_Account")

	router.HandleFunc("/customers/{customer_id:[a-zA-Z0-9_]+}/account/{account_id:[a-zA-Z0-9_]+}", th.CreateTransaction).Methods("POST").Name("New_Transaction")

	authMiddleware := middleware.AuthMiddleware{Repo: repositories.NewAuthRepository()}
	router.Use(authMiddleware.Authentication())

	// starting serve
	addr := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", addr, port), router))
}
