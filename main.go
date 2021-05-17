package main

import (
	"fmt"
	"log"
	"net/http"
)

func greet(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello World!")
}

func main() {
	// defining routes
	http.HandleFunc("/greet", greet)

	// starting serve
	log.Fatal(http.ListenAndServe(":8000", nil))
}
